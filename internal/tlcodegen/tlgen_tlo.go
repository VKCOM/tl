// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/exp/slices"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/tlcodegen/gen_tlo"
)

const (
	natTag  = 0x70659eff // crc32("#")
	typeTag = 0x2cecf817 // crc32("Type")
)

type param struct {
	name       string
	typ        string
	fieldIndex int
	index      int // increments on each Type and #, doesn't decrement after leaving scope
}

type paramScope struct {
	absoluteIndex int
	data          []param
}

func (ps *paramScope) append(name, typ string, index int) {
	ps.data = append(ps.data, param{name, typ, index, ps.absoluteIndex})
	ps.absoluteIndex++
}

func (ps *paramScope) find(f func(e param) bool) (param, bool) {
	for i := len(ps.data) - 1; i >= 0; i-- {
		if f(ps.data[i]) {
			return ps.data[i], true
		}
	}
	return param{}, false
}

func (gen *Gen2) generateTLO() ([]byte, []byte, error) {
	typeTags := map[string]uint32{}
	s := gen_tlo.TlsSchemaV4{
		Version:  0, // always 0
		Date:     int32(time.Now().Unix()),
		TypesNum: 1,
		Types:    make([]gen_tlo.TlsType, 0, len(gen.typeDescriptors)+3), // + 3 is #, Type and _ = ReqResult (last need for backward compatibility with tl2php) TODO - remove when possible
	}
	s.Types = append(s.Types, gen_tlo.TlsType{Name: natTag, Id: "#"})

	sortedTypeDescriptors := make([]string, 0, len(gen.typeDescriptors))
	for typName := range gen.typeDescriptors {
		sortedTypeDescriptors = append(sortedTypeDescriptors, typName)
	}
	slices.Sort(sortedTypeDescriptors)
	typeInd, _ := slices.BinarySearch(sortedTypeDescriptors, "Type")

	// number of types, # was added above
	var typesNum uint32 = 1
	for i, typName := range sortedTypeDescriptors {
		if i == typeInd {
			s.Types = append(s.Types, gen_tlo.TlsType{Name: typeTag, Id: "Type"})
			typesNum++
		}
		typ := gen.typeDescriptors[typName]

		var typeName uint32 // it is actually id, but tl-compiler developers are факинг donkeys
		for _, c := range typ {
			typeName ^= c.Crc32()
		}
		// TODO - remove when possible
		if typName == "ReqResult" {
			tmp := -2079453492
			typeName ^= uint32(uint64(tmp))
		}
		typeTags[typName] = typeName

		// binary encode of params type
		// 0 -> param type is type
		// 1 -> param type is #
		// note:
		//   paramsType depends on right side, NOT left side
		// example 1:
		//   arity = 3
		//   paramsType = 2 (bin    ary 0b010)
		//   foo {X:Type} {Y:#} {Z:Type} = Foo X Y Z;
		//
		// example 2:
		//    arity = 3
		//    paramsType = 1 (binary 0b001)
		//    foo {Y:#} {X:Type} {Z:Type} = Foo Y X Z;
		var paramsType int64
		for i, arg := range typ[0].TemplateArguments {
			if arg.IsNat {
				paramsType |= 1 << i
			}
		}

		var flag int32
		switch typName {
		case "Int", "Long", "Float", "Double", "String":
			flag |= 1 << 0 // FLAG_BARE                = (1 << 0)          ---          Is type expression bare
		}
		if len(typ) > 1 {
			flag |= 1 << 4
		}

		s.Types = append(s.Types, gen_tlo.TlsType{
			Name:            int32(typeName),                      // tl tag
			Id:              typ[0].TypeDecl.Name.String(),        // tl type name
			ConstructorsNum: int32(len(typ)),                      // number of constructors for type
			Flags:           flag,                                 //
			Arity:           int32(len(typ[0].TemplateArguments)), // number of params for type
			ParamsType:      paramsType,                           // written above
		})
		typesNum++
	}
	if typeInd == len(sortedTypeDescriptors) {
		s.Types = append(s.Types, gen_tlo.TlsType{Name: typeTag, Id: "Type"})
		typesNum++
	}
	// TODO - remove when possible
	if reqResultTypeInd, ok := slices.BinarySearch(sortedTypeDescriptors, "ReqResult"); ok {
		s.Types[reqResultTypeInd+1].ConstructorsNum++ // + 1 as '#' was added to types as first element, and it doesn't present in typeDescriptors
		s.Types[reqResultTypeInd+1].Flags |= 1 << 25  // FLAG_DEFAULT_CONSTRUCTOR = (1 << 25)         ---          Does type have a default constructor, e.g. constructor that will be used if no magic is presented.
	}
	s.TypesNum = typesNum

	var constructors, functions []gen_tlo.TlsCombinatorUnion

	var types, funcs []*tlast.Combinator
	for _, c := range gen.allConstructors {
		if c.IsFunction {
			funcs = append(funcs, c)
		} else {
			types = append(types, c)
		}
	}
	slices.SortFunc(types, func(a, b *tlast.Combinator) int {
		return stringCompare(a.Construct.Name.String(), b.Construct.Name.String())
	})
	slices.SortFunc(funcs, func(a, b *tlast.Combinator) int {
		return stringCompare(a.Construct.Name.String(), b.Construct.Name.String())
	})
	sortedConstructors := append(types, funcs...)

	for _, c := range sortedConstructors {
		var typeName uint32
		if c.IsFunction {
			typeName = typeTags[c.FuncDecl.Type.String()]
		} else {
			typeName = typeTags[c.TypeDecl.Name.String()]
		}
		var mc paramScope
		left := gen_tlo.TlsCombinatorLeftUnion{}
		var primitiveFlags int32
		switch c.TypeDecl.Name.String() {
		case "Int", "Long", "Float", "Double", "String":
			left.SetBuiltin()
			primitiveFlags |= 1<<1 | 1<<2 | 1<<3 // purpose of magic for flags of primitives is unknown, but tl-compiler did, so do we
		default:
			args := make([]gen_tlo.TlsArg, 0, len(c.TemplateArguments)+len(c.Fields))
			for i, ta := range c.TemplateArguments {
				var tag int32
				if ta.IsNat {
					tag = natTag
				} else {
					tag = typeTag
				}
				args = append(args, gen_tlo.TlsArg{
					Id:          ta.FieldName,
					Flags:       1<<17 | 1<<0 | 1<<1, // FLAG_OPT_VAR | FLAG_BARE | FLAG_NOCONS
					VarNum:      int32(i),
					ExistVarNum: 0,
					ExistVarBit: 0,
					Type: gen_tlo.TlsTypeExpr{
						Name:  tag,
						Flags: 0, // todo: flags
					}.AsUnion(),
				})
				if ta.IsNat {
					mc.append(ta.FieldName, "#", i)
				} else {
					mc.append(ta.FieldName, "Type", i)
				}
			}
			for i, f := range c.Fields {
				fieldTypeExprUnion, err := gen.combinatorToTypeExprUnion(mc, typeTags, &f, i+len(c.TemplateArguments))
				if err != nil {
					return nil, nil, fmt.Errorf("error on converting field %s to gen_tlo.TlsArg: %w", f.String(), err)
				}
				args = append(args, fieldTypeExprUnion)
				if !f.IsRepeated && f.FieldType.Type.String() == "#" {
					mc.append(f.FieldName, f.FieldType.Type.String(), len(c.TemplateArguments)+i)
				}
			}
			left = gen_tlo.TlsCombinatorLeft{
				ArgsNum: uint32(len(args)),
				Args:    args,
			}.AsUnion()
		}

		var right gen_tlo.TlsTypeExprUnion
		if c.IsFunction {
			// can't use typeRefToTypeExpr: name build from constructors
			tmp := gen.typeRefToExprUnion(mc, typeTags, c.FuncDecl.Args)
			if ctx, ok := mc.find(func(mc param) bool { return mc.name == c.FuncDecl.Type.String() }); ok {
				right = gen_tlo.TlsTypeVar{VarNum: int32(ctx.index)}.AsUnion()
			} else {
				right = gen_tlo.TlsTypeExpr{
					Name:        int32(typeTags[c.FuncDecl.Type.String()]),
					ChildrenNum: uint32(len(tmp)),
					Children:    tmp,
				}.AsUnion()
			}
		} else {
			right = gen_tlo.TlsTypeExpr{
				Name:        int32(typeTags[c.TypeDecl.Name.String()]),
				Flags:       0, // todo: flags
				ChildrenNum: uint32(len(c.TemplateArguments)),
				Children: func() []gen_tlo.TlsExprUnion {
					res := make([]gen_tlo.TlsExprUnion, 0, len(c.TemplateArguments))
					for i, ta := range c.TemplateArguments {
						var exprUnion gen_tlo.TlsExprUnion
						if ta.IsNat {
							exprUnion = gen_tlo.TlsExprNat{Expr: gen_tlo.TlsNatVar{VarNum: int32(i)}.AsUnion()}.AsUnion()
						} else {
							exprUnion = gen_tlo.TlsExprType{Expr: gen_tlo.TlsTypeVar{VarNum: int32(i)}.AsUnion()}.AsUnion()
						}
						res = append(res, exprUnion)
					}
					return res
				}(),
			}.AsUnion()
		}
		combinatorUnion := gen_tlo.TlsCombinatorV4{
			Name:     int32(c.Crc32()),
			Id:       c.Construct.Name.String(),
			TypeName: int32(typeName),
			Left:     left,
			Right:    gen_tlo.TlsCombinatorRight{Value: right},
			Flags:    modifierToFlag(c.Modifiers) | primitiveFlags,
		}.AsUnion()
		if c.IsFunction {
			functions = append(functions, combinatorUnion)
		} else {
			constructors = append(constructors, combinatorUnion)
		}
	}
	// TODO - remove when possible
	reqResult := gen_tlo.TlsCombinatorV4{
		Name:     -2079453492,
		Id:       "_",
		TypeName: -1109277360,
		Left: gen_tlo.TlsCombinatorLeft{
			ArgsNum: 2,
			Args: []gen_tlo.TlsArg{
				{Id: "X", Flags: 131075, Type: gen_tlo.TlsTypeExpr{Name: 753727511}.AsUnion()},
				{Id: "result", Type: gen_tlo.TlsTypeVar{}.AsUnion()},
			}}.AsUnion(),
		Right: gen_tlo.TlsCombinatorRight{
			Value: gen_tlo.TlsTypeExpr{
				Name:        -1109277360,
				ChildrenNum: 1,
				Children: []gen_tlo.TlsExprUnion{gen_tlo.TlsExprType{
					Expr: gen_tlo.TlsTypeVar{}.AsUnion(),
				}.AsUnion()}}.AsUnion()}}.AsUnion()

	s.ConstructorNum = uint32(len(constructors)) + 1 // + _ = ReqResult TODO - remove when possible
	s.Constructors = append(constructors, reqResult) // TODO - remove when possible
	s.FunctionsNum = uint32(len(functions))
	s.Functions = functions
	res, err := s.WriteBoxed(nil)
	if err != nil {
		return nil, nil, fmt.Errorf("can't write TLO boxed: %w", err)
	}
	res2, err := s.WriteJSON(nil)
	if err != nil {
		return nil, nil, fmt.Errorf("can't write TLO json: %w", err)
	}
	return res, res2, nil
}

