// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
)

// TODO - name collision checks

type Kernel struct {
	opts                  *OptionsKernel
	rpcPreferTL2WhiteList Whitelist
	// each type can have up to 3 elements in this map, TL1 constructor, TL1 type and canonical primitive name
	tips         map[string]*KernelType
	tipsOrdered  []*KernelType
	tipsTopLevel []*KernelType

	brackets *KernelType

	instances        map[string]*TypeInstanceRef
	instancesOrdered []*TypeInstanceRef

	filesTL1 tlast.TL
	filesTL2 []tlast.TL2Combinator

	filesTL1Full tlast.TL // for TLO generation. Remove after TLO is removed,

	supportedAnnotations map[string]struct{}
	allAnnotations       []string // position is bit
}

// Add builtin types
func NewKernel(opts *OptionsKernel) *Kernel {
	k := &Kernel{
		opts:                  opts,
		rpcPreferTL2WhiteList: NewWhiteList(opts.RPCPreferTL2WhiteList),
		brackets: &KernelType{
			originTL2: true,
			builtin:   true,
			instances: map[string]*TypeInstanceRef{},
			canBeBare: true,
		},
		tips:      map[string]*KernelType{},
		instances: map[string]*TypeInstanceRef{},
	}
	k.addPrimitive("uint32", "#", "nat", &KernelValueUint32{}, true)
	k.addPrimitive("int32", "int", "int", &KernelValueInt32{}, true)
	k.addPrimitive("float32", "float", "float", &KernelValueInt32{}, false)
	k.addPrimitive("uint64", "", "uint64", &KernelValueUint64{}, true)
	k.addPrimitive("int64", "long", "long", &KernelValueInt64{}, true)
	k.addPrimitive("float64", "double", "double", &KernelValueInt64{}, false)
	k.addPrimitive("byte", "", "byte", &KernelValueByte{}, true)
	k.addPrimitive("bool", "", "bool", &KernelValueBool{}, true)
	k.addPrimitive("bit", "", "bit", &KernelValueBit{}, true)
	k.addPrimitive("string", "string", "string", &KernelValueString{}, true)
	//k.addString()
	k.addTL1Brackets()

	k.supportedAnnotations = map[string]struct{}{"read": {}, "any": {}, "internal": {}, "write": {}, "readwrite": {}, "kphp": {}}
	// TODO - add from options

	return k
}

