// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/VKCOM/tl/internal/tlast"
)

// TODO - name collision checks

type Kernel struct {
	opts                  *OptionsKernel
	rpcPreferTL2WhiteList Whitelist
	tl2WhiteList          Whitelist
	// each type can have up to 3 elements in this map, TL1 constructor, TL1 type and canonical primitive name
	tips         map[string]*KernelType
	tipsOrdered  []*KernelType
	tipsTopLevel []*KernelType

	brackets *KernelType

	instances        map[string]*TypeInstanceRef
	instancesOrdered []*TypeInstanceRef

	filesTL1 tlast.TL
	filesTL2 []tlast.TL2Combinator

	supportedAnnotations map[string]struct{}
	allAnnotations       []string // position is bit
}

// Add builtin types
func NewKernel(opts *OptionsKernel) *Kernel {
	k := &Kernel{
		opts:                  opts,
		rpcPreferTL2WhiteList: NewWhiteList("--rpcPreferTL2WhiteList", opts.RPCPreferTL2WhiteList),
		tl2WhiteList:          NewWhiteList("--tl2WhiteList", opts.TL2WhiteList),
		brackets: &KernelType{
			originTL2: true,
			builtin:   true,
			instances: map[string]*TypeInstanceRef{},
			canBeBare: true,
		},
		tips:      map[string]*KernelType{},
		instances: map[string]*TypeInstanceRef{},
	}
	_ = k.addPrimitive("uint32", "#", "nat", 4, true)
	_ = k.addPrimitive("int32", "int", "int", 4, true)
	_ = k.addPrimitive("float32", "float", "float", 4, false)
	_ = k.addPrimitive("uint64", "", "uint64", 8, true)
	_ = k.addPrimitive("int64", "long", "long", 8, true)
	_ = k.addPrimitive("float64", "double", "double", 8, false)
	_ = k.addPrimitive("byte", "", "byte", 1, true)
	{
		ktBool := k.addPrimitive("bool", "", "bool", 1, true)
		//ktBool.originTL2 = false
		ktBool.canBeBare = false
		ktBool.tl1BoxedName = tlast.TL2TypeName{Name: "bool"}
		// Bool is special, we treat as unions, they are boxed in TL1, therefore we are making them boxed in TL2
		// so that type references are generated exactly same from TL1 and TL2 files.
	}
	_ = k.addPrimitive("bit", "", "bit", 0, true)
	_ = k.addPrimitive("string", "string", "string", 0, true)
	k.addTL1Brackets()
	_ = k.addPrimitive("__function", "__function", "function", 0, false)
	{
		ktFetcher := k.addPrimitive("__function_result", "__function_result", "function_result", 0, false)
		ktFetcher.canBeBare = false
		ktFetcher.tl1BoxedName = tlast.TL2TypeName{Name: "__function_result"}
	}
	k.supportedAnnotations = map[string]struct{}{"read": {}, "any": {}, "internal": {}, "write": {}, "readwrite": {}, "kphp": {}}
	for _, ann := range strings.Split(k.opts.Annotations, ",") {
		// TODO - validate here
		ann = strings.TrimSpace(ann)
		if ann != "" {
			k.supportedAnnotations[ann] = struct{}{}
		}
	}
	return k
}