func (gen *Gen2) typeRefToTypeExpr(mc paramScope, typeTags map[string]uint32, t *tlast.TypeRef, bare bool) gen_tlo.TlsTypeExprUnion {
	if t.Type.String() == "#" {
		return gen_tlo.TlsTypeExpr{Name: natTag}.AsUnion()
	}
	tmp := gen.typeRefToExprUnion(mc, typeTags, t.Args)
	var flags int32
	if t.Bare || bare {
		flags = 1
	}
	return gen_tlo.TlsTypeExpr{
		Name:        int32(gen.allConstructors[t.Type.String()].Crc32()),
		Flags:       flags,            // Is type expression bare
		ChildrenNum: uint32(len(tmp)), // todo t.Args
		Children:    tmp,
	}.AsUnion()
}

func (gen *Gen2) repeatedToTypeExpr(mc paramScope, typeTags map[string]uint32, rws *tlast.RepeatWithScale, fieldIndex int) (gen_tlo.TlsTypeExprUnion, error) {
	var res gen_tlo.TlsArray
	if rws.ExplicitScale {
		if rws.Scale.IsArith {
			res.Multiplicity = gen_tlo.TlsNatConst{Value: int32(rws.Scale.Arith.Res)}.AsUnion()
		} else {
			if c, ok := mc.find(func(mc param) bool { return mc.name == rws.Scale.Scale }); ok {
				res.Multiplicity = gen_tlo.TlsNatVar{VarNum: int32(c.index)}.AsUnion()
			} else {
				return gen_tlo.TlsTypeExprUnion{}, rws.Scale.PR.BeautifulError(errors.New("scale not found"))
			}
		}
	} else {
		if ctx, found := mc.find(func(mc param) bool { return mc.fieldIndex == fieldIndex-1 }); found {
			res.Multiplicity = gen_tlo.TlsNatVar{VarNum: int32(ctx.index)}.AsUnion()
		} else {
			return gen_tlo.TlsTypeExprUnion{}, rws.PR.BeautifulError(errors.New("repeated type used without size: no scale, previous field/argument type not #"))
		}
	}
	res.ArgsNum = uint32(len(rws.Rep))
	outerScopeSize := len(mc.data) // needed to throw out the parameters after leaving the scope
	for i, f := range rws.Rep {
		tlsArg, err := gen.combinatorToTypeExprUnion(mc, typeTags, &f, i)
		if err != nil {
			return gen_tlo.TlsTypeExprUnion{}, fmt.Errorf("error on converting rep %s to gen_tlo.TlsTypeExprUnion: %w", f.String(), err)
		}
		res.Args = append(res.Args, tlsArg)
		if !f.IsRepeated && f.FieldType.Type.String() == "#" {
			mc.append(f.FieldName, f.FieldType.Type.String(), i)
		}
	}
	mc.data = mc.data[:outerScopeSize]
	return res.AsUnion(), nil
}

