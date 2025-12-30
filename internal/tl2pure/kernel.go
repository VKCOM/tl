package tl2pure

import (
	"fmt"
	"log"
	"math/rand"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
)

// common for read/write/json/etc... for simplicity
type TL2Context struct {
}

type KernelValue interface {
	Clone() KernelValue

	Reset()
	Random(rg *rand.Rand)
	WriteTL2(w *ByteBuilder, optimizeEmpty bool, onPath bool, level int, model *UIModel)
	ReadTL2(r []byte, ctx *TL2Context) ([]byte, error)
	WriteJSON(w []byte, ctx *TL2Context) []byte

	UIWrite(sb *strings.Builder, onPath bool, level int, model *UIModel)
	UIFixPath(side int, level int, model *UIModel) int // always called onPath
	UIStartEdit(level int, model *UIModel, createMode int)
	UIKey(level int, model *UIModel, insert bool, delete bool, up bool, down bool)

	CompareForMapKey(other KernelValue) int
}

type KernelType struct {
	comb tlast.TL2Combinator
	// index by canonical name
	instances        map[string]*TypeInstanceRef
	instancesOrdered []*TypeInstanceRef
}

type Kernel struct {
	tips         map[tlast.TL2TypeName]*KernelType
	tipsOrdered  []*KernelType
	tipsTopLevel []*KernelType

	brackets *KernelType

	instances        map[string]*TypeInstanceRef
	instancesOrdered []*TypeInstanceRef

	files []tlast.TL2Combinator
}

// Add builtin types
func NewKernel() *Kernel {
	k := &Kernel{
		brackets:  &KernelType{instances: map[string]*TypeInstanceRef{}},
		tips:      map[tlast.TL2TypeName]*KernelType{},
		instances: map[string]*TypeInstanceRef{},
	}
	k.addPrimitive("uint32", &KernelValueUint32{}, true)
	k.addPrimitive("int32", &KernelValueInt32{}, true)
	k.addPrimitive("uint64", &KernelValueUint64{}, true)
	k.addPrimitive("int64", &KernelValueInt64{}, true)
	k.addPrimitive("byte", &KernelValueByte{}, true)
	k.addPrimitive("bool", &KernelValueBool{}, true)
	k.addPrimitive("bit", &KernelValueBit{}, false)
	k.addString()

	return k
}

func (k *Kernel) addTip(kt *KernelType) error {
	_, ok := k.tips[kt.comb.ReferenceName()]
	if ok {
		return fmt.Errorf("type %v already exists", kt.comb.ReferenceName())
	}
	k.tips[kt.comb.ReferenceName()] = kt
	k.tipsOrdered = append(k.tipsOrdered, kt)
	if kt.comb.IsFunction || len(kt.comb.TypeDecl.TemplateArguments) == 0 {
		k.tipsTopLevel = append(k.tipsTopLevel, kt)
	}
	return nil
}

func (k *Kernel) IsBit(tr tlast.TL2TypeRef) bool {
	return !tr.IsBracket && tr.SomeType.Name.Namespace == "" && tr.SomeType.Name.Name == "bit"
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

func (k *Kernel) GetFunctionInstance(name tlast.TL2TypeName) *TypeInstanceObject {
	tip, ok := k.tips[name]
	if !ok {
		return nil
	}
	ref, ok := tip.instances[name.String()]
	if !ok {
		return nil
	}
	ins2, _ := ref.ins.(*TypeInstanceObject)
	return ins2
}

func (k *Kernel) AddFile(f tlast.TL2File) {
	k.files = append(k.files, f.Combinators...)
}

func (k *Kernel) Compile() error {
	log.Printf("tl2pure: compiling %d combinators", len(k.files))
	// add all declaration to check for name collisions
	for _, comb := range k.files {
		log.Printf("tl2pure: compiling %s", comb)
		kt := &KernelType{
			comb:      comb,
			instances: map[string]*TypeInstanceRef{},
		}
		if !comb.IsFunction {
			names := map[string]int{}
			for i, targ := range comb.TypeDecl.TemplateArguments {
				if _, ok := names[targ.Name]; ok {
					return fmt.Errorf("error adding type %s: template argument %s name collision", comb.TypeDecl.Name, targ.Name)
				}
				names[targ.Name] = i
			}
		}
		if err := k.addTip(kt); err != nil {
			return fmt.Errorf("error adding type %s: %w", comb.ReferenceName(), err)
		}
	}
	// type all declarations by comparing type ref with actual type definition
	for _, tip := range k.tipsOrdered {
		if tip.comb.IsFunction {
			if err := k.typeCheck(tip.comb.FuncDecl.ReturnType, nil); err != nil {
				return err
			}
			if err := k.typeCheckAliasFields(false, tlast.TL2TypeRef{}, tip.comb.FuncDecl.Arguments, nil); err != nil {
				return err
			}
			continue
		}
		if err := k.typeCheck(tip.comb.TypeDecl.Type, tip.comb.TypeDecl.TemplateArguments); err != nil {
			return err
		}
	}
	//instantiate all top-level declarations
	for _, tip := range k.tipsOrdered {
		if tip.comb.IsFunction {
			tr := tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: tip.comb.FuncDecl.Name}}
			if _, err := k.getInstance(tr); err != nil {
				return fmt.Errorf("error adding function %s: %w", tip.comb.FuncDecl.Name, err)
			}
			continue
		}
		typeDecl := tip.comb.TypeDecl
		if len(typeDecl.TemplateArguments) != 0 {
			continue // instantiate templates on demand only
		}
		tr := tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: typeDecl.Name}}
		if _, err := k.getInstance(tr); err != nil {
			return fmt.Errorf("error adding type %s: %w", typeDecl.Name, err)
		}
		//if len(tip.typeDecl.TemplateArguments) == 0 {
		//} else {
		//	for _, arg := range tip.typeDecl.TemplateArguments {
		//		if arg.Category.IsType() {
		//			someType := tlast.TL2TypeApplication{Name: tlast.TL2TypeName{Name: "uint32"}}
		//			tr.SomeType.Arguments = append(tr.SomeType.Arguments,
		//				tlast.TL2TypeArgument{Type: tlast.TL2TypeRef{SomeType: someType}})
		//		} else {
		//			tr.SomeType.Arguments = append(tr.SomeType.Arguments,
		//				tlast.TL2TypeArgument{IsNumber: true, Number: 1})
		//		}
		//	}
		//	// this does nothing, we actually need to have a function with a subset of getInstance and all check
		//	if _, err := k.resolveType(tr, nil, nil); err != nil {
		//		return fmt.Errorf("error checking template %s: %w", tip.typeDecl.Name, err)
		//	}
		//}
	}
	// check recursion cycles
	// TODO - why it is not possible to check all cycles before instantiation
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