func (k *Kernel) addTip(kt *KernelType, name1 string, name2 string) error {
	existing, ok := k.tips[name1]
	if ok {
		if kt.namePR == (tlast.PositionRange{}) { // various built-ins
			return fmt.Errorf("type %v already exists", name1)
		}
		e1 := kt.namePR.BeautifulError(fmt.Errorf("type %v already exists", name1))
		if existing.namePR == (tlast.PositionRange{}) {
			return e1
		}
		e2 := existing.namePR.BeautifulError(errSeeHere)
		return tlast.BeautifulError2(e1, e2)
	}
	if name2 != "" {
		existing, ok := k.tips[name2]
		if ok {
			if kt.namePR == (tlast.PositionRange{}) { // various built-ins
				return fmt.Errorf("type %v already exists", name2)
			}
			e1 := kt.namePR.BeautifulError(fmt.Errorf("type %v already exists", name2))
			if existing.namePR == (tlast.PositionRange{}) {
				return e1
			}
			e2 := existing.namePR.BeautifulError(errSeeHere)
			return tlast.BeautifulError2(e1, e2)
		}
		k.tips[name2] = kt
	}
	k.tips[name1] = kt
	k.tipsOrdered = append(k.tipsOrdered, kt)
	if kt.originTL2 {
		if kt.combTL2.IsFunction || len(kt.combTL2.TypeDecl.TemplateArguments) == 0 {
			k.tipsTopLevel = append(k.tipsTopLevel, kt)
		}
	} else {
		if len(kt.combTL1[0].TemplateArguments) == 0 {
			k.tipsTopLevel = append(k.tipsTopLevel, kt)
		}
	}
	return nil
}

func (k *Kernel) addInstance(canonicalName string, kt *KernelType) *TypeInstanceRef {
	ref := &TypeInstanceRef{}
	if _, ok := kt.instances[canonicalName]; ok {
		panic(fmt.Sprintf("type instance list contains duplicate %q", canonicalName))
	}
	if _, ok := k.instances[canonicalName]; ok {
		panic(fmt.Sprintf("global instance list contains duplicate %q", canonicalName))
	}
	kt.instances[canonicalName] = ref // storing pointer terminates recursion
	kt.instancesOrdered = append(kt.instancesOrdered, ref)

	k.instances[canonicalName] = ref // storing pointer terminates recursion
	k.instancesOrdered = append(k.instancesOrdered, ref)
	return ref
}

func (k *Kernel) AllAnnotations() []string {
	return k.allAnnotations
}

func (k *Kernel) TopLevelTypeInstances() []TypeInstance {
	var result []TypeInstance
	for _, tip := range k.tipsTopLevel {
		for _, ref := range tip.instances {
			result = append(result, ref.ins)
		}
	}
	return result
}

func (k *Kernel) AllTypeInstances() []TypeInstance {
	var result []TypeInstance
	for _, ref := range k.instancesOrdered {
		result = append(result, ref.ins)
	}
	return result
}

func (k *Kernel) AllTypePrimitives() []*TypeInstancePrimitive {
	return AllSelectedTypeInstances[*TypeInstancePrimitive](k)
}

//func (k *Kernel) FunctionInstances() []*TypeInstanceStruct {
//	return k.functionsOrdered
//}

// TODO - remove or fix
func (k *Kernel) GetFunctionInstance(name tlast.TL2TypeName) *TypeInstanceStruct {
	tip, ok := k.tips[name.String()]
	if !ok {
		return nil
	}
	ref, ok := tip.instances[name.String()]
	if !ok {
		return nil
	}
	ins2, _ := ref.ins.(*TypeInstanceStruct)
	return ins2
}

func (k *Kernel) TL1() []*tlast.Combinator {
	return k.filesTL1
}

func (k *Kernel) TL2() []tlast.TL2Combinator {
	return k.filesTL2
}

func (k *Kernel) AddFileTL1(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("error reading schema file %q: %w", file, err)
	}
	dataStr := string(data)
	// TODO - set AllowDirty to false after removing _ {X:Type} result:X = ReqResult X;
	tl, err := tlast.ParseTLFile(dataStr, file, tlast.LexerOptions{AllowDirty: true})
	if err != nil {
		return err // Do not add excess info to already long parse error
	}
	k.filesTL1 = append(k.filesTL1, tl...)
	return nil
}