func (gen *Gen2) combinatorToTypeExprUnion(mc paramScope, typeTags map[string]uint32, f *tlast.Field, fieldIndex int) (gen_tlo.TlsArg, error) {
	res := gen_tlo.TlsArg{Id: f.FieldName}
	if f.FieldType.Type.String() == "#" {
		res.Flags |= 1 << 1 // unknown
		res.VarNum = int32(mc.absoluteIndex)
		res.Type = gen.typeRefToTypeExpr(mc, typeTags, &f.FieldType, false)
	} else if f.IsRepeated {
		repeatedTypeExpr, err := gen.repeatedToTypeExpr(mc, typeTags, &f.ScaleRepeat, fieldIndex)
		if err != nil {
			return gen_tlo.TlsArg{}, fmt.Errorf("error on converting scale repeate %s to gen_tlo.TlsTypeExprUnion: %w", f.ScaleRepeat.String(), err)
		}
		res.Type = repeatedTypeExpr
	} else if c, ok := mc.find(func(mc param) bool { return mc.name == f.FieldType.Type.String() }); ok {
		res.Type = gen_tlo.TlsTypeVar{VarNum: int32(c.index)}.AsUnion()
	} else if _, okConstructor := gen.allConstructors[f.FieldType.Type.String()]; okConstructor {
		res.Type = gen.typeRefToTypeExpr(mc, typeTags, &f.FieldType, true)
	} else if typeTag, okType := typeTags[f.FieldType.Type.String()]; okType {
		tmp := gen.typeRefToExprUnion(mc, typeTags, f.FieldType.Args)
		var flags int32
		if f.FieldType.Bare {
			flags |= 1
		}
		res.Type = gen_tlo.TlsTypeExpr{
			Name:        int32(typeTag),
			Flags:       flags,            // todo: flags
			ChildrenNum: uint32(len(tmp)), // todo uint32(len(f.FieldType.Args)),
			Children:    tmp,
		}.AsUnion()
	} else {
		return gen_tlo.TlsArg{}, f.FieldType.PR.BeautifulError(fmt.Errorf("type %s not found", f.FieldType.Type.String()))
	}
	if f.Excl {
		res.Flags |= 1 << 18 // Is argument a forwarded function (via !)
	}
	if f.Mask == nil {
		return res, nil
	}
	res.Flags |= 1 << 2 // field_mask
	if ctx, ok := mc.find(func(e param) bool { return e.name == f.Mask.MaskName }); ok {
		res.ExistVarNum = int32(ctx.index)
		res.ExistVarBit = int32(f.Mask.BitNumber)
	}
	return res, nil
}

