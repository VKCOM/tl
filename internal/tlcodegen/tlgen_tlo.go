// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//lint:file-ignore U1000 Ignore all unused code, it's not important, and will be removed ASAP
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
	tls "github.com/vkcom/tl/internal/tlcodegen/gentlo/tltls"
)

const (
	natTag    = 0x70659eff // crc32("#")
	typeTag   = 0x2cecf817 // crc32("Type")
	intTag    = 0xa8509bda // crc32(int#a8509bda ? = Int;)
	longTag   = 0x22076cba // crc32(long#22076cba ? = Long;)
	floatTag  = 0x824dab22 // crc32(float#824dab22 ? = Float;)
	doubleTag = 0x2210c154 // crc32(double#2210c154 ? = Double;)
	stringTag = 0xb5286e24 // crc32(string#b5286e24 ? = String;)

	intTadInt32    = int32(-1471112230) // int tag
	longTagInt32   = int32(570911930)   // long tag
	floatTagInt32  = int32(-2108839134) // float tag
	doubleTagInt32 = int32(571523412)   // double tag
	stringTagInt32 = int32(-1255641564) // string tag

	vectorTotalTag          = uint32(0x10133f47) // crc32(vectorTotal t:Type total_count:int vector:%Vector t = VectorTotal t) field of engine.queryShortened
	underscoreTag           = uint32(0x840e0ecc) // crc32(_ X:Type result:X = ReqResult X)
	reqResultTypeTag        = uint32(0xbde1c550) // 0xb527877d ^ 0x8cc84ce1 ^ underscoreTag
	engineQueryTag          = uint32(0xfd232246) // engine.query X:Type query:X = engine.Query
	engineQueryShortenedTag = uint32(0x1edf25a9) // crc32(engine.queryShortened query:%VectorTotal int = engine.Query)
	engineQueryTypeTag      = uint32(0xe3fc07ef) // engineQueryTag ^ engineQueryShortenedTag

	vectorTotalTypeTagInt32      = int32(0x10133f47)  // vectorTotal type tag
	underscoreTagInt32           = int32(-2079453492) // underscore type tag
	reqResultTypeTagInt32        = int32(-1109277360) // reqResult type tag
	engineQueryTagInt32          = int32(-48029114)   // engineQuery type tag
	engineQueryShortenedTagInt32 = int32(0x1edf25a9)  // engineQueryShortened type tag
	engineQueryTypeTagInt32      = int32(-470022161)  // 0xe3fc07ef = engineQueryTag ^ engineQueryShortenedTag
)

// at some point of TL schema compilation tlgen breaks primitive canonical form
// in order to keep TLO compatible predefined primitive combinator form is used
var builtinCombinators = map[string]tls.CombinatorUnion{
	"int":    tls.CombinatorV4{Name: intTadInt32, Id: "int", TypeName: intTadInt32, Left: tls.CombinatorLeftBuiltin{}.AsUnion(), Right: tls.CombinatorRight{Value: tls.TypeExpr{Name: intTadInt32}.AsUnion()}}.AsUnion(),
	"long":   tls.CombinatorV4{Name: longTagInt32, Id: "long", TypeName: longTagInt32, Left: tls.CombinatorLeftBuiltin{}.AsUnion(), Right: tls.CombinatorRight{Value: tls.TypeExpr{Name: longTagInt32}.AsUnion()}}.AsUnion(),
	"float":  tls.CombinatorV4{Name: floatTagInt32, Id: "float", TypeName: floatTagInt32, Left: tls.CombinatorLeftBuiltin{}.AsUnion(), Right: tls.CombinatorRight{Value: tls.TypeExpr{Name: floatTagInt32}.AsUnion()}}.AsUnion(),
	"double": tls.CombinatorV4{Name: doubleTagInt32, Id: "double", TypeName: doubleTagInt32, Left: tls.CombinatorLeftBuiltin{}.AsUnion(), Right: tls.CombinatorRight{Value: tls.TypeExpr{Name: doubleTagInt32}.AsUnion()}}.AsUnion(),
	"string": tls.CombinatorV4{Name: stringTagInt32, Id: "string", TypeName: stringTagInt32, Left: tls.CombinatorLeftBuiltin{}.AsUnion(), Right: tls.CombinatorRight{Value: tls.TypeExpr{Name: stringTagInt32}.AsUnion()}}.AsUnion(),
}