func (k *Kernel) AddFileTL2(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("error reading tl2 schema file %q: %w", file, err)
	}
	dataStr := string(data)
	tl, err := tlast.ParseTL2File(dataStr, file, tlast.LexerOptions{LexerLanguage: tlast.TL2})
	if err != nil {
		return err // Do not add excess info to already long parse error
	}
	k.filesTL2 = append(k.filesTL2, tl.Combinators...)
	return nil
}

func (k *Kernel) functionNeedsCustomResultType(funcDecl tlast.TL2FuncDeclaration) (_ tlast.TL2TypeName, needsCustom bool, customTrue bool) {
	if funcDecl.ReturnType.IsTypeAlias {
		return tlast.TL2TypeName{}, false, false
	}
	structType := funcDecl.ReturnType.StructType
	if !structType.IsUnionType && len(structType.ConstructorFields) == 0 {
		if ktTrue, ok := k.tips["true"]; ok && k.IsTrueType(ktTrue) {
			return tlast.TL2TypeName{Name: "true"}, true, true
		}
	}
	if funcDecl.ReturnType.StructType.IsUnionType ||
		len(funcDecl.ReturnType.StructType.ConstructorFields) != 1 ||
		funcDecl.ReturnType.StructType.ConstructorFields[0].Name != "" ||
		funcDecl.ReturnType.StructType.ConstructorFields[0].IsOptional {
		resultTlName := funcDecl.Name
		resultTlName.Name += "__Result"
		return resultTlName, true, false
	}
	return tlast.TL2TypeName{}, false, false
}

