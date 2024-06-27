// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

//lint:file-ignore U1000 Ignore all unused code, it's not important, and will be removed ASAP
package tlast

import (
	"errors"
	"fmt"
	"sort"
	"time"

	tls "github.com/vkcom/tl/internal/tlast/gentlo/tltls"
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
var builtinCombinators = map[string]tls.Combinator{
	"int":    tls.CombinatorV4{Name: intTadInt32, Id: "int", TypeName: intTadInt32, Left: tls.CombinatorLeftBuiltin{}.AsUnion(), Right: tls.CombinatorRight{Value: tls.TypeExpr0{Name: intTadInt32}.AsUnion()}}.AsUnion(),
	"long":   tls.CombinatorV4{Name: longTagInt32, Id: "long", TypeName: longTagInt32, Left: tls.CombinatorLeftBuiltin{}.AsUnion(), Right: tls.CombinatorRight{Value: tls.TypeExpr0{Name: longTagInt32}.AsUnion()}}.AsUnion(),
	"float":  tls.CombinatorV4{Name: floatTagInt32, Id: "float", TypeName: floatTagInt32, Left: tls.CombinatorLeftBuiltin{}.AsUnion(), Right: tls.CombinatorRight{Value: tls.TypeExpr0{Name: floatTagInt32}.AsUnion()}}.AsUnion(),
	"double": tls.CombinatorV4{Name: doubleTagInt32, Id: "double", TypeName: doubleTagInt32, Left: tls.CombinatorLeftBuiltin{}.AsUnion(), Right: tls.CombinatorRight{Value: tls.TypeExpr0{Name: doubleTagInt32}.AsUnion()}}.AsUnion(),
	"string": tls.CombinatorV4{Name: stringTagInt32, Id: "string", TypeName: stringTagInt32, Left: tls.CombinatorLeftBuiltin{}.AsUnion(), Right: tls.CombinatorRight{Value: tls.TypeExpr0{Name: stringTagInt32}.AsUnion()}}.AsUnion(),
}

var builtinTypeExprUnions = map[string]tls.TypeExpr{
	"#":      tls.TypeExpr0{Name: natTag}.AsUnion(),
	"int":    tls.TypeExpr0{Name: intTadInt32, Flags: 1 << 0}.AsUnion(),
	"long":   tls.TypeExpr0{Name: longTagInt32, Flags: 1 << 0}.AsUnion(),
	"float":  tls.TypeExpr0{Name: floatTagInt32, Flags: 1 << 0}.AsUnion(),
	"double": tls.TypeExpr0{Name: doubleTagInt32, Flags: 1 << 0}.AsUnion(),
	"string": tls.TypeExpr0{Name: stringTagInt32, Flags: 1 << 0}.AsUnion(),
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

func (tl TL) GenerateTLO(version uint32) (tls.SchemaV4, error) {
	allCombinators := make(map[string]*Combinator, len(tl))
	for i, comb := range tl {
		allCombinators[comb.Construct.Name.String()] = tl[i]
	}
	tlsTypes := make(map[string]*tls.Type, len(tl))
	tlsTypes["#"] = &tls.Type{Name: natTag, Id: "#"}
	tlsTypes["Type"] = &tls.Type{Name: typeTag, Id: "Type"}
	typeDeclNames := append(make([]string, 0, len(tl)), "#", "Type")
	for _, combinator := range tl {
		if combinator.IsFunction {
			continue
		}
		typName := combinator.TypeDecl.Name.String()
		_, ok := tlsTypes[typName]
		if !ok {
			typeDeclNames = append(typeDeclNames, typName)
			typ := &tls.Type{
				Name:  0,
				Id:    typName,
				Arity: int32(len(combinator.TypeDecl.Arguments)),
			}
			switch typName {
			case "Int", "Long", "Float", "Double", "String":
				(*typ).Flags |= 1 << 0 // FLAG_BARE                = (1 << 0)          ---          Is type expression bare
			}
			// binary encode of params type
			// 0 -> param type is type
			// 1 -> param type is #
			// note:
			//   paramsType depends on right side, NOT left side;
			//   p.s.:
			//     we are allowed to use range over TemplateArguments as parses prohibits
			//     to use different order in left and right sides
			//
			// example 1:
			//   arity = 3
			//   paramsType = 2 (bin    ary 0b010)
			//   foo {X:Type} {Y:#} {Z:Type} = Foo X Y Z;
			//
			// example 2:
			//    arity = 3
			//    paramsType = 1 (binary 0b001)
			//    foo {Y:#} {X:Type} {Z:Type} = Foo Y X Z;
			for i, arg := range combinator.TemplateArguments {
				if arg.IsNat {
					(*typ).ParamsType |= 1 << i
				}
			}
			tlsTypes[typName] = typ
		}

		typ := tlsTypes[typName]
		for _, f := range combinator.Fields {
			if f.Excl {
				(*typ).Flags |= 1 << 18
			}
		}
		if combinator.Construct.Name.String() == "_" {
			(*typ).Flags |= 1 << 25
		}

		(*typ).Name ^= int32(combinator.Crc32())
		(*typ).ConstructorsNum += 1
		if (*typ).ConstructorsNum > 1 {
			(*typ).Flags |= 1 << 4
		}
	}

	var constructors, functions []tls.Combinator
	for _, c := range tl {
		if builtin, ok := builtinCombinators[c.Construct.Name.String()]; ok {
			constructors = append(constructors, builtin)
			continue
		}
		constructName := c.Construct.Name.String()
		var typeName int32
		if c.IsFunction {
			if tlsType, ok := tlsTypes[c.FuncDecl.Type.String()]; ok {
				typeName = tlsType.Name
			}
		} else {
			if tlsType, ok := tlsTypes[c.TypeDecl.Name.String()]; ok {
				typeName = tlsType.Name
			}
		}
		var mc paramScope
		left := tls.CombinatorLeft{}
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
				Type: tls.TypeExpr0{
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
			if c.Construct.Name.String() == "adsKeyphrasesRanker.hierarchyResponse" && f.FieldName == "children" {
				print()
			}
			fieldTypeExprUnion, err := combinatorToTypeExprUnion(mc, allCombinators, tlsTypes, &f, i+len(c.TemplateArguments))
			if err != nil {
				return tls.SchemaV4{}, fmt.Errorf("error on converting field %s to tls.Arg: %w", f.String(), err)
			}
			args = append(args, fieldTypeExprUnion)
			if !f.IsRepeated && f.FieldType.Type.String() == "#" {
				mc.append(f.FieldName, f.FieldType.Type.String(), len(c.TemplateArguments)+i)
			}
		}
		left = tls.CombinatorLeft0{
			ArgsNum: uint32(len(args)),
			Args:    args,
		}.AsUnion()

		var right tls.TypeExpr
		if c.IsFunction {
			// can't use typeRefToTypeExpr: name build from constructors
			tmp := typeRefToExprUnion(mc, allCombinators, tlsTypes, c.FuncDecl.Args)
			if ctx, ok := mc.find(func(mc param) bool { return mc.name == c.FuncDecl.Type.String() }); ok {
				right = tls.TypeVar{VarNum: int32(ctx.index)}.AsUnion()
			} else {
				right = tls.TypeExpr0{
					Name:        typeName,
					ChildrenNum: uint32(len(tmp)),
					Children:    tmp,
				}.AsUnion()
			}
		} else {
			right = tls.TypeExpr0{
				Name:        typeName,
				Flags:       0, // todo: flags
				ChildrenNum: uint32(len(c.TypeDecl.Arguments)),
				Children: func() []tls.Expr {
					// mess with TypeDecl.Arguments and TemplateArguments caused by
					// engine.query {X:Type} query:!X = engine.Query; the only combinator with right
					// part not equal to left, still uses the assumption, that left and right parts
					// are equal which is might be wrong
					// so TODO - implement naive linear search over TemplateArguments
					res := make([]tls.Expr, 0, len(c.TypeDecl.Arguments))
					for i := range c.TypeDecl.Arguments {
						var exprUnion tls.Expr
						if c.TemplateArguments[i].IsNat {
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
			Id:       constructName,
			TypeName: typeName,
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
	types := make([]tls.Type, 0, len(typeDeclNames))
	sort.Strings(typeDeclNames)
	for _, typName := range typeDeclNames {
		types = append(types, *tlsTypes[typName])
	}
	sort.Slice(functions, func(i, j int) bool {
		v4A, _ := functions[i].AsV4()
		v4B, _ := functions[j].AsV4()
		return v4A.Id < v4B.Id
	})
	date := version
	if version == 0 {
		date = uint32(time.Now().Unix())
	}
	return tls.SchemaV4{
		Version:        int32(version), // TODO - must be changed to # in tl definition
		Date:           int32(date),    // TODO - must be changed to # in tl definition
		TypesNum:       uint32(len(types)),
		Types:          types,
		ConstructorNum: uint32(len(constructors)),
		Constructors:   constructors,
		FunctionsNum:   uint32(len(functions)),
		Functions:      functions,
	}, nil
}

func typeRefToTypeExpr(mc paramScope, allCombinators map[string]*Combinator, tlsTypes map[string]*tls.Type, t *TypeRef, bare bool) tls.TypeExpr {
	if typeExpr, ok := builtinTypeExprUnions[t.Type.String()]; ok {
		return typeExpr
	}
	tmp := typeRefToExprUnion(mc, allCombinators, tlsTypes, t.Args)
	var flags int32
	if t.Bare || bare {
		flags = 1
	}
	return tls.TypeExpr0{
		// safe to use as primitives handled above
		Name:        int32(allCombinators[t.Type.String()].Crc32()),
		Flags:       flags,            // Is type expression bare
		ChildrenNum: uint32(len(tmp)), // todo t.Args
		Children:    tmp,
	}.AsUnion()
}

func repeatedToTypeExpr(mc paramScope, allCombinators map[string]*Combinator, tlsTypes map[string]*tls.Type, rws *RepeatWithScale, fieldIndex int) (tls.TypeExpr, error) {
	var res tls.Array
	if rws.ExplicitScale {
		if rws.Scale.IsArith {
			res.Multiplicity = tls.NatConst{Value: int32(rws.Scale.Arith.Res)}.AsUnion()
		} else {
			if c, ok := mc.find(func(mc param) bool { return mc.name == rws.Scale.Scale }); ok {
				res.Multiplicity = tls.NatVar{VarNum: int32(c.index)}.AsUnion()
			} else {
				return tls.TypeExpr{}, rws.Scale.PR.BeautifulError(errors.New("scale not found"))
			}
		}
	} else {
		if ctx, found := mc.find(func(mc param) bool { return mc.fieldIndex == fieldIndex-1 }); found {
			res.Multiplicity = tls.NatVar{VarNum: int32(ctx.index)}.AsUnion()
		} else {
			return tls.TypeExpr{}, rws.PR.BeautifulError(errors.New("repeated type used without size: no scale, previous field/argument type not #"))
		}
	}
	res.ArgsNum = uint32(len(rws.Rep))
	outerScopeSize := len(mc.data) // needed to throw out the parameters after leaving the scope
	for i, f := range rws.Rep {
		tlsArg, err := combinatorToTypeExprUnion(mc, allCombinators, tlsTypes, &f, i)
		if err != nil {
			return tls.TypeExpr{}, fmt.Errorf("error on converting rep %s to tls.TypeExpr: %w", f.String(), err)
		}
		res.Args = append(res.Args, tlsArg)
		if !f.IsRepeated && f.FieldType.Type.String() == "#" {
			mc.append(f.FieldName, f.FieldType.Type.String(), i)
		}
	}
	mc.data = mc.data[:outerScopeSize]
	return res.AsUnion(), nil
}

func combinatorToTypeExprUnion(mc paramScope, allCombinators map[string]*Combinator, tlsTypes map[string]*tls.Type, f *Field, fieldIndex int) (tls.Arg, error) {
	res := tls.Arg{Id: f.FieldName}
	if f.FieldType.Type.String() == "#" {
		res.Flags |= 1 << 1 // unknown
		res.VarNum = int32(mc.absoluteIndex)
		res.Type = typeRefToTypeExpr(mc, allCombinators, tlsTypes, &f.FieldType, false)
	} else if f.IsRepeated {
		repeatedTypeExpr, err := repeatedToTypeExpr(mc, allCombinators, tlsTypes, &f.ScaleRepeat, fieldIndex)
		if err != nil {
			return tls.Arg{}, fmt.Errorf("error on converting scale repeate %s to tls.TypeExpr: %w", f.ScaleRepeat.String(), err)
		}
		res.Type = repeatedTypeExpr
	} else if c, ok := mc.find(func(mc param) bool { return mc.name == f.FieldType.Type.String() }); ok {
		res.Type = tls.TypeVar{VarNum: int32(c.index)}.AsUnion()
	} else if _, okConstructor := allCombinators[f.FieldType.Type.String()]; okConstructor {
		res.Type = typeRefToTypeExpr(mc, allCombinators, tlsTypes, &f.FieldType, true)
	} else if tlsType, okType := tlsTypes[f.FieldType.Type.String()]; okType {
		tmp := typeRefToExprUnion(mc, allCombinators, tlsTypes, f.FieldType.Args)
		var flags int32
		if f.FieldType.Bare {
			flags |= 1
		}
		res.Type = tls.TypeExpr0{
			Name:        tlsType.Name,
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

func typeRefToExprUnion(mc paramScope, allCombinators map[string]*Combinator, tlsTypes map[string]*tls.Type, aots []ArithmeticOrType) []tls.Expr {
	var res []tls.Expr
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
		if combinator, ok := allCombinators[aotTypeString]; ok {
			flags := int32(1) // constructor is bare type parameter
			if aot.T.Bare {
				flags |= 1
			}
			tmp := typeRefToExprUnion(mc, allCombinators, tlsTypes, aot.T.Args)
			res = append(res, tls.ExprType{Expr: tls.TypeExpr0{
				// safe to use as primitives handled above
				Name:        int32(combinator.Crc32()),
				Flags:       flags,
				ChildrenNum: uint32(len(tmp)),
				Children:    tmp,
			}.AsUnion()}.AsUnion())
			continue
		}
		if tlsType, ok := tlsTypes[aotTypeString]; ok {
			var flags int32
			if aot.T.Bare {
				flags |= 1
			}
			tmp := typeRefToExprUnion(mc, allCombinators, tlsTypes, aot.T.Args)
			res = append(res, tls.ExprType{Expr: tls.TypeExpr0{
				Name:        tlsType.Name,
				Flags:       flags,
				ChildrenNum: uint32(len(tmp)),
				Children:    tmp,
			}.AsUnion()}.AsUnion())
			continue
		}
		// if TypeRef was not found, it is likely the usage of '!'
		// example: @any func {X:Type} field_name:!X = X;
	}
	return res
}

func modifierToFlag(ms []Modifier) (res int32) {
	for _, m := range ms {
		switch m.Name {
		case "any":
			res |= 0
		case "read":
			res |= 1
		case "write":
			res |= 2
		case "readwrite":
			res |= 1 | 2
		case "internal":
			res |= 4
		case "kphp":
			res |= 8
		}
	}
	return res
}
