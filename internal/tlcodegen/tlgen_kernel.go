// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"log"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
)

// Experiments with instantiation kernel.
// TODO - after new kernel finished, remove this code

type KernelResolvedArgument struct {
	isNat    bool
	fixedNat int64 // <0 if not fixed
	tip      *KernelType
	bare     bool // vector Int is not the same as vector int, we must capture the difference somewhere
}

// family of types generated from the same combinator
type KernelCombinator struct {
	tips []*KernelType

	tlName tlast.Name
	origTL []*tlast.Combinator
}

// type generated from combinator, needs actual nat values when used in a field
type KernelType struct {
	combinator *KernelCombinator

	arguments []KernelResolvedArgument
}

type KernelResolvedArgumentWithName struct {
	KernelResolvedArgument

	wrongTypeErr error // we must add all field names to local context, because they must correctly shadow names outside, but we check the type

	NatNamePR tlast.PositionRange
	NatTypePR tlast.PositionRange

	TypeTypePR tlast.PositionRange // original template arg reference for type

	isField    bool // otherwise it is # param with name
	FieldIndex int
	name       string // param name
}

type KernelResolveContext struct {
	localArgs map[string]KernelResolvedArgumentWithName

	allowAnyConstructor bool // we can reference all constructors (functions, union elements) directly internally
}

// for lookup in maps, etc
func (t *KernelType) CanonicalString() string {
	var s strings.Builder
	s.WriteString(t.combinator.tlName.String())
	if len(t.arguments) == 0 {
		return s.String()
	}
	s.WriteString("<")
	for i, a := range t.arguments {
		if i != 0 {
			s.WriteString(",")
		}
		if a.isNat {
			if a.fixedNat >= 0 {
				s.WriteString(fmt.Sprintf("%d", a.fixedNat))
			} else {
				s.WriteString("#")
			}
		} else {
			if a.bare {
				s.WriteString("%")
			}
			s.WriteString(a.tip.CanonicalString())
		}
	}
	s.WriteString(">")
	return s.String()
}

// External arguments
func (t *KernelType) NatArgs(withFixed bool) []string {
	return t.natArgs(withFixed, nil, "")
}

func (t *KernelType) natArgs(withFixed bool, result []string, prefix string) []string {
	for i, a := range t.arguments {
		fieldName := t.combinator.origTL[0].TemplateArguments[i].FieldName
		if a.isNat {
			if withFixed || a.fixedNat < 0 {
				result = append(result, prefix+fieldName)
			}
		} else {
			result = a.tip.natArgs(withFixed, result, prefix+fieldName)
		}
	}
	return result
}

// tlgen kernel. Stores system of types, uses external knowledge to perform actual code generation
type Kernel struct {
	combinators map[tlast.Name]*KernelCombinator
	tips        map[string]*KernelType // key is CanonicalString()
}