func (k *Kernel) Compile() error {
	namespaceTL1SeeHere := map[string]*tlast.ParseError{}
	if err := k.CompileTL1(namespaceTL1SeeHere); err != nil {
		return err
	}
	fmt.Printf("tl2pure: compiling %d TL2 combinators\n", len(k.filesTL2))
	// then add all TL2 declarations
	for _, comb := range k.filesTL2 {
		refName := comb.ReferenceName()
		// fmt.Printf("tl2pure: compiling %s\n", comb)
		if e2, ok := namespaceTL1SeeHere[refName.Namespace]; ok && refName.Namespace != "" {
			e1 := comb.ReferenceNamePR().BeautifulError(fmt.Errorf("namespace %s must be defined entirely in either .tl or in .tl2 file(s)", refName.Namespace))
			return tlast.BeautifulError2(e1, e2)
		}
		kt := &KernelType{
			originTL2:     true,
			combTL2:       comb,
			instances:     map[string]*TypeInstanceRef{},
			isFunction:    comb.IsFunction,
			isTopLevel:    len(comb.TypeDecl.TemplateArguments) == 0,
			canBeBare:     true,
			canonicalName: refName,
		}
		if comb.IsFunction {
			kt.namePR = comb.FuncDecl.PRName
			if resultTlName, needsCustom, customTrue := k.functionNeedsCustomResultType(comb.FuncDecl); needsCustom && !customTrue {
				resultKt := &KernelType{
					originTL2: true,
					combTL2: tlast.TL2Combinator{
						TypeDecl: tlast.TL2TypeDeclaration{
							Name:   resultTlName,
							PRName: comb.FuncDecl.PRName,
							Type:   comb.FuncDecl.ReturnType,
							PR:     comb.FuncDecl.ReturnType.PR,
						},
						IsFunction:    false,
						PR:            tlast.PositionRange{},
						CommentBefore: "",
					},
					instances:     map[string]*TypeInstanceRef{},
					isFunction:    false,
					isTopLevel:    true,
					canBeBare:     true,
					canonicalName: resultTlName,
				}
				// TODO - we must not receive errors here, function
				if err := k.addTip(resultKt, resultTlName.String(), ""); err != nil {
					return fmt.Errorf("internal error when creating function result type %s: %w", resultTlName, err)
				}
			}
			if err := k.addTip(kt, refName.String(), ""); err != nil {
				return err
			}
		} else {
			kt.namePR = comb.TypeDecl.PRName
			kt.templateArguments = comb.TypeDecl.TemplateArguments
			kt.targs = make([]KernelTypeTarg, len(comb.TypeDecl.TemplateArguments))

			kt.tl1Names = map[string]struct{}{}
			kt.tl2Names = map[string]struct{}{refName.String(): {}}
			var nc NameCollision
			for _, targ := range comb.TypeDecl.TemplateArguments {
				if err := nc.AddUniqueName(targ.Name, targ.PR, "template argument"); err != nil {
					return err
				}
			}
			if err := k.addTip(kt, refName.String(), ""); err != nil {
				return err
			}
		}
	}
	// type all declarations by comparing type ref with actual type definition
	for _, tip := range k.tipsOrdered {
		if tip.originTL2 {
			if tip.combTL2.IsFunction {
				if err := k.typeCheck(tip.combTL2.FuncDecl.ReturnType, nil); err != nil {
					return err
				}
				if err := k.typeCheckAliasFields(false, tlast.TL2TypeRef{}, tip.combTL2.FuncDecl.Arguments, nil); err != nil {
					return err
				}
				continue
			}
			if err := k.typeCheck(tip.combTL2.TypeDecl.Type, tip.combTL2.TypeDecl.TemplateArguments); err != nil {
				return err
			}
		} else {
			//if tip.combTL1[0].IsFunction {
			//	comb := tip.combTL1[0]
			//	if err := k.typeCheckAliasFieldsTL1(comb.Fields, nil); err != nil {
			//		return err
			//	}
			//	var leftArgs []tlast.TemplateArgument
			//	for _, f := range comb.Fields {
			//		if f.FieldName != "" && f.FieldType.String() == "#" {
			//			// TODO - add other fields with wrong category to catch references to them
			//			leftArgs = append(leftArgs, tlast.TemplateArgument{
			//				FieldName: f.FieldName,
			//				IsNat:     true,
			//				PR:        f.PR,
			//			})
			//		}
			//	}
			//	if err := k.typeCheckTypeRefTL1(comb.FuncDecl, leftArgs); err != nil {
			//		return err
			//	}
			//	//if err := k.typeCheck(comb..FuncDecl.ReturnType, nil); err != nil {
			//	//	return err
			//	//}
			//	//if err := k.typeCheckAliasFields(false, tlast.TL2TypeRef{}, tip.combTL2.FuncDecl.Arguments, nil); err != nil {
			//	//	return err
			//	//}
			//	continue
			//}
			//for _, comb := range tip.combTL1 {
			//	if !k.shouldSkipDefinition(comb) {
			//		if err := k.typeCheckAliasFieldsTL1(comb.Fields, comb.TemplateArguments); err != nil {
			//			return err
			//		}
			//	}
			//}
		}
	}
	if err := k.checkTagCollisions(); err != nil {
		return err
	}
	if err := k.checkNamespaceCollisions(); err != nil {
		return err
	}

	{
		allAnnotations := map[string]struct{}{}
		// generated RPC code can depend on those annotations, even
		// if none present in current tl file.
		// so we add all supported annotations always.
		for m := range k.supportedAnnotations {
			allAnnotations[m] = struct{}{}
			k.allAnnotations = append(k.allAnnotations, m)
		}
		for _, typ := range k.filesTL1 {
			for _, m := range typ.Modifiers {
				if strings.ToLower(m.Name) != m.Name { // TODO - move into lexer
					return m.PR.BeautifulError(fmt.Errorf("annotations must be lower case"))
				}
				if _, ok := allAnnotations[m.Name]; !ok {
					if _, ok := k.supportedAnnotations[m.Name]; !ok {
						e1 := m.PR.BeautifulError(fmt.Errorf("annotation %q not known to tlgen, please add to --annotations command line argument", m.Name))
						if k.opts.WarningsAreErrors {
							return e1
						}
						e1.PrintWarning(k.opts.ErrorWriter, nil)
					}
					allAnnotations[m.Name] = struct{}{}
					k.allAnnotations = append(k.allAnnotations, m.Name)
					if len(k.allAnnotations) > 32 {
						return m.PR.BeautifulError(errors.New("too many different annotations, max is 32 for now"))
					}
				}
			}
		}
		for _, combinator := range k.filesTL2 {
			if combinator.IsFunction {
				for _, m := range combinator.Annotations {
					if _, ok := allAnnotations[m.Name]; !ok {
						if _, ok := k.supportedAnnotations[m.Name]; !ok {
							e1 := m.PR.BeautifulError(fmt.Errorf("annotation %q not known to tlgen, please add to --annotations command line argument", m.Name))
							if k.opts.WarningsAreErrors {
								return e1
							}
							e1.PrintWarning(k.opts.ErrorWriter, nil)
						}
						allAnnotations[m.Name] = struct{}{}
						k.allAnnotations = append(k.allAnnotations, m.Name)
						if len(k.allAnnotations) > 32 {
							return m.PR.BeautifulError(errors.New("too many different annotations, max is 32 for now"))
						}
					}
				}
			}
		}
		sort.Strings(k.allAnnotations)
	}
	refErrList, err := k.findTL1toTL2References(nil)
	if err != nil {
		return err
	}
	for _, err := range refErrList.migrationErrorList {
		return err // simply return the first error
	}
	for _, tip := range k.tipsOrdered {
		if !tip.isTopLevel {
			continue
		}
		if tip.exclamationArg != nil {
			if !k.opts.InstantiateExclamationWrappers {
				continue
			}
			_, _ = fmt.Fprintf(k.opts.ErrorWriter, "will instantiate exclamation wrapper %s\n", tip.canonicalName)
		}
		tr := tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: tip.canonicalName}}
		if _, _, err := k.getInstance(tr, true); err != nil {
			return err
		}
	}
	for _, ref := range k.instancesOrdered {
		if ins, ok := ref.ins.(*TypeInstanceStruct); ok {
			for i, field := range ins.fields {
				natFieldUsage := ins.GetNatFieldUsage(i, true, true)
				if natFieldUsage.UsedAsSize && natFieldUsage.UsedAsMask {
					e3 := natFieldUsage.usedAsMaskPR.BeautifulError(fmt.Errorf("used as mask here"))
					e3.PrintWarning(k.opts.ErrorWriter, nil)
					e1 := field.pr.BeautifulError(fmt.Errorf("#-field %s is used as tuple size, while already being used as a field mask", field.Name()))
					e2 := natFieldUsage.usedAsSizePR.BeautifulError(fmt.Errorf("used as size here"))
					return tlast.BeautifulError2(e1, e2)
				}
			}
		}
	}
	var cf cycleFinder
	for _, ref := range k.instancesOrdered {
		cf.reset()
		ref.ins.FindCycle(&cf, tlast.PositionRange{}) // never printed
		_ = cf.printCycle()
		//if err != nil {
		//	return err
		//}
		// TODO - this fix is not easy. If all constructors reach original type, this is error,
		// but if only some, this is OK. Complicated property.
		//myPlus a:MyNat2 = MyNat2; // struct with field
		//^^^^^^-- found infinite cycle MyNat2->myPlus->MyNat2, use optional to break it ./internal/tlcodegen/tes
	}
	tl2Children := map[TypeInstance]struct{}{}
	typesCounterMarkTL2 := 0
	for _, v := range k.instancesOrdered {
		if !v.ins.Common().OriginTL2() && k.tl2WhiteList.HasName2(v.ins.Common().tlName) {
			k.markWantsTL2(v.ins, tl2Children)
			typesCounterMarkTL2++
		}
	}
	if k.opts.Verbose && !k.tl2WhiteList.Empty() {
		fmt.Printf("found %d object roots for TL2 versions of types by the following filter: %s\n", typesCounterMarkTL2, k.opts.TL2WhiteList)
	}
	return nil
}

