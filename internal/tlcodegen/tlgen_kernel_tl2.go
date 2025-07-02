package tlcodegen

import (
	"fmt"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
	"sort"
	"strings"
)

type AstInfo struct {
	AllNamesSorted  []string
	AllConstructors map[string]*tlast.Combinator

	BuiltInConstructors map[tlast.Name]bool
	Types               map[tlast.Name][]*tlast.Combinator
	// for non builtins
	ConstructorToType map[tlast.Name]tlast.Name
	Functions         []*tlast.Combinator
}

func initAstInfo(allConstructors map[string]*tlast.Combinator) AstInfo {
	info := AstInfo{AllConstructors: allConstructors}
	allTypes := utils.Keys(info.AllConstructors)

	info.BuiltInConstructors = make(map[tlast.Name]bool)
	info.Types = make(map[tlast.Name][]*tlast.Combinator)
	info.ConstructorToType = make(map[tlast.Name]tlast.Name)

	sort.Slice(allTypes, func(i, j int) bool {
		hasNSi := strings.Contains(allTypes[i], ".")
		hasNSj := strings.Contains(allTypes[j], ".")
		if hasNSi == hasNSj {
			return strings.Compare(allTypes[i], allTypes[j]) <= 0
		} else {
			return hasNSj
		}
	})

	info.AllNamesSorted = allTypes
	for _, name := range info.AllNamesSorted {
		comb := info.AllConstructors[name]
		if comb.Builtin {
			info.BuiltInConstructors[comb.Construct.Name] = true
		} else {
			if comb.IsFunction {
				info.Functions = append(info.Functions, comb)
			} else {
				info.ConstructorToType[comb.Construct.Name] = comb.TypeDecl.Name
				info.Types[comb.TypeDecl.Name] = append(info.Types[comb.TypeDecl.Name], comb)
			}
		}
	}
	return info
}

func (ai AstInfo) TypeFromName(name tlast.Name) tlast.Name {
	if typ, ok := ai.ConstructorToType[name]; ok {
		return typ
	} else {
		return name
	}
}

