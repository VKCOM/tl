package tlcodegen

import (
	"fmt"
	"github.com/vkcom/tl/internal/tlast"
	"math"
	"strconv"
	"strings"
)

func (gen *Gen2) validateTL2AstAndCollectInfo(tl2 tlast.TL2File) error {
	if gen.tl2CombinatorsOrder == nil {
		gen.tl2CombinatorsOrder = make(map[string]int)
		gen.tl2Combinators = make(map[string]*tlast.TL2Combinator)
	}
	for i, combinator := range tl2.Combinators {
		s := combinator.ReferenceName().String()
		if prevCombinator, ok := gen.tl2Combinators[s]; ok {
			return tlast.BeautifulError2(
				combinator.ReferenceNamePR().BeautifulError(fmt.Errorf("this definition of combinator is duplicated")),
				prevCombinator.ReferenceNamePR().BeautifulError(fmt.Errorf("first apperance")),
			)
		}
		if !combinator.IsFunction {
			visitedTypeArgNames := make(map[string]tlast.PositionRange)
			for _, argument := range combinator.TypeDecl.TemplateArguments {
				if pr, ok := visitedTypeArgNames[argument.Name]; ok {
					return tlast.BeautifulError2(
						argument.PRName.BeautifulError(fmt.Errorf("name repeats several times (all names must be unique)")),
						pr.BeautifulError(fmt.Errorf("first apperance")),
					)
				}
				visitedTypeArgNames[argument.Name] = argument.PRName
			}
		}
		gen.tl2CombinatorsOrder[s] = i
		gen.tl2Combinators[s] = &tl2.Combinators[i]
	}
	return nil
}

type ResolvedTL2References struct {
	ResolvedTypes map[string]tlast.TL2TypeRef
	ResolvedNats  map[string]uint32
}