func (k *Kernel) checkTagCollisions() error {
	constructorTags := map[uint32]*tlast.ParseError{}
	for _, typ := range k.filesTL1 {
		crc32 := typ.Crc32()
		if crc32 == 0 {
			// typeA#00000000 = TypeA;
			return typ.Construct.IDPR.BeautifulError(fmt.Errorf("constructor tag 0 is prohibited, even if generated implicitly"))
		}
		if err, ok := constructorTags[crc32]; ok {
			// typeA#dfc15abf = TypeA;
			// typeB#dfc15abf = TypeB;
			e1 := typ.Construct.IDPR.BeautifulError(fmt.Errorf("constructor tag #%08x is used again by %q", crc32, typ.Construct.Name.String()))
			return tlast.BeautifulError2(e1, err)
		}
		constructorTags[crc32] = typ.Construct.IDPR.BeautifulError(errSeeHere)
	}
	for _, typ := range k.filesTL2 {
		if typ.IsFunction {
			crc32 := typ.FuncDecl.Magic
			if crc32 != 0 {
				if err, ok := constructorTags[crc32]; ok {
					e1 := typ.FuncDecl.PRID.BeautifulError(fmt.Errorf("constructor tag #%08x is used again by %q", crc32, typ.FuncDecl.Name.String()))
					return tlast.BeautifulError2(e1, err)
				}
				constructorTags[crc32] = typ.FuncDecl.PRID.BeautifulError(errSeeHere)
			}
			continue
		}
		crc32 := typ.TypeDecl.Magic
		if crc32 != 0 {
			if err, ok := constructorTags[crc32]; ok {
				e1 := typ.TypeDecl.PRID.BeautifulError(fmt.Errorf("constructor tag #%08x is used again by %q", crc32, typ.TypeDecl.Name.String()))
				return tlast.BeautifulError2(e1, err)
			}
			constructorTags[crc32] = typ.TypeDecl.PRID.BeautifulError(errSeeHere)
		}
	}
	return nil
}