var builtinTypeExprUnions = map[string]tls.TypeExprUnion{
	"#":      tls.TypeExpr{Name: natTag}.AsUnion(),
	"int":    tls.TypeExpr{Name: intTadInt32, Flags: 1 << 0}.AsUnion(),
	"long":   tls.TypeExpr{Name: longTagInt32, Flags: 1 << 0}.AsUnion(),
	"float":  tls.TypeExpr{Name: floatTagInt32, Flags: 1 << 0}.AsUnion(),
	"double": tls.TypeExpr{Name: doubleTagInt32, Flags: 1 << 0}.AsUnion(),
	"string": tls.TypeExpr{Name: stringTagInt32, Flags: 1 << 0}.AsUnion(),
}

var absentTypes = map[string]tls.Type{
	"#":            {Name: natTag, Id: "#"},
	"Type":         {Name: typeTag, Id: "Type"},
	"ReqResult":    {Name: reqResultTypeTagInt32, Id: "ReqResult", ConstructorsNum: 3, Flags: 1<<4 | 1<<25, Arity: 1},
	"engine.Query": {Name: engineQueryTypeTagInt32, Id: "engine.Query", ConstructorsNum: 2, Flags: 1 << 4},
}

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
	s := tls.SchemaV4{
		Version: 0, // always 0
		Date:    int32(time.Now().Unix()),
	}
	types := make([]tls.Type, 0, len(gen.typeDescriptors)+3) // + 3 is #, Type and _ = ReqResult (last need for backward compatibility with tl2php) TODO - remove when possible
	sortedTypeDescriptors := make([]string, 0, len(gen.typeDescriptors)+3)
	for typName := range gen.typeDescriptors {
		sortedTypeDescriptors = append(sortedTypeDescriptors, typName)
	}
	sortedTypeDescriptors = append(sortedTypeDescriptors, "#")
	sortedTypeDescriptors = append(sortedTypeDescriptors, "Type")
	// should be already in schema sortedTypeDescriptors = append(sortedTypeDescriptors, "ReqResult")
	sortedTypeDescriptors = append(sortedTypeDescriptors, "engine.Query")

	slices.Sort(sortedTypeDescriptors)
	// number of types, # was added above
	for _, typName := range sortedTypeDescriptors {
		if tlsType, ok := absentTypes[typName]; ok {
			types = append(types, tlsType)
			typeTags[typName] = uint32(tlsType.Name)
			continue
		}
		typ := gen.typeDescriptors[typName]

		var typeName uint32 // it is actually id, but tl-compiler developers are факинг donkeys
		for _, c := range typ {
			typeName ^= c.Crc32()
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

		types = append(types, tls.Type{
			Name:            int32(typeName),                      // tl tag
			Id:              typ[0].TypeDecl.Name.String(),        // tl type name
			ConstructorsNum: int32(len(typ)),                      // number of constructors for type
			Flags:           flag,                                 //
			Arity:           int32(len(typ[0].TemplateArguments)), // number of params for type
			ParamsType:      paramsType,                           // written above
		})
	}

	var constructors, functions []tls.CombinatorUnion

	var typs, funcs []*tlast.Combinator
	for _, c := range gen.allConstructors {
		if c.IsFunction {
			funcs = append(funcs, c)
		} else {
			typs = append(typs, c)
		}
	}
	slices.SortFunc(typs, func(a, b *tlast.Combinator) int {
		return stringCompare(a.Construct.Name.String(), b.Construct.Name.String())
	})
	slices.SortFunc(funcs, func(a, b *tlast.Combinator) int {
		return stringCompare(a.Construct.Name.String(), b.Construct.Name.String())
	})
	sortedConstructors := append(typs, funcs...)

	for _, c := range sortedConstructors {
		constructName := c.Construct.Name.String()
		combinatorV4, ok := builtinCombinators[constructName]
		switch {
		// builtins of tlgen type system
		case constructName == "#" || constructName == "__vector" || constructName == "__tuple":
			continue
		case ok:
			constructors = append(constructors, combinatorV4)
			continue
		default:
		}

		var typeName uint32
		if c.IsFunction {
			typeName = typeTags[c.FuncDecl.Type.String()]
		} else {
			typeName = typeTags[c.TypeDecl.Name.String()]
		}
		var mc paramScope
		left := tls.CombinatorLeftUnion{}
		args := make([]tls.Arg, 0, len(c.TemplateArguments)+len(c.Fields))
		for i, ta := range c.TemplateArguments {
			var tag int32
			if ta.IsNat {
				tag = natTag
			} else {
				tag = typeTag
			}
			args = append(args, tls.Arg{
				Id:          ta.FieldName,
				Flags:       1<<17 | 1<<0 | 1<<1, // FLAG_OPT_VAR | FLAG_BARE | FLAG_NOCONS
				VarNum:      int32(i),
				ExistVarNum: 0,
				ExistVarBit: 0,
				Type: tls.TypeExpr{
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
				return nil, nil, fmt.Errorf("error on converting field %s to tls.Arg: %w", f.String(), err)
			}
			args = append(args, fieldTypeExprUnion)
			if !f.IsRepeated && f.FieldType.Type.String() == "#" {
				mc.append(f.FieldName, f.FieldType.Type.String(), len(c.TemplateArguments)+i)
			}
		}
		left = tls.CombinatorLeft{
			ArgsNum: uint32(len(args)),
			Args:    args,
		}.AsUnion()

		var right tls.TypeExprUnion
		if c.IsFunction {
			// can't use typeRefToTypeExpr: name build from constructors
			tmp := gen.typeRefToExprUnion(mc, typeTags, c.FuncDecl.Args)
			if ctx, ok := mc.find(func(mc param) bool { return mc.name == c.FuncDecl.Type.String() }); ok {
				right = tls.TypeVar{VarNum: int32(ctx.index)}.AsUnion()
			} else {
				right = tls.TypeExpr{
					Name:        int32(typeTags[c.FuncDecl.Type.String()]),
					ChildrenNum: uint32(len(tmp)),
					Children:    tmp,
				}.AsUnion()
			}
		} else {
			right = tls.TypeExpr{
				Name:        int32(typeTags[c.TypeDecl.Name.String()]),
				Flags:       0, // todo: flags
				ChildrenNum: uint32(len(c.TemplateArguments)),
				Children: func() []tls.ExprUnion {
					res := make([]tls.ExprUnion, 0, len(c.TemplateArguments))
					for i, ta := range c.TemplateArguments {
						var exprUnion tls.ExprUnion
						if ta.IsNat {
							exprUnion = tls.ExprNat{Expr: tls.NatVar{VarNum: int32(i)}.AsUnion()}.AsUnion()
						} else {
							exprUnion = tls.ExprType{Expr: tls.TypeVar{VarNum: int32(i)}.AsUnion()}.AsUnion()
						}
						res = append(res, exprUnion)
					}
					return res
				}(),
			}.AsUnion()
		}
		combinatorUnion := tls.CombinatorV4{
			Name:     int32(c.Crc32()), // primitives were corrupted, but is safe to use crc32 here, as all critical combinators have been taken care in the begging of the loop over types
			Id:       c.Construct.Name.String(),
			TypeName: int32(typeName),
			Left:     left,
			Right:    tls.CombinatorRight{Value: right},
			Flags:    modifierToFlag(c.Modifiers),
		}.AsUnion()
		if c.IsFunction {
			functions = append(functions, combinatorUnion)
		} else {
			constructors = append(constructors, combinatorUnion)
		}
	}
	// TODO - remove when possible
	reqResult := tls.CombinatorV4{
		Name:     underscoreTagInt32,
		Id:       "_",
		TypeName: reqResultTypeTagInt32,
		Left: tls.CombinatorLeft{
			ArgsNum: 2,
			Args: []tls.Arg{
				{Id: "X", Flags: 131075, Type: tls.TypeExpr{Name: typeTag}.AsUnion()},
				{Id: "result", Type: tls.TypeVar{}.AsUnion()},
			}}.AsUnion(),
		Right: tls.CombinatorRight{
			Value: tls.TypeExpr{
				Name:        reqResultTypeTagInt32,
				ChildrenNum: 1,
				Children: []tls.ExprUnion{tls.ExprType{
					Expr: tls.TypeVar{}.AsUnion(),
				}.AsUnion()}}.AsUnion()}}.AsUnion()
	engineQuery := tls.CombinatorV4{
		Name:     engineQueryTagInt32,
		Id:       "engine.query",
		TypeName: engineQueryTypeTagInt32,
		Left: tls.CombinatorLeft{
			ArgsNum: 2,
			Args: []tls.Arg{
				{Id: "X", Flags: 1<<1 | 1<<0 | 1<<17, Type: tls.TypeExpr{Name: typeTag}.AsUnion()},
				{Id: "query", Flags: 1 << 18, Type: tls.TypeVar{}.AsUnion()},
			},
		}.AsUnion(),
		Right: tls.CombinatorRight{
			Value: tls.TypeExpr{
				Name: engineQueryTypeTagInt32,
			}.AsUnion()},
	}.AsUnion()
	engineQueryShortened := tls.CombinatorV4{
		Name:     engineQueryShortenedTagInt32,
		Id:       "engine.queryShortened",
		TypeName: engineQueryTypeTagInt32,
		Left: tls.CombinatorLeft{
			ArgsNum: 1,
			Args: []tls.Arg{
				{Id: "query", Type: tls.TypeExpr{
					Name:        vectorTotalTypeTagInt32,
					Flags:       1 << 0,
					ChildrenNum: 1,
					Children: []tls.ExprUnion{
						tls.ExprType{Expr: tls.TypeExpr{
							Name:  intTadInt32,
							Flags: 1 << 0,
						}.AsUnion()}.AsUnion(),
					}}.AsUnion()},
			},
		}.AsUnion(),
		Right: tls.CombinatorRight{Value: tls.TypeExpr{
			Name: engineQueryTypeTagInt32,
		}.AsUnion()},
	}.AsUnion()
	s.TypesNum = uint32(len(types))
	s.Types = types
	s.ConstructorNum = uint32(len(constructors)) + 3
	s.Constructors = append(constructors, reqResult, engineQuery, engineQueryShortened) // TODO - remove when possible
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

func (gen *Gen2) typeRefToTypeExpr(mc paramScope, typeTags map[string]uint32, t *tlast.TypeRef, bare bool) tls.TypeExprUnion {
	if typeExpr, ok := builtinTypeExprUnions[t.Type.String()]; ok {
		return typeExpr
	}
	tmp := gen.typeRefToExprUnion(mc, typeTags, t.Args)
	var flags int32
	if t.Bare || bare {
		flags = 1
	}
	return tls.TypeExpr{
		// safe to use as primitives handled above
		Name:        int32(gen.allConstructors[t.Type.String()].Crc32()),
		Flags:       flags,            // Is type expression bare
		ChildrenNum: uint32(len(tmp)), // todo t.Args
		Children:    tmp,
	}.AsUnion()
}

func (gen *Gen2) repeatedToTypeExpr(mc paramScope, typeTags map[string]uint32, rws *tlast.RepeatWithScale, fieldIndex int) (tls.TypeExprUnion, error) {
	var res tls.Array
	if rws.ExplicitScale {
		if rws.Scale.IsArith {
			res.Multiplicity = tls.NatConst{Value: int32(rws.Scale.Arith.Res)}.AsUnion()
		} else {
			if c, ok := mc.find(func(mc param) bool { return mc.name == rws.Scale.Scale }); ok {
				res.Multiplicity = tls.NatVar{VarNum: int32(c.index)}.AsUnion()
			} else {
				return tls.TypeExprUnion{}, rws.Scale.PR.BeautifulError(errors.New("scale not found"))
			}
		}
	} else {
		if ctx, found := mc.find(func(mc param) bool { return mc.fieldIndex == fieldIndex-1 }); found {
			res.Multiplicity = tls.NatVar{VarNum: int32(ctx.index)}.AsUnion()
		} else {
			return tls.TypeExprUnion{}, rws.PR.BeautifulError(errors.New("repeated type used without size: no scale, previous field/argument type not #"))
		}
	}
	res.ArgsNum = uint32(len(rws.Rep))
	outerScopeSize := len(mc.data) // needed to throw out the parameters after leaving the scope
	for i, f := range rws.Rep {
		tlsArg, err := gen.combinatorToTypeExprUnion(mc, typeTags, &f, i)
		if err != nil {
			return tls.TypeExprUnion{}, fmt.Errorf("error on converting rep %s to tls.TypeExprUnion: %w", f.String(), err)
		}
		res.Args = append(res.Args, tlsArg)
		if !f.IsRepeated && f.FieldType.Type.String() == "#" {
			mc.append(f.FieldName, f.FieldType.Type.String(), i)
		}
	}
	mc.data = mc.data[:outerScopeSize]
	return res.AsUnion(), nil
}

func (gen *Gen2) combinatorToTypeExprUnion(mc paramScope, typeTags map[string]uint32, f *tlast.Field, fieldIndex int) (tls.Arg, error) {
	res := tls.Arg{Id: f.FieldName}
	if f.FieldType.Type.String() == "#" {
		res.Flags |= 1 << 1 // unknown
		res.VarNum = int32(mc.absoluteIndex)
		res.Type = gen.typeRefToTypeExpr(mc, typeTags, &f.FieldType, false)
	} else if f.IsRepeated {
		repeatedTypeExpr, err := gen.repeatedToTypeExpr(mc, typeTags, &f.ScaleRepeat, fieldIndex)
		if err != nil {
			return tls.Arg{}, fmt.Errorf("error on converting scale repeate %s to tls.TypeExprUnion: %w", f.ScaleRepeat.String(), err)
		}
		res.Type = repeatedTypeExpr
	} else if c, ok := mc.find(func(mc param) bool { return mc.name == f.FieldType.Type.String() }); ok {
		res.Type = tls.TypeVar{VarNum: int32(c.index)}.AsUnion()
	} else if _, okConstructor := gen.allConstructors[f.FieldType.Type.String()]; okConstructor {
		res.Type = gen.typeRefToTypeExpr(mc, typeTags, &f.FieldType, true)
	} else if typeTag, okType := typeTags[f.FieldType.Type.String()]; okType {
		tmp := gen.typeRefToExprUnion(mc, typeTags, f.FieldType.Args)
		var flags int32
		if f.FieldType.Bare {
			flags |= 1
		}
		res.Type = tls.TypeExpr{
			Name:        int32(typeTag),
			Flags:       flags,            // todo: flags
			ChildrenNum: uint32(len(tmp)), // todo uint32(len(f.FieldType.Args)),
			Children:    tmp,
		}.AsUnion()
	} else {
		return tls.Arg{}, f.FieldType.PR.BeautifulError(fmt.Errorf("type %s not found", f.FieldType.Type.String()))
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

func (gen *Gen2) typeRefToExprUnion(mc paramScope, typeTags map[string]uint32, aots []tlast.ArithmeticOrType) []tls.ExprUnion {
	var res []tls.ExprUnion
	for _, aot := range aots {
		aotTypeString := aot.T.Type.String()
		if aot.IsArith {
			res = append(res, tls.ExprNat{Expr: tls.NatConst{Value: int32(aot.Arith.Res)}.AsUnion()}.AsUnion())
			continue
		}
		if typeExpr, ok := builtinTypeExprUnions[aotTypeString]; ok {
			res = append(res, tls.ExprType{Expr: typeExpr}.AsUnion())
			continue
		}
		if ctx, ok := mc.find(func(e param) bool {
			return e.name == aotTypeString
		}); ok {
			switch ctx.typ {
			case "#":
				res = append(res, tls.ExprNat{Expr: tls.NatVar{VarNum: int32(ctx.index)}.AsUnion()}.AsUnion())
			case "Type":
				res = append(res, tls.ExprType{Expr: tls.TypeVar{VarNum: int32(ctx.index)}.AsUnion()}.AsUnion())
			}
			continue
		}
		if c, ok := gen.allConstructors[aotTypeString]; ok {
			flags := int32(1) // constructor is bare type parameter
			if aot.T.Bare {
				flags |= 1
			}
			tmp := gen.typeRefToExprUnion(mc, typeTags, aot.T.Args)
			res = append(res, tls.ExprType{Expr: tls.TypeExpr{
				// safe to use as primitives handled above
				Name:        int32(c.Crc32()),
				Flags:       flags,
				ChildrenNum: uint32(len(tmp)),
				Children:    tmp,
			}.AsUnion()}.AsUnion())
			continue
		}
		if typeTag, ok := typeTags[aotTypeString]; ok {
			var flags int32
			if aot.T.Bare {
				flags |= 1
			}
			tmp := gen.typeRefToExprUnion(mc, typeTags, aot.T.Args)
			res = append(res, tls.ExprType{Expr: tls.TypeExpr{
				Name:        int32(typeTag),
				Flags:       flags,
				ChildrenNum: uint32(len(tmp)),
				Children:    tmp,
			}.AsUnion()}.AsUnion())
			continue
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
	if !strings.HasSuffix(filepathName, tloExt) {
		filepathName = filepath.Join(filepathName, gen.RootPackageName+tloExt)
	}
	return os.WriteFile(filepathName, gen.TLO, 0644)
}