func (rtl2c *ResolvedTL2References) resolveRef(ref tlast.TL2TypeRef) (newRef tlast.TL2TypeRef, err error) {
	newRef.PR = ref.PR
	if ref.IsBracket {
		newRef.IsBracket = true
		if ref.BracketType == nil {
			err = ref.PR.BeautifulError(fmt.Errorf("no bracket parsed"))
			return
		}
		newRef.BracketType = new(tlast.TL2BracketType)
		oldBracket := ref.BracketType
		newBracket := newRef.BracketType

		newBracket.PR = oldBracket.PR

		if oldBracket.IndexType != nil {
			newBracket.IndexType = new(tlast.TL2TypeArgument)

			oldIndex := oldBracket.IndexType
			newIndex := newBracket.IndexType

			newIndex.PR = oldIndex.PR

			if oldIndex.IsNumber {
				newIndex.IsNumber = true
				newIndex.Number = oldIndex.Number
			} else {
				if resolvedNumber, ok := rtl2c.ResolvedNats[oldIndex.Type.String()]; ok {
					newIndex.IsNumber = true
					newIndex.Number = resolvedNumber
				} else {
					newIndex.Type, err = rtl2c.resolveRef(oldIndex.Type)
					if err != nil {
						return
					}
				}
			}
		}
		newBracket.ArrayType, err = rtl2c.resolveRef(oldBracket.ArrayType)
		if err != nil {
			return
		}
	} else {
		oldType := ref.SomeType
		if oldType == nil {
			err = ref.PR.BeautifulError(fmt.Errorf("expected type to be parsed"))
			return
		}
		tp := ref.SomeType
		refName := tp.Name
		if resolvedRef, ok := rtl2c.ResolvedTypes[refName.String()]; ok {
			if len(tp.Arguments) != 0 {
				err = tp.PRArguments.BeautifulError(fmt.Errorf("generic type can't have arguments"))
				return
			}
			newRef = resolvedRef
			return
		}
		if _, ok := rtl2c.ResolvedNats[refName.String()]; ok {
			err = tp.PRArguments.BeautifulError(fmt.Errorf("reference to number generic can't be type"))
			return
		}

		newRef.SomeType = new(tlast.TL2TypeApplication)
		newType := newRef.SomeType

		newType.Name = oldType.Name
		newType.PRName = oldType.PRName
		newType.PRArguments = oldType.PRArguments

		for i, argument := range oldType.Arguments {
			newType.Arguments = append(newType.Arguments, tlast.TL2TypeArgument{})
			newType.Arguments[i].PR = argument.PR

			if argument.IsNumber {
				newType.Arguments[i].IsNumber = true
				newType.Arguments[i].Number = argument.Number
			} else {
				if !argument.Type.IsBracket {
					if argument.Type.SomeType == nil {
						err = ref.PR.BeautifulError(fmt.Errorf("expected type to be parsed"))
						return
					}
					if resolvedNumber, ok := rtl2c.ResolvedNats[argument.Type.SomeType.Name.String()]; ok {
						newType.Arguments[i].IsNumber = true
						newType.Arguments[i].Number = resolvedNumber
						continue
					}
				}
				newType.Arguments[i].Type, err = rtl2c.resolveRef(argument.Type)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func (gen *Gen2) genTypeTL2(resolvedRef tlast.TL2TypeRef) (*TypeRWWrapper, error) {
	reduction := resolvedRef.String()
	if wr, ok := gen.generatedTypes[reduction]; ok {
		return wr, nil
	}
	if pr, ok := gen.builtinTypes[reduction]; ok {
		return pr, nil
	}
	kernelType := TypeRWWrapper{
		gen:              gen,
		originateFromTL2: true,
		wantsTL2:         true,
	}

	if !gen.isTL1Ref(resolvedRef) {
		gen.generatedTypes[reduction] = &kernelType
		gen.generatedTypesList = append(gen.generatedTypesList, &kernelType)
	}

	if resolvedRef.IsBracket {
		if resolvedRef.BracketType == nil {
			return nil, resolvedRef.PR.BeautifulError(fmt.Errorf("expected bracket type declaration but it wasn't parsed"))
		}
		return gen.genBracketTypeTL2(&kernelType, *resolvedRef.BracketType)
	}
	if resolvedRef.SomeType == nil {
		return nil, resolvedRef.PR.BeautifulError(fmt.Errorf("expected reference to type but it wasn't parsed"))
	}
	typeApplication := *resolvedRef.SomeType
	name := typeApplication.Name
	comb, ok := gen.tl2Combinators[name.String()]
	if name.String() == "bool" || name.String() == "legacy_bool" {
		return gen.genBoolTL2(&kernelType, name.String() == "legacy_bool")
	}
	if !ok {
		return nil, typeApplication.PRName.BeautifulError(fmt.Errorf("reference to unknown type %q", name))
	}
	if comb.IsFunction {
		return gen.genFunctionTL2(&kernelType, comb)
	}
	typeDeclaration := comb.TypeDecl
	if len(typeApplication.Arguments) != len(typeDeclaration.TemplateArguments) {
		return nil, tlast.BeautifulError2(
			typeApplication.PRArguments.BeautifulError(
				fmt.Errorf("unexpected number of type arguments (%d instead if %d)",
					len(typeApplication.Arguments),
					len(typeDeclaration.TemplateArguments),
				),
			),
			typeDeclaration.PR.BeautifulError(fmt.Errorf("original")),
		)
	}

	argNamespaces := map[string]struct{}{}
	resolveMapping := ResolvedTL2References{
		ResolvedNats:  map[string]uint32{},
		ResolvedTypes: map[string]tlast.TL2TypeRef{},
	}

	for i, argument := range typeDeclaration.TemplateArguments {
		actualArg := typeApplication.Arguments[i]
		if argument.Category.IsUint32() {
			if !actualArg.IsNumber {
				return nil, actualArg.PR.BeautifulError(fmt.Errorf("by definition of this type here can be either number or generic if uint32"))
			}
			resolveMapping.ResolvedNats[argument.Name] = actualArg.Number
			kernelType.arguments = append(kernelType.arguments, ResolvedArgument{isNat: true, isArith: true, Arith: tlast.Arithmetic{Res: actualArg.Number}})
		} else if argument.Category.IsType() {
			if actualArg.IsNumber {
				return nil, actualArg.PR.BeautifulError(fmt.Errorf("by definition of this type here can be only type reference"))
			}
			resolveMapping.ResolvedTypes[argument.Name] = actualArg.Type
			wr, err := gen.genTypeTL2(actualArg.Type)
			if err != nil {
				return nil, err
			}
			_, isUnion := wr.trw.(*TypeRWUnion)
			kernelType.arguments = append(kernelType.arguments, ResolvedArgument{tip: wr, bare: !isUnion})
			collectArgsNamespaces(wr, argNamespaces)
		}
	}

	if comb.HasAnnotation(tl1Ref) {
		kernelType.originateFromTL2 = false

		notParsedError := comb.TypeDecl.PRName.BeautifulError(fmt.Errorf("can't find reference to tl1-type"))
		found, comment := extractTLGenTag(comb.CommentBefore, "tlgen:tl1type")

		if !found {
			return nil, notParsedError
		}
		parts := strings.Split(comment, " ")
		if len(parts) == 0 {
			return nil, notParsedError
		}
		typeNamespace := ""
		typeName := parts[0]
		if i := strings.Index(typeName, "."); i != -1 {
			typeNamespace = typeName[:i]
			typeName = typeName[i+1:]
		}
		tl1Ref := tlast.TypeRef{
			Type: TypeName{Namespace: typeNamespace, Name: typeName},
		}
		tl1Context := LocalResolveContext{
			localTypeArgs: map[string]LocalTypeArg{},
			localNatArgs:  map[string]LocalNatArg{},
		}
		for _, arg := range parts[1:] {
			tl1Ref.Args = append(tl1Ref.Args, tlast.ArithmeticOrType{T: tlast.TypeRef{Type: TypeName{Name: arg}}})
			tl2Arg := strings.ToLower(arg[:1]) + arg[1:]
			if number, ok := resolveMapping.ResolvedNats[tl2Arg]; ok {
				tl1Context.localNatArgs[arg] = LocalNatArg{
					natArg: ActualNatArg{
						isArith: true,
						Arith: tlast.Arithmetic{
							Nums: []uint32{number},
							Res:  number,
						},
						isField: false,
					},
				}
			} else if ref, ok := resolveMapping.ResolvedTypes[tl2Arg]; ok {
				wr, err := gen.genTypeTL2(ref)
				if err != nil {
					return nil, err
				}
				_, isUnion := wr.trw.(*TypeRWUnion)
				tl1Context.localTypeArgs[tl2Arg] = LocalTypeArg{
					arg: ResolvedArgument{
						tip:  wr,
						bare: !isUnion,
					},
				}
			} else {
				tl1Context.localNatArgs[arg] = LocalNatArg{
					natArg: ActualNatArg{
						isArith: false,
						isField: false,
						name:    tl2Arg,
					},
				}
			}
		}
		wr, _, _, _, err := gen.getType(tl1Context, tl1Ref, nil)
		wr.wantsTL2 = true
		return wr, err
	}

	if comb.HasAnnotation(tl2Maybe) {
		return gen.genMaybeTL2(&kernelType, comb, resolveMapping)
	}

	// some namespace optimization
	replaceNamespace := func(n string) *Namespace {
		newNamespace := n
		if n == "" && len(argNamespaces) == 1 {
			for ns := range argNamespaces {
				newNamespace = ns
			}
		}
		return gen.getNamespace(newNamespace)
	}
	argTail := kernelType.wrapperNameTail()

	// calculate exact type
	kernelType.tl2Name = comb.TypeDecl.Name
	kernelType.tl2Origin = comb

	// TODO: for tl1 meta
	kernelType.tlName = tlast.Name{
		Namespace: kernelType.tl2Name.Namespace,
		Name:      kernelType.tl2Name.Name,
	}
	if comb.TypeDecl.ID != nil {
		kernelType.tlTag = *comb.TypeDecl.ID
	}

	kernelType.ns = replaceNamespace(comb.TypeDecl.Name.Namespace)
	kernelType.ns.types = append(kernelType.ns.types, &kernelType)
	kernelType.fileName = comb.TypeDecl.Name.String()
	kernelType.goLocalName, kernelType.goGlobalName = getCombinatorNames(*comb, argTail)

	err := gen.genTypeDeclarationTL2(&kernelType, comb.TypeDecl.Type, resolveMapping, resolvedRef)

	return &kernelType, err
}

func (gen *Gen2) genTypeDeclarationTL2(
	kernelType *TypeRWWrapper,
	typeDecl tlast.TL2TypeDefinition,
	resolveMapping ResolvedTL2References,
	originalRef tlast.TL2TypeRef,
) error {
	argTail := kernelType.wrapperNameTail()

	switch {
	case typeDecl.IsUnionType:
		union := TypeRWUnion{
			wr:     kernelType,
			Fields: []Field{},
			IsEnum: false,
		}
		kernelType.trw = &union
		hasNonEnum := false
		yetCreatedNames := make(map[string]bool)
		yetParsedNames := make(map[string]int)
		for i, variant := range typeDecl.UnionType.Variants {
			if !variant.IsTypeAlias {
				hasNonEnum = hasNonEnum || len(variant.Fields) > 0
			}
			variantType := TypeRWStruct{}
			variantWrapper := TypeRWWrapper{
				gen: gen,
				trw: &variantType,

				originateFromTL2: true,
				wantsTL2:         true,

				fileName: kernelType.fileName,

				unionParent: &union,
				unionIndex:  i,

				// TODO: keep it sync with tl1
				tlTag: uint32(i),
			}
			variantType.wr = &variantWrapper
			field := Field{
				t: &variantWrapper,
			}
			var err error
			var targetNamespace string
			targetNamespace, variantWrapper.goLocalName, variantWrapper.goGlobalName, variantWrapper.tlName, field.goName, err = getVariantNames(kernelType.tl2Name, typeDecl.UnionType.Variants, i, yetCreatedNames, argTail)

			field.originalName = variantWrapper.tlName.String()
			field.goName = union.fieldsDec.deconflictName(field.goName)

			yetCreatedNames[field.goName] = true
			if index, ok := yetParsedNames[variantWrapper.tlName.String()]; ok {
				return tlast.BeautifulError2(
					variant.PRName.BeautifulError(fmt.Errorf("can't have such name (or it's tl1 legacy variant) due to duplication")),
					typeDecl.UnionType.Variants[index].PRName.BeautifulError(fmt.Errorf("first apperance of such name (or it's tl1 legacy variant)")),
				)
			} else {
				yetParsedNames[variantWrapper.tlName.String()] = i
			}

			variantWrapper.ns = gen.getNamespace(targetNamespace)
			variantWrapper.ns.types = append(variantWrapper.ns.types, &variantWrapper)

			currentRef := originalRef
			currentRef.SomeType = new(tlast.TL2TypeApplication)
			*currentRef.SomeType = *originalRef.SomeType
			currentRef.SomeType.Name.Name = originalRef.SomeType.Name.Name + variant.Name

			variantReduction := currentRef.String()
			gen.generatedTypes[variantReduction] = &variantWrapper
			gen.generatedTypesList = append(gen.generatedTypesList, &variantWrapper)

			if err != nil {
				return err
			}
			if variant.IsTypeAlias {
				resolvedTypedef, err := resolveMapping.resolveRef(variant.TypeAlias)
				if err != nil {
					return err
				}
				typeDefWr, err := gen.genTypeTL2(resolvedTypedef)
				if err != nil {
					return err
				}
				variantType.Fields = append(variantType.Fields, Field{})
				variantType.Fields[0].t = typeDefWr
				_, isUnion := typeDefWr.trw.(*TypeRWUnion)
				variantType.Fields[0].bare = !isUnion
			} else {
				err = gen.genFieldsTL2(resolveMapping, &variantType.Fields, variant.Fields)
				if err != nil {
					return err
				}
			}

			_, isUnion := field.t.trw.(*TypeRWUnion)
			field.bare = !isUnion
			union.Fields = append(union.Fields, field)
		}
		union.IsEnum = !hasNonEnum
	case typeDecl.IsConstructorFields:
		strct := TypeRWStruct{
			wr:     kernelType,
			Fields: []Field{},
		}
		kernelType.trw = &strct
		err := gen.genFieldsTL2(resolveMapping, &strct.Fields, typeDecl.ConstructorFields)
		if err != nil {
			return err
		}
	default:
		kernelInterface := TypeRWStruct{
			wr: kernelType,
			Fields: []Field{
				{
					originalName: "",
				},
			},
		}
		kernelType.trw = &kernelInterface
		resolvedTypedef, err := resolveMapping.resolveRef(typeDecl.TypeAlias)
		if err != nil {
			return err
		}
		typeDefWr, err := gen.genTypeTL2(resolvedTypedef)
		if err != nil {
			return err
		}
		_, isUnion := typeDefWr.trw.(*TypeRWUnion)
		kernelInterface.Fields[0].t = typeDefWr
		kernelInterface.Fields[0].bare = !isUnion
	}

	return nil
}

func (gen *Gen2) genFieldsTL2(resolveMapping ResolvedTL2References, fields *[]Field, refFields []tlast.TL2Field) error {
	for i, refField := range refFields {
		// init
		field := Field{
			originalName: refField.Name,
			goName:       snakeToCamelCase(refField.Name),
		}
		// add fieldmask
		if refField.IsOptional {
			field.fieldMask = new(ActualNatArg)
			field.fieldMask.isField = true
			field.fieldMask.FieldIndex = -((i+1)/8 + 1)
			field.BitNumber = uint32((i + 1) % 8)

			field.goName = strings.ToLower(field.goName[:1]) + field.goName[1:]
		}
		// add type
		resolvedRefType, err := resolveMapping.resolveRef(refField.Type)
		if err != nil {
			return err
		}
		field.t, err = gen.genTypeTL2(resolvedRefType)
		if err != nil {
			return err
		}
		if bl, isBool := field.t.trw.(*TypeRWBool); isBool && refField.IsOptional && !bl.isTL2Legacy {
			return refField.PRName.BeautifulError(fmt.Errorf("field with type \"bool\" can't be optional (use any maybe-like wrapper)"))
		}
		// tl1 boxed only for union
		_, isUnion := field.t.trw.(*TypeRWUnion)
		field.bare = !isUnion
		// for tl1 unknown params call
		for paramI := 0; paramI < len(field.t.NatParams); paramI++ {
			field.natArgs = append(field.natArgs, ActualNatArg{
				isArith:        true,
				isTL2FakeArith: true,
				Arith: tlast.Arithmetic{
					Res:  math.MaxUint32,
					Nums: []uint32{math.MaxUint32},
				},
			})
		}
		// generate tl2 legacy setters
		if strings.Contains(refField.CommentBefore, "tlgen:addLegacySetters") {
			maskNameFound, maskName := extractTLGenTag(refField.CommentBefore, "tlgen:addLegacySetters:name")
			maskBitFound, maskBit := extractTLGenTag(refField.CommentBefore, "tlgen:addLegacySetters:bit")
			if maskNameFound && maskBitFound {
				field.GenerateLegacySettersForTL2Name = maskName
				bitUint64, err := strconv.ParseUint(maskBit, 10, 32)
				if err != nil {
					return err
				}
				field.BitNumber = uint32(bitUint64)
			}
		}
		*fields = append(*fields, field)
	}
	return nil
}

func (gen *Gen2) genFunctionTL2(kernelType *TypeRWWrapper, comb *tlast.TL2Combinator) (wr *TypeRWWrapper, err error) {
	// set up wrapper
	kernelType.tl2Name = comb.FuncDecl.Name
	kernelType.tl2Origin = comb
	kernelType.tlTag = *comb.FuncDecl.ID

	// TODO: for tl1 meta
	kernelType.tlName = tlast.Name{
		Namespace: comb.FuncDecl.Name.Namespace,
		Name:      comb.FuncDecl.Name.Name,
	}

	kernelType.ns = gen.getNamespace(comb.FuncDecl.Name.Namespace)
	kernelType.ns.types = append(kernelType.ns.types, kernelType)
	kernelType.fileName = comb.FuncDecl.Name.String()
	kernelType.goLocalName, kernelType.goGlobalName = getCombinatorNames(*comb, "")

	functionType := TypeRWStruct{
		wr: kernelType,
	}
	kernelType.trw = &functionType

	err = gen.genFieldsTL2(
		ResolvedTL2References{
			ResolvedNats:  map[string]uint32{},
			ResolvedTypes: map[string]tlast.TL2TypeRef{},
		},
		&functionType.Fields,
		comb.FuncDecl.Arguments,
	)

	if err != nil {
		return nil, err
	}

	if comb.FuncDecl.ReturnType.IsAlias() {
		// for less diff with tl1 generation inplace type
		functionType.ResultType, err = gen.genTypeTL2(comb.FuncDecl.ReturnType.TypeAlias)
	} else {
		// set up wrapper for result
		functionType.ResultType = &TypeRWWrapper{
			gen: gen,
			ns:  gen.getNamespace(comb.FuncDecl.Name.Namespace),

			goGlobalName: kernelType.goGlobalName + "Result",
			goLocalName:  kernelType.goLocalName + "Result",

			fileName: kernelType.fileName,

			originateFromTL2: true,
			wantsTL2:         true,

			tl2IsResult: true,
			tl2Name: tlast.TL2TypeName{
				Namespace: kernelType.tl2Name.Namespace,
				Name:      kernelType.tl2Name.Name + "Result",
			},
			tl2Origin: kernelType.tl2Origin,
		}
		functionType.ResultType.ns.types = append(functionType.ResultType.ns.types, functionType.ResultType)

		gen.generatedTypes[functionType.ResultType.goGlobalName] = functionType.ResultType
		gen.generatedTypesList = append(gen.generatedTypesList, functionType.ResultType)

		err = gen.genTypeDeclarationTL2(
			functionType.ResultType,
			comb.FuncDecl.ReturnType,
			ResolvedTL2References{
				ResolvedNats:  map[string]uint32{},
				ResolvedTypes: map[string]tlast.TL2TypeRef{},
			},
			tlast.TL2TypeRef{
				SomeType: &tlast.TL2TypeApplication{
					Name: functionType.ResultType.tl2Name,
				},
			},
		)
	}

	return kernelType, err
}

func (gen *Gen2) genBracketTypeTL2(kernelType *TypeRWWrapper, br tlast.TL2BracketType) (*TypeRWWrapper, error) {
	var err error
	bracketType := TypeRWBrackets{
		wr: kernelType,
	}
	kernelType.trw = &bracketType
	kernelType.tl2IsBuiltinBrackets = true

	elementRef := &bracketType.element
	if br.IndexType != nil {
		if br.IndexType.IsNumber {
			bracketType.dynamicSize = false
			bracketType.size = br.IndexType.Number

			kernelType.goLocalName, kernelType.goGlobalName = "BuiltinTuple", "BuiltinTuple"
			kernelType.arguments = append(kernelType.arguments, ResolvedArgument{
				isNat:   true,
				isArith: true,
				Arith:   tlast.Arithmetic{Res: bracketType.size},
			})
		} else {
			bracketType.dictLike = true
			bracketType.dictKeyField = Field{goName: "Key"}
			bracketType.dictKeyField.t, err = gen.genTypeTL2(br.IndexType.Type)
			if err != nil {
				return nil, err
			}

			if pr, ok := bracketType.dictKeyField.t.trw.(*TypeRWPrimitive); ok && pr.goType == "string" {
				bracketType.dictKeyString = true
			}

			_, isUnion := bracketType.dictKeyField.t.trw.(*TypeRWUnion)
			kernelType.goLocalName, kernelType.goGlobalName = "BuiltinMap", "BuiltinMap"
			kernelType.arguments = append(kernelType.arguments, ResolvedArgument{
				isNat:   false,
				isArith: false,
				tip:     bracketType.dictKeyField.t,
				bare:    !isUnion,
			})
			elementRef = &bracketType.dictValueField
		}
	} else {
		bracketType.vectorLike = true
		kernelType.goLocalName, kernelType.goGlobalName = "BuiltinVector", "BuiltinVector"
	}

	elementRef.goName = "Value"
	elementRef.t, err = gen.genTypeTL2(br.ArrayType)
	if err != nil {
		return nil, err
	}

	_, isUnion := elementRef.t.trw.(*TypeRWUnion)
	kernelType.arguments = append(kernelType.arguments, ResolvedArgument{
		isNat:   false,
		isArith: false,
		tip:     elementRef.t,
		bare:    !isUnion,
	})

	kernelType.goLocalName += kernelType.wrapperNameTail()
	kernelType.goGlobalName += kernelType.wrapperNameTail()

	kernelType.ns = elementRef.t.ns
	kernelType.ns.types = append(kernelType.ns.types, kernelType)
	kernelType.fileName = elementRef.t.fileName

	return kernelType, nil
}

func (gen *Gen2) genBoolTL2(kernelType *TypeRWWrapper, isLegacy bool) (*TypeRWWrapper, error) {
	boolType := TypeRWBool{wr: kernelType}
	kernelType.trw = &boolType

	if isLegacy {
		kernelType.tl2Name = tlast.TL2TypeName{Name: "legacy_bool"}

		kernelType.fileName = "legacy_bool"
		kernelType.goGlobalName = "LegacyBool"
		kernelType.goLocalName = "LegacyBool"

		boolType.isTL2Legacy = true
	} else {
		kernelType.tl2Name = tlast.TL2TypeName{Name: "bool"}

		kernelType.fileName = "bool"
		kernelType.goGlobalName = "Bool"
		kernelType.goLocalName = "Bool"
	}

	kernelType.ns = gen.getNamespace("")
	kernelType.ns.types = append(kernelType.ns.types, kernelType)

	return kernelType, nil
}

func (gen *Gen2) genMaybeTL2(kernelType *TypeRWWrapper, comb *tlast.TL2Combinator, resolveMapping ResolvedTL2References) (*TypeRWWrapper, error) {
	// check can be maybe
	nonMaybeErr := fmt.Errorf("this type can't be maybe due to definition of maybe (ideal: @maybe maybe<t:type> = None | Some t;)")

	// {0?}<t:type> = {1?} | {2?};
	if len(comb.TypeDecl.TemplateArguments) != 1 ||
		!comb.TypeDecl.TemplateArguments[0].Category.IsType() ||
		!comb.TypeDecl.Type.IsUnionType ||
		len(comb.TypeDecl.Type.UnionType.Variants) != 2 {
		return nil, comb.TypeDecl.PRName.BeautifulError(nonMaybeErr)
	}
	variants := comb.TypeDecl.Type.UnionType.Variants
	emptyVariant, valueVariant := variants[0], variants[1]
	if emptyVariant.IsTypeAlias ||
		len(emptyVariant.Fields) != 0 {
		return nil, emptyVariant.PRName.BeautifulError(fmt.Errorf("first constructor must be empty"))
	}

	valueRef := valueVariant.TypeAlias
	if valueVariant.IsTypeAlias {
		// {0?}<{4?}:type> = {1?} | {2?} {4?};
		if valueVariant.TypeAlias.String() != comb.TypeDecl.TemplateArguments[0].Name {
			return nil, valueVariant.PRName.BeautifulError(fmt.Errorf("second constructor can be alias only to first type argument"))
		}
	} else {
		// {0?}<{4?}:type> = {1?} | {2?} {3?}:{4?};
		if len(valueVariant.Fields) == 0 {
			return nil, valueVariant.PRName.BeautifulError(fmt.Errorf("second constructor must be non-empty"))
		}
		valueRef = valueVariant.Fields[0].Type
		if len(valueVariant.Fields) != 1 ||
			valueVariant.Fields[0].IsOptional ||
			valueVariant.Fields[0].IsIgnored ||
			valueVariant.Fields[0].Type.IsBracket ||
			valueVariant.Fields[0].Type.SomeType == nil ||
			valueVariant.Fields[0].Type.SomeType.Name.String() != comb.TypeDecl.TemplateArguments[0].Name {
			return nil, comb.TypeDecl.PRName.BeautifulError(nonMaybeErr)
		}
	}

	var err error
	valueRef, err = resolveMapping.resolveRef(valueRef)
	if err != nil {
		return nil, err
	}

	maybeType := TypeRWMaybe{wr: kernelType}
	maybeType.element.goName = "Value"
	maybeType.element.t, err = gen.genTypeTL2(valueRef)

	if err != nil {
		return nil, err
	}

	kernelType.trw = &maybeType
	kernelType.tl2Name = comb.TypeDecl.Name
	kernelType.tl2Origin = comb

	elementRef := maybeType.element

	_, isUnion := elementRef.t.trw.(*TypeRWUnion)
	kernelType.arguments = append(kernelType.arguments, ResolvedArgument{
		isNat:   false,
		isArith: false,
		tip:     elementRef.t,
		bare:    !isUnion,
	})

	suffix := snakeToCamelCase(comb.TypeDecl.Name.Namespace + "_" + comb.TypeDecl.Name.Name)

	kernelType.goLocalName = snakeToCamelCase(elementRef.t.goGlobalName + ifString(!isUnion, suffix, "Boxed"+suffix))
	kernelType.goGlobalName = snakeToCamelCase(elementRef.t.goGlobalName + ifString(!isUnion, suffix, "Boxed"+suffix))

	kernelType.ns = elementRef.t.ns
	kernelType.ns.types = append(kernelType.ns.types, kernelType)
	kernelType.fileName = elementRef.t.fileName

	return kernelType, nil
}

func snakeToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	newS := ""
	for _, part := range parts {
		if len(part) > 0 {
			newS += strings.ToUpper(part[:1]) + part[1:]
		}
	}
	if newS == "" {
		return newS
	}
	return strings.ToUpper(newS[:1]) + newS[1:]
}

func getCombinatorNames(combinator tlast.TL2Combinator, argTail string) (localName string, globalName string) {
	tn := combinator.ReferenceName()
	if combinator.HasAnnotation(tl2Ext) {
		suffix := ""
		for _, argument := range combinator.TypeDecl.TemplateArguments {
			if argument.Category.IsUint32() {
				suffix += "_" + strings.ToUpper(argument.Name[:1]) + argument.Name[1:]
			}
		}
		tn.Name, _ = strings.CutSuffix(tn.Name, suffix)
	}
	return getTypeNames(tn, argTail)
}

func getTypeNames(tl2Name tlast.TL2TypeName, argTail string) (localName string, globalName string) {
	tName := tl2Name.Name
	tNs := tl2Name.Namespace
	return snakeToCamelCase(tName) + argTail, snakeToCamelCase(tNs+"_"+tName) + argTail
}

func getVariantNames(tl2Name tlast.TL2TypeName, constructors []tlast.TL2UnionConstructor, constructorId int, yetCreatedNames map[string]bool, argTail string) (namespace string, localName string, globalName string, tlName tlast.Name, fieldName string, err error) {
	// TODO: tl1 legacy...
	const Tag = "tlgen:tl1name"
	prefix := strings.ToLower(tl2Name.Name)
	suffix := strings.ToLower(tl2Name.Name)

	for _, unionConstructor := range constructors {
		localFound, localComment := extractTLGenTag(unionConstructor.CommentBefore, Tag)
		if localFound {
			_, tl1Name := "", localComment
			if strings.Contains(tl1Name, ".") {
				_, tl1Name, _ = strings.Cut(tl1Name, ".")
			}
			if !strings.HasPrefix(tl1Name, prefix) {
				prefix = ""
			}
			if !strings.HasSuffix(tl1Name, suffix) {
				suffix = ""
			}
		}
	}

	constructor := constructors[constructorId]
	found, comment := extractTLGenTag(constructor.CommentBefore, Tag)

	if found {
		tl1Namespace, tl1Name := "", comment
		if strings.Contains(tl1Name, ".") {
			tl1Namespace, tl1Name, _ = strings.Cut(tl1Name, ".")
		}
		namespace = tl1Namespace
		tlName = tlast.Name{Namespace: tl1Namespace, Name: tl1Name}
		localName, globalName = getTypeNames(
			tlast.TL2TypeName{
				Namespace: tl1Namespace,
				Name:      tl1Name,
			},
			argTail,
		)
		nameSuffix := tl1Name
		if prefix != "" && strings.HasPrefix(strings.ToLower(tl1Name), prefix) {
			nameSuffix = tl1Name[len(prefix):]
		} else if suffix != "" && strings.HasSuffix(strings.ToLower(tl1Name), suffix) {
			nameSuffix = tl1Name[:len(tl1Name)-len(suffix)]
		}
		fieldName = snakeToCamelCase(nameSuffix)
		if yetCreatedNames[fieldName] {
			fieldName = snakeToCamelCase(tl1Namespace + "_" + nameSuffix)
		}
		return
	} else {
		namespace = tl2Name.Namespace
		tlName = tlast.Name{Namespace: namespace, Name: tl2Name.Name + constructor.Name}
		localName, globalName = getTypeNames(
			tlast.TL2TypeName{
				Namespace: tl2Name.Namespace,
				Name:      tl2Name.Name + "_" + constructor.Name,
			},
			argTail,
		)
		fieldName = snakeToCamelCase(constructor.Name)
		return
	}
}

func (w *TypeRWWrapper) wrapperNameTail() (tail string) {
	b := strings.Builder{}
	for _, a := range w.arguments {
		if a.isNat {
			if a.isTL2FakeArith {
				b.WriteString("FakeUint32Max")
			} else {
				b.WriteString(strconv.FormatUint(uint64(a.Arith.Res), 10))
			}
		} else {
			b.WriteString(a.tip.goGlobalName)
		}
	}
	return b.String()
}

func (gen *Gen2) isTL1Ref(ref tlast.TL2TypeRef) bool {
	if ref.IsBracket {
		return false
	}
	if ref.SomeType == nil {
		return false
	}
	typeApplication := *ref.SomeType
	name := typeApplication.Name
	comb, ok := gen.tl2Combinators[name.String()]
	if !ok {
		return false
	}
	if comb.IsFunction {
		return false
	}
	return comb.HasAnnotation(tl1Ref)
}

// finds value for tags in comment with template {tag}:"{value}"
func extractTLGenTag(comment string, tag string) (found bool, value string) {
	index := strings.Index(comment, tag)
	if index == -1 {
		return
	}
	comment = comment[index+len(tag):]
	index = strings.Index(comment, ":")
	if index == -1 {
		return
	}
	comment = comment[index+1:]
	index = strings.Index(comment, "\"")
	if index == -1 {
		return
	}
	comment = comment[index+1:]
	index = strings.Index(comment, "\"")
	if index == -1 {
		return
	}
	found = true
	value = comment[:index]
	return
}