func (k *Kernel) addTip(kt *KernelType, name1 string, name2 string) error {
	_, ok := k.tips[name1]
	if ok {
		return fmt.Errorf("type %v already exists", name1)
	}
	if name2 != "" {
		_, ok := k.tips[name2]
		if ok {
			return fmt.Errorf("type %v already exists", name2)
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
	kt.instances[canonicalName] = ref
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

func (k *Kernel) TL1FullForTLO() []*tlast.Combinator {
	return k.filesTL1Full
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
	fullTL, err := tlast.ParseTLFile(dataStr, file, tlast.LexerOptions{AllowDirty: true})
	if err != nil {
		return err // Do not add excess info to already long parse error
	}

	// we do not want to support those, they are soon to be removed forever
	dataStr = strings.ReplaceAll(dataStr, "_ {X:Type} result:X = ReqResult X;", "")
	dataStr = strings.ReplaceAll(dataStr, "engine.query {X:Type} query:!X = engine.Query;", "")
	dataStr = strings.ReplaceAll(dataStr, "engine.queryShortened query:%(VectorTotal int) = engine.Query;", "")

	tl, err := tlast.ParseTLFile(dataStr, file, tlast.LexerOptions{AllowDirty: false})
	if err != nil {
		return err // Do not add excess info to already long parse error
	}
	k.filesTL1 = append(k.filesTL1, tl...)
	k.filesTL1Full = append(k.filesTL1Full, fullTL...)
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

func (k *Kernel) Compile() error {
	namespaceTL1SeeHere := map[string]*tlast.ParseError{}
	if err := k.CompileTL1(namespaceTL1SeeHere); err != nil {
		return err
	}
	log.Printf("tl2pure: compiling %d TL2 combinators", len(k.filesTL2))
	// then add all TL2 declarations
	for _, comb := range k.filesTL2 {
		refName := comb.ReferenceName()
		log.Printf("tl2pure: compiling %s", comb)
		if e2, ok := namespaceTL1SeeHere[refName.Namespace]; ok && refName.Namespace != "" {
			e1 := comb.ReferenceNamePR().BeautifulError(fmt.Errorf("namespace %s must be defined entirely in either .tl or in .tl2 file(s)", refName.Namespace))
			return tlast.BeautifulError2(e1, e2)
		}
		kt := &KernelType{
			originTL2:      true,
			combTL2:        comb,
			instances:      map[string]*TypeInstanceRef{},
			isFunction:     comb.IsFunction,
			isTopLevel:     len(comb.TypeDecl.TemplateArguments) == 0,
			canBeBare:      true,
			canonicalName:  tlast.Name(refName),
			historicalName: tlast.Name(refName),
		}
		if !comb.IsFunction {
			kt.tl1Names = map[string]struct{}{}
			kt.tl2Names = map[string]struct{}{refName.String(): {}}
			var nc NameCollision
			for _, targ := range comb.TypeDecl.TemplateArguments {
				if err := nc.AddUniqueName(targ.Name, targ.PR, "template argument"); err != nil {
					return err
				}
			}
		}
		if err := k.addTip(kt, refName.String(), ""); err != nil {
			return fmt.Errorf("error adding type %s: %w", refName, err)
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
					if _, ok := k.supportedAnnotations[m.Name]; !ok && utils.DoLint(typ.CommentRight) {
						e1 := m.PR.BeautifulError(fmt.Errorf("annotation %q not known to tlgen", m.Name))
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
					// to ignore diagonal legacy
					//if m.Name == tl1Diagonal {
					//	continue
					//}
					if _, ok := allAnnotations[m.Name]; !ok {
						if _, ok := k.supportedAnnotations[m.Name]; !ok {
							e1 := m.PR.BeautifulError(fmt.Errorf("annotation %q not known to tlgen", m.Name))
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
	//instantiate all top-level declarations
	for _, tip := range k.tipsOrdered {
		if tip.originTL2 {
			if tip.combTL2.IsFunction {
				tr := tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: tip.combTL2.FuncDecl.Name}}
				if _, err := k.getInstanceTL2(tr, true); err != nil {
					return fmt.Errorf("error adding function %s: %w", tr.String(), err)
				}
				continue
			}
			typeDecl := tip.combTL2.TypeDecl
			if len(typeDecl.TemplateArguments) != 0 {
				continue // instantiate templates on demand only
			}
			tr := tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: typeDecl.Name}}
			if _, err := k.getInstanceTL2(tr, true); err != nil {
				return fmt.Errorf("error adding type %s: %w", tr.String(), err)
			}
		}
	}
	for _, tip := range k.tipsTopLevel {
		if !tip.originTL2 {
			tr := tlast.TypeRef{Type: tip.canonicalName}
			if _, _, err := k.getInstanceTL1(tr, true); err != nil {
				return err
			}
		}
	}
	var cf cycleFinder
	for _, ref := range k.instancesOrdered {
		cf.reset()
		ref.ins.FindCycle(&cf)
		res := cf.printCycle()
		if res == "TODO - implement later" {
			return fmt.Errorf("found recursive cycle %s", res)
		}
	}
	tl2WhiteList := NewWhiteList(k.opts.TL2WhiteList)
	tl2Children := map[TypeInstance]struct{}{}
	typesCounterMarkTL2 := 0
	for _, v := range k.instancesOrdered {
		if tl2WhiteList.HasName(v.ins.Common().tlName) {
			k.markWantsTL2(v.ins, tl2Children)
			typesCounterMarkTL2++
		}
	}
	if k.opts.Verbose && !tl2WhiteList.Empty() {
		log.Printf("found %d object roots for TL2 versions of types by the following filter: %s", typesCounterMarkTL2, k.opts.TL2WhiteList)
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
