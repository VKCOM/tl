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
	WriteTL2(w []byte, optimizeEmpty bool, ctx *TL2Context) []byte
	ReadTL2(r []byte, ctx *TL2Context) ([]byte, error)
	WriteJSON(w []byte, ctx *TL2Context) []byte

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

// lrc is local resolve context, or actual values of template arguments, for pair<x,y> = ...; it could be <uint32, uint32>
func (k *Kernel) resolveType(tr tlast.TL2TypeRef, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []tlast.TL2TypeArgument) (tlast.TL2TypeRef, error) {
	ac, err := k.resolveArgument(tlast.TL2TypeArgument{Type: tr}, leftArgs, actualArgs)
	if err != nil {
		return tr, err
	}
	if ac.IsNumber {
		return tr, fmt.Errorf("type argument %s resolved to a number %d", tr, ac.Number)
	}
	return ac.Type, nil
}

func (k *Kernel) CanonicalName(t tlast.TL2TypeRef) string {
	sb := strings.Builder{}
	t.Print(&sb)
	return sb.String()
}

func (k *Kernel) IsBit(tr tlast.TL2TypeRef) bool {
	return !tr.IsBracket && tr.SomeType.Name.Namespace == "" && tr.SomeType.Name.Name == "bit"
}

func (k *Kernel) resolveArgument(tr tlast.TL2TypeArgument, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []tlast.TL2TypeArgument) (tlast.TL2TypeArgument, error) {
	before := tr
	was := before.Type.String()
	tr, err := k.resolveArgumentImpl(tr, leftArgs, actualArgs)
	after := before.Type.String()
	if was != after {
		panic(fmt.Sprintf("tl2pure: internal error, resolveArgument destroyed %s original value %s due to golang aliasing", after, was))
	}
	return tr, err
}

func (k *Kernel) resolveArgumentImpl(tr tlast.TL2TypeArgument, leftArgs []tlast.TL2TypeTemplate,
	actualArgs []tlast.TL2TypeArgument) (tlast.TL2TypeArgument, error) {
	// numbers resolve to numbers
	if tr.IsNumber {
		tr.Category = tlast.TL2TypeCategoryNat
		return tr, nil
	}
	if tr.Type.IsBracket {
		bracketType := *tr.Type.BracketType
		if bracketType.HasIndex {
			ic, err := k.resolveArgument(bracketType.IndexType, leftArgs, actualArgs)
			if err != nil {
				return tr, err
			}
			bracketType.IndexType = ic
		}
		ac, err := k.resolveType(bracketType.ArrayType, leftArgs, actualArgs)
		if err != nil {
			return tr, err
		}
		bracketType.ArrayType = ac
		tr.Type.BracketType = &bracketType
		return tr, nil
	}
	// names found in local arguments have priprity over global type names
	someType := tr.Type.SomeType
	if someType.Name.Namespace == "" {
		for i, targ := range leftArgs {
			if targ.Name == someType.Name.Name {
				if len(someType.Arguments) != 0 {
					return tr, fmt.Errorf("reference to template argument %s cannot have arguments", targ.Name)
				}
				return actualArgs[i], nil
			}
		}
		// probably ref to global type or a typo
	}
	//return tr, nil
	someType.Arguments = append([]tlast.TL2TypeArgument{}, someType.Arguments...) // preserve original
	for i, arg := range someType.Arguments {
		rt, err := k.resolveArgument(arg, leftArgs, actualArgs)
		if err != nil {
			return tr, err
		}
		someType.Arguments[i] = rt
	}
	tr.Type.SomeType = someType
	return tr, nil
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
		if err := k.addTip(kt); err != nil {
			return fmt.Errorf("error adding type %s: %w", comb.TypeDecl.Name, err)
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
