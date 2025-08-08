package tlcodegen

import (
	"fmt"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
	"golang.org/x/exp/slices"
	"os"
	"path/filepath"
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

type MigratingPartInfo struct {
	File string

	CombinatorsToMigrate   []*tlast.Combinator
	CombinatorsToReference []*tlast.Combinator
}

func (gen *Gen2) SplitMigratingTypes() ([]MigratingPartInfo, error) {
	if !gen.options.TL2MigrateByNamespaces {
		allConstructors := utils.Values(gen.allConstructors)
		sort.Slice(allConstructors, func(i, j int) bool {
			return allConstructors[i].OriginalOrderIndex < allConstructors[j].OriginalOrderIndex
		})

		return []MigratingPartInfo{
			{
				File:                 gen.options.TL2MigrationFile,
				CombinatorsToMigrate: allConstructors,
			},
		}, nil
	}

	typesInfo := processCombinators(gen.allConstructors)

	namespaceToReductions := make(map[string]map[string]*TypeReduction)
	namespaceToTypes := make(map[string][]tlast.Name)
	allReductions := make(map[string]*TypeReduction)
	reductionToNamespaceUsages := make(map[string]map[string]bool)
	reductionToNamespaceOrigin := make(map[string]string)

	allTypes := utils.Keys(typesInfo.Types)
	sort.Slice(allTypes, func(i, j int) bool {
		return compareNames(allTypes[i], allTypes[j])
	})

	for _, name := range allTypes {
		namespaceToTypes[name.Namespace] = append(namespaceToTypes[name.Namespace], name)
	}

	allNamespaces := utils.Keys(namespaceToTypes)
	sort.Strings(allNamespaces)

	for _, namespace := range allNamespaces {
		typeReductions := make(map[string]*TypeReduction)
		visitedTypes := make(map[TypeName]bool)

		for _, name := range namespaceToTypes[namespace] {
			comb := typesInfo.Types[name]
			if len(comb.TypeArguments) != 0 {
				continue
			}
			reduce(
				TypeReduction{
					IsType: true,
					Type:   typesInfo.Types[comb.Name],
				},
				&typeReductions,
				&typesInfo.Types,
				&typesInfo.Constructors,
				&visitedTypes,
			)
			if len(comb.Constructors) == 1 && comb.Constructors[0].Result != nil {
				resultRed := typesInfo.ResultTypeReduction(&TypeReduction{
					IsType:      false,
					Constructor: comb.Constructors[0],
				})
				reduce(
					*resultRed.Type,
					&typeReductions,
					&typesInfo.Types,
					&typesInfo.Constructors,
					&visitedTypes,
				)
			}
		}

		namespaceToReductions[namespace] = typeReductions
		for s, reduction := range typeReductions {
			allReductions[s] = reduction
		}
	}

	// view results
	if false {
		file, _ := os.Create("./reductions.txt")

		namespaces := utils.Keys(namespaceToReductions)
		sort.Strings(namespaces)

		for _, namespace := range namespaces {
			fmt.Fprintf(file, ">>> %s\n", namespace)

			reductions := utils.Keys(namespaceToReductions[namespace])
			sort.Strings(reductions)

			for _, reduction := range reductions {
				fmt.Fprintf(file, "%s\n", reduction)
			}
		}

	}

	for ns, reds := range namespaceToReductions {
		for red, redType := range reds {
			if _, ok := reductionToNamespaceUsages[red]; !ok {
				reductionToNamespaceUsages[red] = make(map[string]bool)
			}
			reductionToNamespaceUsages[red][ns] = true
			reductionToNamespaceOrigin[red] = redType.ReferenceName().Namespace
		}
	}

	// view results
	if false {
		file2, _ := os.Create("./common_reductions.txt")
		filter := true

		if filter {
			fmt.Fprintf(file2, "Filter to show only reductions in multiple namespaces\n")
		}

		reds := utils.Keys(reductionToNamespaceUsages)
		sort.Strings(reds)

		for _, red := range reds {
			nss := utils.Keys(reductionToNamespaceUsages[red])
			if filter {
				if len(nss) < 2 {
					continue
				}
			}

			sort.Strings(nss)

			fmt.Fprintf(file2, ">>> %s: [", red)
			for _, ns := range nss {
				if ns == "" {
					ns = "__empty"
				}
				fmt.Fprintf(file2, "%s,", ns)
			}

			fmt.Fprintf(file2, "]\n")
		}
	}

	filter := prepareNameFilter(gen.options.TL2MigratingWhitelist)
	referencingConstructors := make(map[string]map[string]*tlast.Combinator)
	migratingConstructors := make(map[string]map[string]*tlast.Combinator)

	for s, combinator := range gen.allConstructors {
		consideredName := combinator.TypeDecl.Name
		if combinator.IsFunction || combinator.Builtin {
			consideredName = combinator.Construct.Name
		}
		if inNameFilter(consideredName, filter) {
			if _, ok := migratingConstructors[consideredName.Namespace]; !ok {
				migratingConstructors[consideredName.Namespace] = make(map[string]*tlast.Combinator)
			}
			migratingConstructors[consideredName.Namespace][s] = combinator
		}
	}

	migratingNamespaces := utils.Keys(migratingConstructors)
	sort.Strings(migratingNamespaces)

	for _, ns := range migratingNamespaces {
		for red, redType := range namespaceToReductions[ns] {
			origin := reductionToNamespaceOrigin[red]
			if _, ok := migratingConstructors[origin]; !ok {
				refName := redType.ReferenceName().String()
				if refConstructor, ok := gen.allConstructors[refName]; ok {
					if refConstructor.TypeDecl.Name.Namespace != origin {
						// for constructor with another namespace
						continue
					}
					if _, ok := referencingConstructors[origin]; !ok {
						referencingConstructors[origin] = make(map[string]*tlast.Combinator)
					}
					referencingConstructors[origin][refName] = refConstructor

				}
			}
		}
	}

	absentNamespaceDependencies := make(map[string]map[string]bool)

	for _, ns := range allNamespaces {
		if _, ok := migratingConstructors[ns]; ok {
			continue
		}
		for red := range namespaceToReductions[ns] {
			origin := reductionToNamespaceOrigin[red]
			if _, ok := migratingConstructors[origin]; ok {
				if _, ok := absentNamespaceDependencies[origin]; !ok {
					absentNamespaceDependencies[origin] = make(map[string]bool)
				}
				absentNamespaceDependencies[origin][ns] = true
			}
		}
	}

	if len(absentNamespaceDependencies) != 0 {
		errorText := strings.Builder{}
		errorText.WriteString("can't migrate current set of namespaces, cause non-migrating namespaces below depends on them:\n{\n")

		keys := utils.Keys(absentNamespaceDependencies)
		sort.Strings(keys)

		for i, key := range keys {
			deps := utils.Keys(absentNamespaceDependencies[key])
			sort.Strings(deps)

			errorText.WriteString(fmt.Sprintf("\t\"%s\": [", key))
			for j, dep := range deps {
				errorText.WriteString(fmt.Sprintf("\"%s\"", dep))
				if j != len(deps)-1 {
					errorText.WriteString(",")
				}
			}
			errorText.WriteString("]")
			if i != len(keys)-1 {
				errorText.WriteString(",")
			}
			errorText.WriteString("\n")
		}

		errorText.WriteString("}\n(add depended namespaces to whitelist or remove unnecessary ones)")
		return nil, fmt.Errorf("%s", errorText.String())
	}

	affectedNamespaces := utils.SliceToSet(utils.Keys(referencingConstructors))
	utils.AppendMap(
		utils.SliceToSet(utils.Keys(migratingConstructors)),
		&affectedNamespaces,
	)

	results := make([]MigratingPartInfo, 0)

	affectedNamespacesList := utils.Keys(affectedNamespaces)
	sort.Strings(affectedNamespacesList)

	for _, ns := range affectedNamespacesList {
		nsFileName := ns
		if nsFileName == "" {
			nsFileName = "__common_namespace"
		}
		newInfo := MigratingPartInfo{File: filepath.Join(gen.options.TL2MigrationFile, "namespaces", nsFileName+".tl2")}
		nsRefConstructors := referencingConstructors[ns]

		constructors := utils.Keys(nsRefConstructors)
		slices.SortFunc(constructors, func(a, b string) int {
			return nsRefConstructors[a].OriginalOrderIndex - nsRefConstructors[b].OriginalOrderIndex
		})

		for _, constructor := range constructors {
			newInfo.CombinatorsToReference = append(newInfo.CombinatorsToReference, nsRefConstructors[constructor])
		}

		nsMigratingConstructors := migratingConstructors[ns]
		constructors = utils.Keys(nsMigratingConstructors)
		slices.SortFunc(constructors, func(a, b string) int {
			return nsMigratingConstructors[a].OriginalOrderIndex - nsMigratingConstructors[b].OriginalOrderIndex
		})

		for _, constructor := range constructors {
			newInfo.CombinatorsToMigrate = append(newInfo.CombinatorsToMigrate, nsMigratingConstructors[constructor])
		}

		results = append(results, newInfo)
	}

	return results, nil
}

type FileToWrite struct {
	Path string
	Ast  tlast.TL2File
}

func (gen *Gen2) MigrateToTL2(prevState []FileToWrite) (newState []FileToWrite, _ error) {
	parts, err := gen.SplitMigratingTypes()
	if err != nil {
		return nil, err
	}

	prevTypeTL2Info := getTypesInfoFromTL2State(prevState)

	info := initAstInfo(gen.allConstructors)
	associatedWrappers := make(map[*tlast.Combinator][]*TypeRWWrapper)
	associatedCombinator := make(map[tlast.Name]*tlast.Combinator)

	for _, wrapper := range gen.generatedTypesList {
		for _, combinator := range wrapper.origTL {
			if combinator.IsFunction {
				continue
			}
			associatedWrappers[combinator] = append(associatedWrappers[combinator], wrapper)
		}
	}

	for _, combinator := range gen.allConstructors {
		// constructor name
		associatedCombinator[combinator.Construct.Name] = combinator
		// type name if presented
		if !combinator.Builtin {
			associatedCombinator[combinator.TypeDecl.Name] = combinator
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

	var resolveType func(ref tlast.TypeRef, natIsConstant map[string]bool, natTemples map[string]bool) (newRef tlast.TL2TypeRef)
	resolveType = func(ref tlast.TypeRef, natIsConstant map[string]bool, natTemples map[string]bool) (newRef tlast.TL2TypeRef) {
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
			tname := info.TypeFromName(ref.Type)

			// check bracket case
			isBracket := false
			if len(associatedWrappers[comb]) > 0 {
				wrapper := associatedWrappers[comb][0]
				_, isBracket = wrapper.trw.(*TypeRWBrackets)
			} else {
				isBracket = comb.Builtin && (comb.Construct.Name.String() == BuiltinTupleName ||
					comb.Construct.Name.String() == BuiltinVectorName)
			}

			// check bool case
			isBool := false
			if len(associatedWrappers[comb]) > 0 {
				wrapper := associatedWrappers[comb][0]
				_, isBool = wrapper.trw.(*TypeRWBool)
			}

			if isBracket {
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
				newRef.BracketType.ArrayType = resolveType(ref.Args[arrayIndex].T, natIsConstant, natTemples)
			} else if isBool {
				newRef.SomeType = &tlast.TL2TypeApplication{
					Name: tlast.TL2TypeName{
						Namespace: "",
						Name:      "bool",
					},
				}
			} else {
				newRef.SomeType = &tlast.TL2TypeApplication{
					Name: tlast.TL2TypeName{
						Namespace: tname.Namespace,
						Name:      lowerFirst(tname.Name),
					},
				}
				for i, arg := range ref.Args {
					if arg.IsArith || natTemples[arg.String()] {
						if arg.IsArith || natIsConstant[arg.String()] {
							newRef.SomeType.Name.Name += "_" + upperFirst(associatedCombinator[ref.Type].TemplateArguments[i].FieldName)
						} else {
							continue
						}
					}
					newArg := tlast.TL2TypeArgument{}
					if arg.IsArith {
						newArg.IsNumber = true
						newArg.Number = arg.Arith.Res
					} else {
						newArg.Type = resolveType(arg.T, natIsConstant, natTemples)
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

	addFields := func(oldFields []tlast.Field, natIsConstant map[string]bool, natTemplates map[string]bool) (newFields []tlast.TL2Field) {
		originalNatTemplates := utils.CopyMap(natTemplates)
		natTemplates = utils.CopyMap(natTemplates)
		for _, field := range oldFields {
			newField := tlast.TL2Field{}
			appendLegacySetterComment := false
			// name
			newField.Name = lowerFirst(field.FieldName)
			if newField.Name == "" {
				newField.Name = "_"
				newField.IsIgnored = true
			}
			// if it has field mask
			isTrue := field.FieldType.Type.String() == "true" || field.FieldType.Type.String() == "True"
			if field.Mask != nil {
				// local nat dependencies convert to optional (expect x:n.0?Bool)
				if !originalNatTemplates[field.Mask.MaskName] {
					newField.IsOptional = true
					if !field.IsRepeated {
						comb := associatedCombinator[field.FieldType.Type]
						if comb != nil && len(associatedWrappers[comb]) > 0 {
							wrapper := associatedWrappers[comb][0]
							if _, ok := wrapper.trw.(*TypeRWBool); ok {
								newField.IsOptional = true
								newField.Type.SomeType = &tlast.TL2TypeApplication{
									Name: tlast.TL2TypeName{Name: "legacy_bool"},
								}
								newFields = append(newFields, newField)
								continue
							}
						}
					}
				}
				// convert x:n.0?true -> x:bool
				if field.FieldType.Bare && isTrue {
					newField.IsOptional = false
					newField.Type.SomeType = &tlast.TL2TypeApplication{
						Name: tlast.TL2TypeName{Name: "bool"},
					}
					newFields = append(newFields, newField)
					continue
				}
				// outer (non-local) nat dependencies add comment for legacy Set/IsSet
				if originalNatTemplates[field.Mask.MaskName] && !natIsConstant[field.Mask.MaskName] {
					appendLegacySetterComment = true
				}
			}
			calculatingType := &newField.Type
			// if is repeated
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
			// calculate type
			*calculatingType = resolveType(field.FieldType, natIsConstant, natTemplates)
			// advance context
			if field.FieldType.String() == "#" {
				natTemplates[field.FieldName] = true
			}
			// add comment
			newField.CommentBefore = appendComment(
				newField.CommentBefore,
				field.CommentBefore,
			)
			newField.CommentBefore = appendComment(
				newField.CommentBefore,
				field.CommentRight,
			)
			if appendLegacySetterComment {
				newField.CommentBefore = appendComment(
					newField.CommentBefore,
					fmt.Sprintf("// tlgen:addLegacySetters:name:\"%s\"", field.Mask.MaskName),
				)
				newField.CommentBefore = appendComment(
					newField.CommentBefore,
					fmt.Sprintf("// tlgen:addLegacySetters:bit:\"%d\"", field.Mask.BitNumber),
				)
			}
			newFields = append(newFields, newField)
		}
		return
	}

	getOrder := func(name tlast.Name) int {
		combinators := info.Types[name]
		minIndex := combinators[0].OriginalOrderIndex
		for _, combinator := range combinators {
			if combinator.OriginalOrderIndex < minIndex {
				minIndex = combinator.OriginalOrderIndex
			}
		}
		return minIndex
	}

	for _, part := range parts {
		file := tlast.TL2File{}

		localConstructors := make(map[string]*tlast.Combinator)
		referencingConstructors := make(map[string]bool)

		for _, combinator := range part.CombinatorsToMigrate {
			localConstructors[combinator.Construct.Name.String()] = combinator
		}
		for _, combinator := range part.CombinatorsToReference {
			localConstructors[combinator.Construct.Name.String()] = combinator
			referencingConstructors[combinator.Construct.Name.String()] = true
		}

		localInfo := initAstInfo(localConstructors)
		localTypeNames := utils.Keys(localInfo.Types)
		sort.Slice(localTypeNames, func(i, j int) bool {
			return compareNames(localTypeNames[i], localTypeNames[j])
		})

		originalOrderTypeNames := slices.Clone(localTypeNames)
		sort.Slice(originalOrderTypeNames, func(i, j int) bool {
			return getOrder(originalOrderTypeNames[i]) < getOrder(originalOrderTypeNames[j])
		})

		for _, name := range originalOrderTypeNames {
			combinators := info.Types[name]
			if len(combinators) == 0 {
				continue
			}
			sort.Slice(combinators, func(i, j int) bool {
				return combinators[i].OriginalOrderIndex <= combinators[j].OriginalOrderIndex
			})
			combinator0 := combinators[0]

			isMaybe := false
			// work with legacy types
			if len(associatedWrappers[combinator0]) > 0 {
				// remove all alterations of bool
				wrapper := associatedWrappers[combinator0][0]
				if _, ok := wrapper.trw.(*TypeRWBool); ok {
					continue
				}
				if _, ok := wrapper.trw.(*TypeRWMaybe); ok {
					isMaybe = true
				}
			}

			// basic info
			baseName := tlast.TL2TypeName{Namespace: name.Namespace, Name: lowerFirst(name.Name)}
			isRef := referencingConstructors[combinator0.Construct.Name.String()]

			combinatorConstableArgs := make(map[int]bool)
			minimalConstableArgsSubSet := make(map[int]bool)

			for i, argument := range combinator0.TemplateArguments {
				if argument.IsNat {
					influencedNats := natUsage.GetInfluencedNatFieldsToTemplate(combinator0.TypeDecl.Name, i)

					usedAsConst := len(natUsage.GetConstantsPassingThroughArgument(name, i)) > 0
					usedAsUnknownValue := len(influencedNats) > 0
					if usedAsConst {
						combinatorConstableArgs[i] = true
					}
					// it used for constants only
					if !usedAsUnknownValue && usedAsConst {
						minimalConstableArgsSubSet[i] = true
					}
				}
			}

			allConstableArgsSubSets := utils.SetSubSets(combinatorConstableArgs)
			allConstableArgsSubSets = utils.FilterSlice(
				allConstableArgsSubSets,
				func(m map[int]bool) bool {
					for i := range minimalConstableArgsSubSet {
						if !m[i] {
							return false
						}
					}
					return true
				},
			)

			// add from previous state
			for _, combinator := range prevTypeTL2Info[baseName] {
				inheritedSubSet := make(map[int]bool)
				for i, argument := range combinator.TypeDecl.TemplateArguments {
					if argument.Category.IsUint32() {
						inheritedSubSet[i] = true
					}
				}
				inheritedSubSetList := utils.Keys(inheritedSubSet)
				sort.Ints(inheritedSubSetList)
				if !slices.ContainsFunc(allConstableArgsSubSets, func(m map[int]bool) bool {
					l := utils.Keys(m)
					sort.Ints(l)
					return slices.Equal(inheritedSubSetList, l)
				}) {
					allConstableArgsSubSets = append(allConstableArgsSubSets, inheritedSubSet)
				}
			}

			// sort them
			sort.Slice(allConstableArgsSubSets, func(i, j int) bool {
				if len(allConstableArgsSubSets[i]) == len(allConstableArgsSubSets[j]) {
					iList := utils.Keys(allConstableArgsSubSets[i])
					sort.Ints(iList)
					jList := utils.Keys(allConstableArgsSubSets[j])
					sort.Ints(jList)
					return slices.Compare(iList, jList) <= 0
				}
				return len(allConstableArgsSubSets[i]) <= len(allConstableArgsSubSets[j])
			})

			for _, setOfConstantNatArgs := range allConstableArgsSubSets {
				tl2Combinator := tlast.TL2Combinator{IsFunction: false}
				// init name
				tl2Combinator.TypeDecl.Name = baseName

				constantArgs := utils.Keys(setOfConstantNatArgs)
				sort.Ints(constantArgs)
				for _, arg := range constantArgs {
					tl2Combinator.TypeDecl.Name.Name += "_" + upperFirst(combinator0.TypeDecl.Arguments[arg])
				}
				// add annotations
				if isRef {
					tl2Combinator.Annotations = append(tl2Combinator.Annotations, tlast.TL2Annotation{Name: tl1Ref})
				} else {
					// only tl2 special info
					if isMaybe {
						tl2Combinator.Annotations = append(tl2Combinator.Annotations, tlast.TL2Annotation{Name: tl2Maybe})
					}
				}
				if len(constantArgs) != 0 {
					tl2Combinator.Annotations = append(tl2Combinator.Annotations, tlast.TL2Annotation{Name: tl2Ext})
				}

				// init templates
				natTemplates := make(map[string]bool)
				natIsConstant := make(map[string]bool)
				generic := make(map[string]bool)

				for i, argument := range combinator0.TemplateArguments {
					category := "type"
					if argument.IsNat {
						category = "uint32"
						natTemplates[argument.FieldName] = true
						if !setOfConstantNatArgs[i] {
							continue
						}
						natIsConstant[argument.FieldName] = true
					} else {
						generic[argument.FieldName] = true
					}

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

				if !isRef {
					if len(combinators) == 1 {
						if len(combinator0.Fields) == 1 && combinator0.Fields[0].FieldName == "" {
							tl2Combinator.TypeDecl.Type.TypeAlias = resolveType(combinator0.Fields[0].FieldType, natIsConstant, natTemplates)
						} else {
							tl2Combinator.TypeDecl.Type.IsConstructorFields = true
							tl2Combinator.TypeDecl.Type.ConstructorFields = addFields(combinator0.Fields, natIsConstant, natTemplates)
						}
					} else {
						tl2Combinator.TypeDecl.Type.IsUnionType = true
						for i, combinator := range combinators {
							newVariant := tlast.TL2UnionConstructor{}
							// add original comment
							if i != 0 {
								newVariant.CommentBefore = appendComment(
									newVariant.CommentBefore,
									combinator.CommentBefore,
								)
							}

							newVariant.CommentBefore = appendComment(
								newVariant.CommentBefore,
								combinator.CommentRight,
							)
							newVariant.Name = upperFirst(combinator.Construct.Name.Name)
							// name convert
							{
								writeOriginalName := false
								if combinator.Construct.Name.Namespace != combinator.TypeDecl.Name.Namespace {
									writeOriginalName = true
								} else {
									typeName := combinator.TypeDecl.Name.Name
									constructorName := combinator.Construct.Name.Name

									if strings.HasPrefix(strings.ToLower(constructorName), strings.ToLower(typeName)) {
										nameAfterPrefix := constructorName[len(typeName):]
										if len(nameAfterPrefix) > 0 {
											firstLetter := strings.ToLower(nameAfterPrefix)[0]
											if 'a' <= firstLetter && firstLetter <= 'z' {
												newVariant.Name = upperFirst(nameAfterPrefix)
											} else {
												writeOriginalName = true
											}
										} else {
											writeOriginalName = true
										}
									} else {
										writeOriginalName = true
									}
								}
								if writeOriginalName {
									newVariant.CommentBefore = appendComment(
										newVariant.CommentBefore,
										fmt.Sprintf("// tlgen:tl1name:\"%s\"", combinator.Construct.Name.String()),
									)
								}
							}

							if len(combinator.Fields) == 1 && combinator.Fields[0].FieldName == "" {
								newVariant.IsTypeAlias = true
								newVariant.TypeAlias = resolveType(combinator.Fields[0].FieldType, natIsConstant, natTemplates)
							} else {
								newVariant.IsTypeAlias = false
								newVariant.Fields = addFields(combinator.Fields, natIsConstant, natTemplates)
							}

							tl2Combinator.TypeDecl.Type.UnionType.Variants = append(tl2Combinator.TypeDecl.Type.UnionType.Variants, newVariant)
						}
					}
				} else {
					tl2Combinator.TypeDecl.Type.IsConstructorFields = true
				}

				if isRef {
					tl2Combinator.CommentBefore = appendComment(
						tl2Combinator.CommentBefore,
						fmt.Sprintf("// tlgen:tl1type:\"%s\"", combinator0.TypeDecl.String()),
					)
				} else {
					// add comment
					tl2Combinator.CommentBefore = appendComment(
						tl2Combinator.CommentBefore,
						combinator0.CommentBefore,
					)
					tl2Combinator.CommentBefore = appendComment(
						tl2Combinator.CommentBefore,
						combinator0.CommentRight,
					)
				}
				file.Combinators = append(file.Combinators, tl2Combinator)
			}
		}

		orderedFunctions := slices.Clone(localInfo.Functions)
		sort.Slice(orderedFunctions, func(i, j int) bool {
			return orderedFunctions[i].OriginalOrderIndex < orderedFunctions[j].OriginalOrderIndex
		})

		for _, function := range orderedFunctions {
			tl2Combinator := tlast.TL2Combinator{IsFunction: true}
			// add modifiers
			for _, modifier := range function.Modifiers {
				tl2Combinator.Annotations = append(tl2Combinator.Annotations,
					tlast.TL2Annotation{
						Name: modifier.Name,
					},
				)
			}
			if len(function.TemplateArguments) > 0 {
				tl2Combinator.Annotations = append(tl2Combinator.Annotations,
					tlast.TL2Annotation{
						Name: tl1Diagonal,
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
			if len(function.TemplateArguments) == 0 {
				// add arguments
				tl2Combinator.FuncDecl.Arguments = addFields(function.Fields, make(map[string]bool), make(map[string]bool))
				// add return type
				nats := make(map[string]bool)
				for _, field := range function.Fields {
					if field.FieldType.String() == "#" {
						nats[field.FieldName] = true
					}
				}
				tl2Combinator.FuncDecl.ReturnType = tlast.TL2TypeDefinition{
					TypeAlias: resolveType(function.FuncDecl, make(map[string]bool), nats),
				}
			}
			// add comment
			tl2Combinator.CommentBefore = appendComment(tl2Combinator.CommentBefore, function.CommentBefore)
			tl2Combinator.CommentBefore = appendComment(tl2Combinator.CommentBefore, function.CommentRight)

			file.Combinators = append(file.Combinators, tl2Combinator)
		}

		newState = append(newState, FileToWrite{Path: part.File, Ast: file})
	}

	if DEBUG {
		printDebugInfo(associatedWrappers, natUsage, info)
	}

	return gen.MergeMigrationState(prevState, newState)
}

func (gen *Gen2) MergeMigrationState(oldState, newState []FileToWrite) ([]FileToWrite, error) {
	if !gen.options.TL2MigrateByNamespaces || !gen.options.TL2ContinuousMigration {
		return newState, nil
	}

	resultingState := make([]FileToWrite, 0)

	presentingNewFiles := make(map[string]FileToWrite)
	presentingOldFiles := make(map[string]FileToWrite)

	for _, file := range oldState {
		presentingOldFiles[file.Path] = file
	}

	for _, file := range newState {
		presentingNewFiles[file.Path] = file
	}

	// in old but not in new
	for path, file := range presentingOldFiles {
		if _, ok := presentingNewFiles[path]; !ok {
			resultingState = append(resultingState, file)
		}
	}

	// in new but not in old
	for path, file := range presentingNewFiles {
		if _, ok := presentingOldFiles[path]; !ok {
			resultingState = append(resultingState, file)
		}
	}

	// in both
	for path, oldFile := range presentingOldFiles {
		if newFile, ok := presentingNewFiles[path]; ok {
			resultingFile, err := gen.mergeTL2Files(oldFile.Ast, newFile.Ast)
			if err != nil {
				return nil, err
			}
			resultingState = append(resultingState, FileToWrite{Path: path, Ast: resultingFile})
		}
	}

	sort.Slice(resultingState, func(i, j int) bool {
		return strings.Compare(resultingState[i].Path, resultingState[j].Path) > 0
	})

	return resultingState, nil
}

type TL2CombinatorIndexed struct {
	Name       tlast.TL2TypeName
	Combinator tlast.TL2Combinator

	Index int
}

func indexCombinators(file tlast.TL2File) (result []TL2CombinatorIndexed) {
	result = make([]TL2CombinatorIndexed, len(file.Combinators))
	for i, combinator := range file.Combinators {
		result[i].Index = i
		result[i].Combinator = combinator
		result[i].Name = combinator.TypeDecl.Name
		if combinator.IsFunction {
			result[i].Name = combinator.FuncDecl.Name
		}
	}
	return
}

func id[T any](x T) T { return x }

func (gen *Gen2) mergeTL2Files(oldFile, newFile tlast.TL2File) (tlast.TL2File, error) {
	resultingMapping := make(map[tlast.TL2TypeName]TL2CombinatorIndexed)

	oldCombs := indexCombinators(oldFile)
	newCombs := indexCombinators(newFile)

	getName := func(c TL2CombinatorIndexed) tlast.TL2TypeName { return c.Name }

	oldMapping := utils.SliceToMap(oldCombs, getName, id[TL2CombinatorIndexed])
	newMapping := utils.SliceToMap(newCombs, getName, id[TL2CombinatorIndexed])

	oldNames := utils.Keys(oldMapping)
	newNames := utils.Keys(newMapping)

	intersection := utils.SetIntersection(
		utils.SliceToSet(oldNames),
		utils.SliceToSet(newNames),
	)

	for name := range intersection {
		options := tlast.NewCanonicalFormatOptions()

		oldComb := oldMapping[name]
		newComb := newMapping[name]

		oldStr := strings.Builder{}
		newStr := strings.Builder{}

		oldComb.Combinator.Print(&oldStr, options)
		newComb.Combinator.Print(&newStr, options)

		if oldStr.String() == newStr.String() {
			resultingMapping[name] = newMapping[name]
		} else {
			oldIsRef := oldComb.Combinator.HasAnnotation("tl1")
			newIsRef := newComb.Combinator.HasAnnotation("tl1")

			if oldIsRef && !newIsRef {
				resultingMapping[name] = newMapping[name]
			} else {
				return tlast.TL2File{}, fmt.Errorf("can't continue migration due to incompetible versions to \"%s\"", name)
			}
		}
	}

	union := utils.SetUnion(
		utils.SliceToSet(oldNames),
		utils.SliceToSet(newNames),
	)

	for name := range union {
		if _, ok := intersection[name]; ok {
			continue
		}
		if comb, ok := oldMapping[name]; ok {
			resultingMapping[name] = comb
			continue
		}
		if comb, ok := newMapping[name]; ok {
			resultingMapping[name] = comb
			continue
		}
	}

	combs := utils.Values(resultingMapping)
	sort.Slice(combs, func(i, j int) bool {
		_, isOldI := oldMapping[combs[i].Name]
		_, isOldJ := oldMapping[combs[j].Name]

		if isOldI == isOldJ {
			return combs[i].Index <= combs[j].Index
		} else {
			return isOldI
		}
	})

	file := tlast.TL2File{}
	for _, comb := range combs {
		file.Combinators = append(file.Combinators, comb.Combinator)
	}

	return file, nil
}

func getTypesInfoFromTL2State(state []FileToWrite) map[tlast.TL2TypeName][]tlast.TL2Combinator {
	result := make(map[tlast.TL2TypeName][]tlast.TL2Combinator)
	for _, file := range state {
		for _, combinator := range file.Ast.Combinators {
			if combinator.IsFunction {
				continue
			}
			name := combinator.TypeDecl.Name
			if name.String() == "vector" {
				//debugf("debug")
			}
			if combinator.HasAnnotation(tl2Ext) {
				suffix := ""
				for _, argument := range combinator.TypeDecl.TemplateArguments {
					if argument.Category.IsUint32() {
						suffix += "_" + strings.ToUpper(argument.Name[:1]) + argument.Name[1:]
					}
				}
				name.Name, _ = strings.CutSuffix(name.Name, suffix)
			}
			result[name] = append(result[name], combinator)
		}
	}
	return result
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
		//file, _ := os.Open("usages.json")
		currefs := natUsage.CombinatorsNatFieldsToArraySizeReference

		tps := utils.Keys(currefs)
		sort.Slice(tps, func(i, j int) bool {
			return compareNames(tps[i], tps[j])
		})

		for _, tp := range tps {
			args := utils.Keys(currefs[tp])
			sort.Ints(args)

			fmt.Printf("<<<<<<\n%s:\n[\n", tp.String())
			for arg := range args {
				refs := currefs[tp][arg]

				refsKeys := utils.Keys(refs)
				sort.Slice(refsKeys, func(i, j int) bool {
					return compareNames(refsKeys[i], refsKeys[j])
				})

				fmt.Printf("\t%d: [\n", arg)
				for _, ref := range refsKeys {
					fields := refs[ref]

					fieldsKeys := utils.Keys(fields)
					sort.Ints(fieldsKeys)

					for _, field := range fieldsKeys {
						path := fields[field]
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

		fmt.Printf("List to remove deps (%d): [\n", len(tps))
		for _, tp := range tps {
			fmt.Printf("\t%s,\n", tp)
		}
		fmt.Println("]")
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

func appendComment(comment string, newLines string) string {
	if newLines == "" {
		return comment
	}
	if comment == "" {
		return newLines
	}
	return comment + "\n" + newLines
}
