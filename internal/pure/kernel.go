// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
)

// TODO - name collision checks

type KernelType struct {
	originTL2 bool
	combTL1   []*tlast.Combinator
	combTL2   tlast.TL2Combinator
	// index by canonical name
	instances map[string]*TypeInstanceRef
	// order of instantiation
	instancesOrdered []*TypeInstanceRef
}

type Kernel struct {
	tips         map[string]*KernelType // TL1 single constructor names are also here
	tipsOrdered  []*KernelType
	tipsTopLevel []*KernelType

	brackets *KernelType

	instances        map[string]*TypeInstanceRef
	instancesOrdered []*TypeInstanceRef

	filesTL1 tlast.TL
	filesTL2 []tlast.TL2Combinator
}

// Add builtin types
func NewKernel() *Kernel {
	k := &Kernel{
		brackets:  &KernelType{originTL2: true, instances: map[string]*TypeInstanceRef{}},
		tips:      map[string]*KernelType{},
		instances: map[string]*TypeInstanceRef{},
	}
	k.addPrimitive("uint32", false, &KernelValueUint32{}, true)
	k.addPrimitive("int32", false, &KernelValueInt32{}, true)
	k.addPrimitive("uint64", true, &KernelValueUint64{}, true)
	k.addPrimitive("int64", false, &KernelValueInt64{}, true)
	k.addPrimitive("byte", true, &KernelValueByte{}, true)
	k.addPrimitive("bool", true, &KernelValueBool{}, true)
	k.addPrimitive("bit", true, &KernelValueBit{}, false)
	k.addString()

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
		if kt.combTL1[0].IsFunction || len(kt.combTL1[0].TemplateArguments) == 0 {
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

func (k *Kernel) TopLeveTypeInstances() []TypeInstance {
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

func (k *Kernel) AddParsedFileTL1(f tlast.TL) {
	k.filesTL1 = append(k.filesTL1, f...)
}

func (k *Kernel) AddParsedFileTL2(f tlast.TL2File) {
	k.filesTL2 = append(k.filesTL2, f.Combinators...)
}

func (k *Kernel) AddFileTL1(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("error reading schema file %q: %w", file, err)
	}
	dataStr := string(data)
	tl, err := tlast.ParseTLFile(dataStr, file, tlast.LexerOptions{AllowDirty: true})
	if err != nil {
		return err // Do not add excess info to already long parse error
	}
	k.AddParsedFileTL1(tl)
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
	k.AddParsedFileTL2(tl)
	return nil
}

func (k *Kernel) normalizeName(s string) string {
	s = strings.ReplaceAll(s, "_", "")
	return strings.ToLower(s)
}

func (k *Kernel) Compile(opts *OptionsKernel) error {
	if err := k.CompileTL1(); err != nil {
		return err
	}
	log.Printf("tl2pure: compiling %d TL2 combinators", len(k.filesTL2))
	// then add all TL2 declarations
	for _, comb := range k.filesTL2 {
		log.Printf("tl2pure: compiling %s", comb)
		kt := &KernelType{
			originTL2: true,
			combTL2:   comb,
			instances: map[string]*TypeInstanceRef{},
		}
		if !comb.IsFunction {
			namesNormalized := map[string]int{}
			names := map[string]int{}
			for i, targ := range comb.TypeDecl.TemplateArguments {
				if _, ok := names[targ.Name]; ok {
					return fmt.Errorf("error adding type %s: template argument %s name collision", comb.TypeDecl.Name, targ.Name)
				}
				nn := k.normalizeName(targ.Name)
				if _, ok := names[nn]; ok {
					return fmt.Errorf("error adding type %s: template argument %s normalized name collision", comb.TypeDecl.Name, nn)
				}
				namesNormalized[nn] = i
			}
		}
		if err := k.addTip(kt, comb.ReferenceName().String(), ""); err != nil {
			return fmt.Errorf("error adding type %s: %w", comb.ReferenceName(), err)
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
			if tip.combTL1[0].IsFunction {
				comb := tip.combTL1[0]
				if err := k.typeCheckAliasFieldsTL1(comb.Fields, nil); err != nil {
					return err
				}
				var leftArgs []tlast.TemplateArgument
				for _, f := range comb.Fields {
					if f.FieldName != "" && f.FieldType.String() == "#" {
						// TODO - add other fields with wrong category to catch references to them
						leftArgs = append(leftArgs, tlast.TemplateArgument{
							FieldName: f.FieldName,
							IsNat:     true,
							PR:        f.PR,
						})
					}
				}
				if err := k.typeCheckTypeRefTL1(comb.FuncDecl, leftArgs); err != nil {
					return err
				}
				//if err := k.typeCheck(comb..FuncDecl.ReturnType, nil); err != nil {
				//	return err
				//}
				//if err := k.typeCheckAliasFields(false, tlast.TL2TypeRef{}, tip.combTL2.FuncDecl.Arguments, nil); err != nil {
				//	return err
				//}
				continue
			}
			for _, comb := range tip.combTL1 {
				if err := k.typeCheckAliasFieldsTL1(comb.Fields, comb.TemplateArguments); err != nil {
					return err
				}
			}
		}
	}
	if err := k.CheckTagCollisions(opts); err != nil {
		return err
	}
	if err := k.CheckNamespaceCollisions(opts); err != nil {
		return err
	}
	//instantiate all top-level declarations
	for _, tip := range k.tipsOrdered {
		if tip.originTL2 {
			if tip.combTL2.IsFunction {
				tr := tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: tip.combTL2.FuncDecl.Name}}
				if _, err := k.getInstance(tr); err != nil {
					return fmt.Errorf("error adding function %s: %w", tr.String(), err)
				}
				continue
			}
			typeDecl := tip.combTL2.TypeDecl
			if len(typeDecl.TemplateArguments) != 0 {
				continue // instantiate templates on demand only
			}
			tr := tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: typeDecl.Name}}
			if _, err := k.getInstance(tr); err != nil {
				return fmt.Errorf("error adding type %s: %w", tr.String(), err)
			}
		} else {
			comb := tip.combTL1[0]
			if comb.IsFunction {
				tr := tlast.TypeRef{Type: comb.Construct.Name}
				if _, err := k.getInstanceTL1(tr); err != nil {
					return fmt.Errorf("error adding function %s: %w", tr.String(), err)
				}
				continue
			}
			if len(comb.TemplateArguments) != 0 {
				continue // instantiate templates on demand only
			}
			tr := tlast.TypeRef{Type: comb.Construct.Name}
			if len(tip.combTL1) != 1 {
				tr = tlast.TypeRef{Type: comb.TypeDecl.Name}
			}
			if _, err := k.getInstanceTL1(tr); err != nil {
				return fmt.Errorf("error adding type %s: %w", tr.String(), err)
			}
		}
	}
	// It is not easy to check all cycles before instantiation, so we do it afterward.
	var cf cycleFinder
	for _, ref := range k.instancesOrdered {
		cf.reset()
		ref.ins.FindCycle(&cf)
		res := cf.printCycle()
		if res != "" {
			return fmt.Errorf("found recursive cycle %s", res)
		}
	}
	return nil
}

func (k *Kernel) CheckTagCollisions(opts *OptionsKernel) error {
	// TODO
	return nil
}

func (k *Kernel) CheckNamespaceCollisions(opts *OptionsKernel) error {
	// TODO
	return nil
}