func (gen *Gen2) MigrateToTL2() tlast.TL2File {
	file := tlast.TL2File{}
	info := initAstInfo(gen.allConstructors)

	// print for debug
	if false {
		fmt.Println("[")
		for _, typ := range info.AllNamesSorted {
			suffix := ""
			if info.AllConstructors[typ].Builtin {
				suffix = "(builtin)"
			} else if !gen.allConstructors[typ].IsFunction {
				suffix = "(" + gen.allConstructors[typ].TypeDecl.Name.String() + ")"
			}
			fmt.Println("\t", typ, suffix)
		}
		fmt.Println("]")
	}

	associatedWrappers := make(map[*tlast.Combinator][]*TypeRWWrapper)
	associatedCombinator := make(map[tlast.Name]*tlast.Combinator)

	for _, wrapper := range gen.generatedTypesList {
		for _, combinator := range wrapper.origTL {
			if combinator.IsFunction {
				continue
			}
			associatedCombinator[combinator.Construct.Name] = combinator
			if !combinator.Builtin {
				associatedCombinator[combinator.TypeDecl.Name] = combinator
			}
			associatedWrappers[combinator] = append(associatedWrappers[combinator], wrapper)
		}
	}

	var combinedTL tlast.TL = utils.Values(info.AllConstructors)
	natUsage := checkNatUsages(&combinedTL)

	typeNames := utils.Keys(info.Types)
	sort.Slice(typeNames, func(i, j int) bool {
		return compareNames(typeNames[i], typeNames[j])
	})

	lowerFirst := func(s string) string {
		if s == "" {
			return s
		}
		return strings.ToLower(s[:1]) + s[1:]
	}

	upperFirst := func(s string) string {
		if s == "" {
			return s
		}
		return strings.ToUpper(s[:1]) + s[1:]
	}

	var resolveType func(ref tlast.TypeRef, natIsConstant map[string]bool) (newRef tlast.TL2TypeRef)
	resolveType = func(ref tlast.TypeRef, natIsConstant map[string]bool) (newRef tlast.TL2TypeRef) {
		if ref.Type.String() == "#" {
			newRef.SomeType = &tlast.TL2TypeApplication{
				Name: tlast.TL2TypeName{
					Name: "uint32",
				},
			}
			return
		}
		comb := associatedCombinator[ref.Type]
		if comb != nil {
			wrapper := associatedWrappers[comb][0]
			tname := info.TypeFromName(ref.Type)
			if _, ok := wrapper.trw.(*TypeRWBrackets); ok {
				newRef.IsBracket = true
				newRef.BracketType = new(tlast.TL2BracketType)
				arrayIndex := 0
				if len(ref.Args) == 2 {
					arrayIndex = 1
					indexType := ref.Args[0]
					if indexType.IsArith {
						newRef.BracketType.IndexType = new(tlast.TL2TypeArgument)
						newRef.BracketType.IndexType.IsNumber = true
						newRef.BracketType.IndexType.Number = indexType.Arith.Res
					} else if natIsConstant[indexType.String()] {
						newRef.BracketType.IndexType = new(tlast.TL2TypeArgument)
						newRef.BracketType.IndexType.Type.SomeType = &tlast.TL2TypeApplication{
							Name: tlast.TL2TypeName{
								Name: lowerFirst(indexType.T.Type.Name),
							},
						}
					}
				}
				newRef.BracketType.ArrayType = resolveType(ref.Args[arrayIndex].T, natIsConstant)
			} else {
				newRef.SomeType = &tlast.TL2TypeApplication{
					Name: tlast.TL2TypeName{
						Namespace: tname.Namespace,
						Name:      lowerFirst(tname.Name),
					},
				}
				for i, arg := range ref.Args {
					if len(natUsage.GetInfluencedNatFieldsToTemplate(tname, i)) > 0 {
						continue
					}
					newArg := tlast.TL2TypeArgument{}
					if arg.IsArith {
						newArg.IsNumber = true
						newArg.Number = arg.Arith.Res
					} else {
						newArg.Type = resolveType(arg.T, natIsConstant)
					}
					newRef.SomeType.Arguments = append(newRef.SomeType.Arguments, newArg)
				}
			}
		} else {
			newRef.SomeType = &tlast.TL2TypeApplication{
				Name: tlast.TL2TypeName{
					Name: lowerFirst(ref.Type.Name),
				},
			}
		}
		return
	}

	addFields := func(oldFields []tlast.Field, natIsConstant map[string]bool) (newFields []tlast.TL2Field) {
		for _, field := range oldFields {
			newField := tlast.TL2Field{}
			newField.Name = lowerFirst(field.FieldName)
			if newField.Name == "" {
				newField.Name = "_"
				newField.IsIgnored = true
			}
			if field.Mask != nil {
				newField.IsOptional = true
				if !field.IsRepeated {
					comb := associatedCombinator[field.FieldType.Type]
					if comb != nil {
						wrapper := associatedWrappers[comb][0]
						if _, ok := wrapper.trw.(*TypeRWBool); ok {
							newField.IsOptional = false
							newField.Type.SomeType = &tlast.TL2TypeApplication{
								Name: tlast.TL2TypeName{Name: "maybe"},
								Arguments: []tlast.TL2TypeArgument{
									{
										Type: tlast.TL2TypeRef{
											SomeType: &tlast.TL2TypeApplication{
												Name: tlast.TL2TypeName{Name: "bool"},
											},
										},
									},
								},
							}
							newFields = append(newFields, newField)
							continue
						}
					}
				}
			}
			calculatingType := &newField.Type
			if field.IsRepeated {
				newField.Type.IsBracket = true
				newField.Type.BracketType = new(tlast.TL2BracketType)
				calculatingType = &newField.Type.BracketType.ArrayType
				if !field.ScaleRepeat.ExplicitScale {
					newField.Type.BracketType.IndexType = nil
				} else {
					newField.Type.BracketType.IndexType = new(tlast.TL2TypeArgument)
					if field.ScaleRepeat.Scale.IsArith {
						newField.Type.BracketType.IndexType.IsNumber = true
						newField.Type.BracketType.IndexType.Number = field.ScaleRepeat.Scale.Arith.Res
					} else {
						newField.Type.BracketType.IndexType.Type.SomeType = &tlast.TL2TypeApplication{
							Name: tlast.TL2TypeName{Name: lowerFirst(field.ScaleRepeat.Scale.Scale)},
						}
					}
				}
			}
			*calculatingType = resolveType(field.FieldType, natIsConstant)
			newFields = append(newFields, newField)
		}
		return
	}

	for _, name := range typeNames {
		combinators := info.Types[name]
		if len(combinators) == 0 {
			continue
		}
		combinator0 := combinators[0]
		tl2Combinator := tlast.TL2Combinator{IsFunction: false}
		// init name
		tl2Combinator.TypeDecl.Name.Namespace = name.Namespace
		tl2Combinator.TypeDecl.Name.Name = lowerFirst(name.Name)
		// init templates
		natTemplates := make(map[string]int)
		natIsConstant := make(map[string]bool)
		generic := make(map[string]bool)

		for i, argument := range combinator0.TemplateArguments {
			category := "type"
			if argument.IsNat {
				category = "uint32"
				natTemplates[argument.FieldName] = i
				if len(natUsage.GetInfluencedNatFieldsToTemplate(combinator0.TypeDecl.Name, i)) > 0 {
					continue
				}
			} else {
				generic[argument.FieldName] = true
			}

			natIsConstant[argument.FieldName] = true
			tl2Combinator.TypeDecl.TemplateArguments = append(tl2Combinator.TypeDecl.TemplateArguments,
				tlast.TL2TypeTemplate{
					Name:     lowerFirst(argument.FieldName),
					Category: tlast.TL2TypeCategory(category),
				},
			)
		}
		//// init magic
		//if combinator0.Construct.ID != nil {
		//	tl2Combinator.TypeDecl.ID = new(uint32)
		//	*tl2Combinator.TypeDecl.ID = *combinator0.Construct.ID
		//}

		if len(combinators) == 1 {
			if len(combinator0.Fields) == 1 && combinator0.Fields[0].FieldName == "" {
				tl2Combinator.TypeDecl.Type.TypeAlias = resolveType(combinator0.Fields[0].FieldType, natIsConstant)
			} else {
				tl2Combinator.TypeDecl.Type.IsConstructorFields = true
				tl2Combinator.TypeDecl.Type.ConstructorFields = addFields(combinator0.Fields, natIsConstant)
			}
		} else {
			tl2Combinator.TypeDecl.Type.IsUnionType = true
			for _, combinator := range combinators {
				newVariant := tlast.TL2UnionTypeVariant{}
				if len(combinator.Fields) == 1 && combinator.Fields[0].FieldName == "" {
					newVariant.TypeAlias = resolveType(combinator.Fields[0].FieldType, natIsConstant)
				} else {
					newVariant.IsConstructor = true
					newVariant.Constructor = tlast.TL2UnionConstructor{
						Name: upperFirst(combinator.Construct.Name.Name),
					}
					newVariant.Constructor.Fields = addFields(combinator.Fields, natIsConstant)
				}
				tl2Combinator.TypeDecl.Type.UnionType.Variants = append(tl2Combinator.TypeDecl.Type.UnionType.Variants, newVariant)
			}
		}

		file.Combinators = append(file.Combinators, tl2Combinator)
	}

	for _, function := range info.Functions {
		tl2Combinator := tlast.TL2Combinator{IsFunction: true}
		// add modifiers
		for _, modifier := range function.Modifiers {
			tl2Combinator.Annotations = append(tl2Combinator.Annotations,
				tlast.TL2Annotation{
					Name: modifier.Name,
				},
			)
		}
		// add name
		tl2Combinator.FuncDecl.Name = tlast.TL2TypeName{
			Namespace: function.Construct.Name.Namespace,
			Name:      lowerFirst(function.Construct.Name.Name),
		}
		// add magic
		if function.Construct.ID != nil {
			tl2Combinator.FuncDecl.ID = new(uint32)
			*tl2Combinator.FuncDecl.ID = *function.Construct.ID
		}
		// add arguments
		tl2Combinator.FuncDecl.Arguments = addFields(function.Fields, make(map[string]bool))
		// add return type
		tl2Combinator.FuncDecl.ReturnType = tlast.TL2TypeDefinition{
			TypeAlias: resolveType(function.FuncDecl, make(map[string]bool)),
		}
		file.Combinators = append(file.Combinators, tl2Combinator)
	}

	printDebugInfo(associatedWrappers, natUsage, info)

	return file
}