func (k *Kernel) generateSomething(gen *Gen2, tl tlast.TL) error {
	for _, typ := range tl {
		/*
			if len(typ.TemplateArguments) == 1 && typ.TemplateArguments[0].IsNat {
				t := tlast.TypeRef{Type: typ.TypeDecl.Name, PR: typ.TypeDecl.PR}
				argT := tlast.TypeRef{Type: tlast.Name{
					Namespace: "",
					Name:      "ArgumentN",
				}}
				t.Args = append(t.Args, tlast.ArithmeticOrType{
					IsArith: false,
					T:       argT,
				})
				lrc := LocalResolveContext{allowAnyConstructor: true, localNatArgs: map[string]LocalNatArg{}}
				lrc.localNatArgs["ArgumentN"] = LocalNatArg{
					natArg: ActualNatArg{isField: true, FieldIndex: 0},
				}
				_, _, _, err = gen.getType(lrc, t)
				if err != nil {
					return nil, err
				}
			}
		*/
		if len(typ.TemplateArguments) == 0 {
			t := tlast.TypeRef{Type: typ.Construct.Name, PR: typ.Construct.NamePR}
			if !typ.IsFunction {
				t = tlast.TypeRef{Type: typ.TypeDecl.Name, PR: typ.TypeDecl.PR}
			}
			_, _, _, err := k.getType(gen, KernelResolveContext{allowAnyConstructor: true}, t)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (k *Kernel) getType(gen *Gen2, localContext KernelResolveContext, t tlast.TypeRef) (*KernelType, bool, []ActualNatArg, error) {
	var actualNatArgs []ActualNatArg
	//var resolvedT ResolvedType
	tName := t.Type.String()
	// Each named reference is either global type, global constructor, local param or local field
	if localArg, ok := localContext.localArgs[tName]; ok {
		if localArg.isNat {
			e1 := t.PR.BeautifulError(fmt.Errorf("reference to %s %q where type is required", ifString(localArg.isField, "field", "#-param"), tName))
			e2 := localArg.NatNamePR.BeautifulError(errSeeHere)
			return nil, false, nil, tlast.BeautifulError2(e1, e2)
		}
		if len(t.Args) != 0 {
			e1 := t.PR.BeautifulError(fmt.Errorf("reference to template type arg %q cannot have arguments", tName))
			e2 := localArg.NatTypePR.BeautifulError(fmt.Errorf("defined here"))
			return nil, false, nil, tlast.BeautifulError2(e1, e2)
		}
		if t.Bare { // overwrite bare
			// myUnionA = MyUnion;
			// myUnionB b:int = MyUnion;
			// wrapper {T:Type} a:%T = Wrapper T;
			// useWarpper xx:(wrapper MyUnion) = UseWrapper;
			localArg.bare = true
		}
		for _, p := range localArg.tip.NatArgs(true) {
			actualNatArgs = append(actualNatArgs, ActualNatArg{
				name: p,
			})
		}
		return localArg.tip, localArg.bare, actualNatArgs, nil
	}
	var td *tlast.Combinator
	if lt, ok := gen.typeDescriptors[tName]; ok { // order of this if-else chain is important for built-ins
		if len(lt) > 1 && t.Bare {
			// myUnionA = MyUnion;
			// myUnionB b:int = MyUnion;
			// useUnion a:%MyUnion = UseUnion;
			e1 := t.PR.BeautifulError(fmt.Errorf("reference to union %q cannot be bare", tName))
			e2 := lt[0].TypeDecl.NamePR.BeautifulError(fmt.Errorf("see more"))
			//return ResolvedType{}, nil,
			return nil, false, nil, tlast.BeautifulError2(e1, e2)
		}
		td = lt[0] // for type checking, any constructor is ok for us, because they all must have the same args
		conName := td.Construct.Name.String()
		if con2, ok := gen.singleConstructors[conName]; ok && t.Bare && !con2.IsFunction && con2.TypeDecl.Name.String() == "_" {
			// bare references to wrappers %int have int canonical form,
			// otherwise vectors, maybes and other templates will be generated twice
			t.Type = td.Construct.Name
		}
	} else if lt, ok := gen.singleConstructors[tName]; ok {
		td = lt
		t.Bare = true
		if td.TypeDecl.Name.String() != "_" {
			// We use "_" in type declaration for internal types which cannot be boxed
			// We could wish to extend this definition to user types in the future
			// If there is no boxed version, constructor name is canonical reference, otherwise
			// Type name is canonical reference. We need canonical references to avoid generating type more than once
			t.Type = td.TypeDecl.Name
		}
	} else if lt, ok := gen.allConstructors[tName]; ok {
		if !localContext.allowAnyConstructor {
			e1 := t.PR.BeautifulError(fmt.Errorf("reference to %s constructor %q is not allowed", ifString(lt.IsFunction, "function", "union"), tName))
			e2 := lt.Construct.NamePR.BeautifulError(fmt.Errorf("see more"))
			return nil, false, nil, tlast.BeautifulError2(e1, e2)
		}
		// Here type name is already in canonical form, because this code path is only internal for union members and functions
		td = lt
		t.Bare = true
	}
	if td == nil {
		return nil, false, nil, t.PR.BeautifulError(fmt.Errorf("error resolving name %q", tName))
	}
	if len(td.TemplateArguments) > len(t.Args) {
		arg := td.TemplateArguments[len(t.Args)]
		e1 := t.PRArgs.CollapseToEnd().BeautifulError(fmt.Errorf("missing template argument %q here", arg.FieldName))
		e2 := arg.PR.BeautifulError(fmt.Errorf("declared here"))
		return nil, false, nil, tlast.BeautifulError2(e1, e2)
	}
	if len(td.TemplateArguments) < len(t.Args) {
		arg := t.Args[len(td.TemplateArguments)]
		e1 := arg.T.PR.BeautifulError(fmt.Errorf("excess template argument %q here", arg.String()))
		e2 := td.TemplateArgumentsPR.BeautifulError(fmt.Errorf("arguments declared here"))
		return nil, false, nil, tlast.BeautifulError2(e1, e2)
	}
	comb := &KernelCombinator{
		tlName: t.Type,
		origTL: []*tlast.Combinator{td},
	}
	kernelType := &KernelType{combinator: comb}
	t.Args = append([]tlast.ArithmeticOrType{}, t.Args...) // copy args to avoid damaging source type
	for i, a := range t.Args {
		ta := td.TemplateArguments[i]
		aName := a.T.Type.String()
		if ta.IsNat {
			if a.IsArith {
				kernelType.arguments = append(kernelType.arguments, KernelResolvedArgument{
					isNat:    true,
					fixedNat: int64(a.Arith.Res),
				})
				actualNatArgs = append(actualNatArgs, ActualNatArg{isArith: true, Arith: a.Arith})
				continue
			}
			if localArg, ok := localContext.localArgs[aName]; ok {
				if !localArg.isNat {
					e1 := a.T.PR.BeautifulError(fmt.Errorf("error resolving reference %q to #-param %q", aName, ta.FieldName))
					e2 := localArg.NatTypePR.BeautifulError(localArg.wrongTypeErr)
					return nil, false, nil, tlast.BeautifulError2(e1, e2)
				}
				if localArg.wrongTypeErr != nil {
					e1 := a.T.PR.BeautifulError(fmt.Errorf("error resolving reference %q to #-param %q", aName, ta.FieldName))
					e2 := localArg.NatTypePR.BeautifulError(localArg.wrongTypeErr)
					return nil, false, nil, tlast.BeautifulError2(e1, e2)
				}
				kernelType.arguments = append(kernelType.arguments, KernelResolvedArgument{
					isNat:    localArg.isNat, // true due to check above
					fixedNat: localArg.fixedNat,
				})
				actualNatArgs = append(actualNatArgs, ActualNatArg{
					isField:    localArg.isField,
					FieldIndex: localArg.FieldIndex,
					name:       localArg.name,
				})
				continue
			}
			e1 := a.T.PR.BeautifulError(fmt.Errorf("error resolving reference %q to #-param %q", aName, ta.FieldName))
			e2 := ta.PR.BeautifulError(fmt.Errorf("see more"))
			return nil, false, nil, tlast.BeautifulError2(e1, e2)
		}
		if a.IsArith {
			e1 := a.T.PR.BeautifulError(fmt.Errorf("passing constant %q to Type-param %q is impossible", a.Arith.String(), ta.FieldName))
			e2 := ta.PR.BeautifulError(fmt.Errorf("declared here"))
			return nil, false, nil, tlast.BeautifulError2(e1, e2)
		}
		internalType, internalBare, internalNatArgs, err := k.getType(gen, localContext, a.T)
		if err != nil {
			return nil, false, nil, err
		}
		kernelType.arguments = append(kernelType.arguments, KernelResolvedArgument{
			tip:  internalType,
			bare: internalBare,
		})
		actualNatArgs = append(actualNatArgs, internalNatArgs...)
	}
	canonicalName := kernelType.CanonicalString()
	exist, ok := k.tips[canonicalName]
	if !ok {
		log.Printf("adding canonical type: %s\n", canonicalName)
		k.combinators[comb.tlName] = comb
		comb.tips = append(comb.tips, kernelType)
		k.tips[canonicalName] = kernelType
		// We added our type already, so others can reference it
		// Now we will iterate over our fields so all types we need are also generated
		if err := k.generateType(gen, kernelType); err != nil {
			return nil, false, nil, err
		}
		return kernelType, t.Bare, actualNatArgs, nil
	}
	exist.combinator.tips = append(exist.combinator.tips, kernelType)
	return exist, t.Bare, actualNatArgs, nil
}

func (k *Kernel) generateType(gen *Gen2, kernelType *KernelType) error {
	log.Printf("---- generating type %s", kernelType.CanonicalString())
	typeName := kernelType.combinator.tlName.String()
	tlType, ok := gen.typeDescriptors[typeName]
	if !ok {
		// we are OK with generating by constructor reference
		// resolveType must prohibit such references depending on its own logic
		if lt, ok := gen.allConstructors[typeName]; ok {
			tlType = append(tlType[:0], lt)
		} else {
			//return rt2.PR.LogicError(fmt.Errorf("attempt to generate unknown type %q", rt2.Type.String()))
			panic("attempt to generate unknown type")
		}
	}
	lrc := KernelResolveContext{
		localArgs: make(map[string]KernelResolvedArgumentWithName),
	}
	// var natParams []string
	for i, a := range tlType[0].TemplateArguments { // they are the same for all constructors
		if _, ok := lrc.localArgs[a.FieldName]; ok {
			return fmt.Errorf("collision of arguments") // TODO - beautiful error
		}
		ra := kernelType.arguments[i]
		lrc.localArgs[a.FieldName] = KernelResolvedArgumentWithName{
			KernelResolvedArgument: ra,
			isField:                false,
			FieldIndex:             0,
			name:                   a.FieldName,
		}
	}

	if typeName == "__tuple" {
		return nil // TODO
	}
	if typeName == "__vector" {
		return nil // TODO
	}
	// All customizations we want are here

	if len(tlType) != 1 {
		return nil // TODO
	}
	return k.generateTypeStruct(gen, lrc, kernelType, tlType[0])
}

func (k *Kernel) generateTypeStruct(gen *Gen2, lrc KernelResolveContext, kernelType *KernelType, tlType *tlast.Combinator) error {
	for i, field := range tlType.Fields {
		fieldType, fieldTypeBare, fieldNatArgs, err := k.getType(gen, lrc, field.FieldType)
		if err != nil {
			return err
		}
		log.Printf("    field %s type %s bare %v args %v", field.FieldName, fieldType.CanonicalString(), fieldTypeBare, fieldNatArgs)
		/*
			fieldName := field.FieldName
			if fieldName == "" {
				// TODO - it would be nice to prohibit anonymous field name, unless it is single field
				fieldName = "a" + strconv.Itoa(i)
			}
			newField := Field{
				t:            wr,
				goName:       res.fieldsDec.deconflictName(CNameToCamelName(fieldName)),
				cppName:      res.fieldsDecCPP.deconflictName(fieldName),
				originalName: field.FieldName,
				originalType: field.FieldType,
				resolvedType: fieldResolvedType,
				natArgs:      fieldNatArgs,
			}
			if field.Mask != nil {
				if field.Mask.BitNumber >= 32 {
					return field.Mask.PRBits.BeautifulError(fmt.Errorf("bitmask (%d) must be in range [0..32)", field.Mask.BitNumber))
				}
				newField.BitNumber = field.Mask.BitNumber
				localArg, ok := lrc.localNatArgs[field.Mask.MaskName]
				if !ok {
					return field.Mask.PRName.BeautifulError(fmt.Errorf("failed to resolve field mask %q reference", field.Mask.MaskName))
				}
				if localArg.wrongTypeErr != nil {
					e1 := field.Mask.PRName.BeautifulError(fmt.Errorf("field mask %q reference to field of wrong type", field.Mask.MaskName))
					e2 := localArg.TypePR.BeautifulError(localArg.wrongTypeErr)
					return tlast.BeautifulError2(e1, e2)
				}
				newField.fieldMask = &localArg.natArg
			}
			res.Fields = append(res.Fields, newField)
		*/
		if field.FieldName == "" {
			continue
		}
		if _, ok := lrc.localArgs[field.FieldName]; ok {
			return fmt.Errorf("collision of arguments") // TODO - beautiful error
		}
		lrc.localArgs[field.FieldName] = KernelResolvedArgumentWithName{
			KernelResolvedArgument: KernelResolvedArgument{
				isNat:    field.FieldType.Type.String() == "#",
				fixedNat: -1,
			},
			isField:    true,
			FieldIndex: i,
			// name:             field.FieldName,
		}
		/*
			arg := LocalNatArg{
				NamePR: field.PRName,
				TypePR: field.FieldType.PR,
				natArg: ActualNatArg{isField: true, FieldIndex: i},
			}
			if field.FieldType.Type.String() != "#" {
				arg.wrongTypeErr = fmt.Errorf("referenced field %q must have type #", field.FieldName)
			}
			if field.FieldName == "" {
				continue
			}
			if err := lrc.checkArgsCollision(field.FieldName, field.PRName, errFieldNameCollision); err != nil {
				return err
			}
			lrc.localNatArgs[field.FieldName] = arg
		*/
	}
	if tlType.IsFunction {
		resultResolvedType, resultResolvedTypeBare, resultNatArgs, err := k.getType(gen, lrc, tlType.FuncDecl)
		if err != nil {
			return err
		}
		if resultResolvedTypeBare {
			// @read a.TypeA = int;
			// @read a.TypeB = %Int;
			return tlType.FuncDecl.PR.BeautifulError(fmt.Errorf("function %q result cannot be bare", tlType.Construct.Name.String()))
		}
		log.Printf("    result type %s bare %v args %v", resultResolvedType.CanonicalString(), resultResolvedTypeBare, resultNatArgs)
	}
	return nil
}