func (k *Kernel) checkNamespaceCollisions() error {
	var nc NameCollision
	for _, comb := range k.filesTL1 {
		if err := nc.AddSameCaseName(comb.Construct.Name.Namespace, comb.Construct.NamePR, "namespace"); err != nil {
			return err
		}
		if !comb.IsFunction {
			if err := nc.AddSameCaseName(comb.TypeDecl.Name.Namespace, comb.TypeDecl.NamePR, "namespace"); err != nil {
				return err
			}
		}
	}
	for _, comb := range k.filesTL2 {
		if comb.IsFunction {
			if err := nc.AddSameCaseName(comb.FuncDecl.Name.Namespace, comb.FuncDecl.PRName, "namespace"); err != nil {
				return err
			}
		} else {
			if err := nc.AddSameCaseName(comb.TypeDecl.Name.Namespace, comb.TypeDecl.PRName, "namespace"); err != nil {
				return err
			}
		}
	}
	return nil
}

func (k *Kernel) markWantsTL2(node TypeInstance, visitedNodes map[TypeInstance]struct{}) {
	if _, ok := visitedNodes[node]; ok {
		return
	}
	node.Common().hasTL2 = true
	visitedNodes[node] = struct{}{}
	children := make([]TypeInstance, 0, 4) // avoids the majority of heap allocations
	children = node.GetChildren(children, true)
	for _, child := range children {
		k.markWantsTL2(child, visitedNodes)
	}
}

func (k *Kernel) PrintUnusedWarnings() {
	if err := k.rpcPreferTL2WhiteList.UnusedWarning(); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := k.tl2WhiteList.UnusedWarning(); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func AllSelectedTypeInstances[T any](k *Kernel) []T {
	result := make([]T, 0)
	for _, tip := range k.AllTypeInstances() {
		if exactTip, ok := tip.(T); ok {
			result = append(result, exactTip)
		}
	}
	return result
}