func printDebugInfo(associatedWrappers map[*tlast.Combinator][]*TypeRWWrapper, natUsage NatUsagesInfo, info AstInfo) {
	// print for debug
	if false {
		for combinator, wrappers := range associatedWrappers {
			fmt.Println(">>>>>")
			fmt.Println(combinator)
			fmt.Printf("[info] isBuiltin:%t, isFunction:%t\n", combinator.Builtin, combinator.IsFunction)
			fmt.Println("[")
			for _, wrapper := range wrappers {
				fmt.Print("\t")
				fmt.Print(wrapper.goGlobalName)
				fmt.Println(",")
			}
			fmt.Println("]")
		}
	}

	// print for debug
	if false {
		for name, m := range natUsage.TypeNameAndArgToAffectingConstants {
			fmt.Printf(">>>>>>\n%s: \n[\n", name)
			for index, cs := range m {
				fmt.Printf("\t%d: [", index)
				for c := range cs {
					fmt.Printf("%d, ", c)
				}
				fmt.Printf("],\n")
			}
			fmt.Print("]\n")
		}
	}

	// print for debug
	if false {
		var tl tlast.TL = utils.Values(info.AllConstructors)
		sort.Slice(tl, func(i, j int) bool {
			return compareNames(tl[i].Construct.Name, tl[j].Construct.Name)
		})
		fmt.Println(tl)
	}

	// print for debug
	if false {
		for tp, args := range natUsage.TypeArgumentToArraySizeReference {
			fmt.Printf(">>>>>>\n%s:\n[\n", tp.String())
			for arg, refs := range args {
				fmt.Printf("\t%d: [\n", arg)
				for ref, fields := range refs {
					for field, path := range fields {
						pathStr := ""
						for _, edge := range path {
							pathStr += fmt.Sprintf(" <- ((%s, %d), %d)", edge.Type, edge.ArgIndex, edge.FieldIndex)
						}
						fmt.Printf("\t\t(\"%s\", %d)%s,\n", ref.String(), field, pathStr)
					}
				}
				fmt.Println("\t],")
			}
			fmt.Println("]")
		}
	}

	// print for debug
	if false {
		for tp, args := range natUsage.CombinatorsNatFieldsToArraySizeReference {
			fmt.Printf("<<<<<<\n%s:\n[\n", tp.String())
			for arg, refs := range args {
				fmt.Printf("\t%d: [\n", arg)
				for ref, fields := range refs {
					for field := range fields {
						fmt.Printf("\t\t(\"%s\", %d),\n", ref.String(), field)
					}
				}
				fmt.Println("\t],")
			}
			fmt.Println("]")
		}
	}

	// print for debug
	if false {
		fmt.Println("--------")

		for comb, fields := range natUsage.CombinatorsNatFieldIndexToBitsUsed {
			for field, bits := range fields {
				if len(bits) == 0 {
					continue
				}
				sizeRefs := natUsage.GetArraySizeReferenceForField(comb, field)
				if len(sizeRefs) != 0 {
					fmt.Printf("WARNING: %s's field #%d is used as a field mask and as array size !!!!!\n", comb.String(), field)
				}
			}
		}

		fmt.Println("--------")

		for typ, args := range natUsage.TypeNameAndArgIndexToBitsUsedByNat {
			for arg, bits := range args {
				if len(bits) == 0 {
					continue
				}
				sizeRefs := natUsage.GetArgumentUsagesAsSize(typ, arg)
				if len(sizeRefs) != 0 {
					fmt.Printf("WARNING: %s's template argument #%d is used as a field mask and as array size !!!!!\n", typ.String(), arg)
				}
			}
		}
	}
}

func compareNames(name1, name2 tlast.Name) bool {
	emptyI, emptyJ := name1.Namespace == "", name2.Namespace == ""
	if emptyI == emptyJ {
		return strings.Compare(name1.String(), name2.String()) <= 0
	} else {
		return emptyI
	}
}