func (gen *Gen2) typeRefToExprUnion(mc paramScope, typeTags map[string]uint32, aots []tlast.ArithmeticOrType) []gen_tlo.TlsExprUnion {
	var res []gen_tlo.TlsExprUnion
	for _, aot := range aots {
		aotTypeString := aot.T.Type.String()
		if aot.IsArith {
			res = append(res, gen_tlo.TlsExprNat{Expr: gen_tlo.TlsNatConst{Value: int32(aot.Arith.Res)}.AsUnion()}.AsUnion())
			continue
		}
		if ctx, ok := mc.find(func(e param) bool {
			return e.name == aotTypeString
		}); ok {
			switch ctx.typ {
			case "#":
				res = append(res, gen_tlo.TlsExprNat{Expr: gen_tlo.TlsNatVar{VarNum: int32(ctx.index)}.AsUnion()}.AsUnion())
			case "Type":
				res = append(res, gen_tlo.TlsExprType{Expr: gen_tlo.TlsTypeVar{VarNum: int32(ctx.index)}.AsUnion()}.AsUnion())
			}
		}
		if c, ok := gen.allConstructors[aotTypeString]; ok {
			flags := int32(1) // constructor is bare type parameter
			if aot.T.Bare {
				flags |= 1
			}
			tmp := gen.typeRefToExprUnion(mc, typeTags, aot.T.Args)
			res = append(res, gen_tlo.TlsExprType{Expr: gen_tlo.TlsTypeExpr{
				Name:        int32(c.Crc32()),
				Flags:       flags,
				ChildrenNum: uint32(len(tmp)),
				Children:    tmp,
			}.AsUnion()}.AsUnion())
		}
		if typeTag, ok := typeTags[aotTypeString]; ok {
			var flags int32
			if aot.T.Bare {
				flags |= 1
			}
			tmp := gen.typeRefToExprUnion(mc, typeTags, aot.T.Args)
			res = append(res, gen_tlo.TlsExprType{Expr: gen_tlo.TlsTypeExpr{
				Name:        int32(typeTag),
				Flags:       flags,
				ChildrenNum: uint32(len(tmp)),
				Children:    tmp,
			}.AsUnion()}.AsUnion())
		}
		if aotTypeString == "#" {
			res = append(res, gen_tlo.TlsExprType{Expr: gen_tlo.TlsTypeExpr{
				Name:  natTag,
				Flags: 0, // todo: flags
			}.AsUnion()}.AsUnion())
		}
	}
	return res
}

func modifierToFlag(ms []tlast.Modifier) (res int32) {
	for _, m := range ms {
		switch m.Name {
		case "@any":
			res |= 0
		case "@read":
			res |= 1
		case "@write":
			res |= 2
		case "@readwrite":
			res |= 1 | 2
		case "@internal":
			res |= 4
		case "@kphp":
			res |= 8
		}
	}
	return res
}

func (gen *Gen2) WriteTLO() error {
	if gen.options.TLOPath == "" {
		return nil
	}
	filepathName := gen.options.TLOPath
	if !strings.HasSuffix(filepathName, ".tlo") {
		filepathName = filepath.Join(filepathName, gen.RootPackageName+tloExt)
	}
	return os.WriteFile(filepathName, gen.TLO, 0644)
}
