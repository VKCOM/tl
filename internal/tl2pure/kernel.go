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
	scratch []byte
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
	tip tlast.TL2TypeDeclaration
	// index by canonical name
	instances map[string]*TypeInstanceRef
}

type Kernel struct {
	tips         map[tlast.TL2TypeName]*KernelType
	tipsOrdered  []*KernelType
	tipsTopLevel []*KernelType

	instances        map[string]*TypeInstanceRef
	instancesOrdered []*TypeInstanceRef

	files []tlast.TL2Combinator
}

// Add builtin types
func NewKernel() *Kernel {
	k := &Kernel{
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

	return k
}

func (k *Kernel) addTip(kt *KernelType) error {
	_, ok := k.tips[kt.tip.Name]
	if ok {
		return fmt.Errorf("type %v already exists", kt.tip.Name)
	}
	k.tips[kt.tip.Name] = kt
	k.tipsOrdered = append(k.tipsOrdered, kt)
	if len(kt.tip.TemplateArguments) == 0 {
		k.tipsTopLevel = append(k.tipsTopLevel, kt)
	}
	return nil
}

// lrc is local resolve context, or actual values of template arguments, for pair<x,y> = ...; it could be <uint32, uint32>
func (k *Kernel) resolveType(tr tlast.TL2TypeRef, templateArguments []tlast.TL2TypeTemplate,
	lrc []tlast.TL2TypeArgument) (tlast.TL2TypeRef, error) {
	ac, err := k.resolveArgument(tlast.TL2TypeArgument{Type: tr}, templateArguments, lrc)
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
	k.canonicalName(t, &sb)
	return sb.String()
}

func (k *Kernel) canonicalName(t tlast.TL2TypeRef, sb *strings.Builder) {
	if t.IsBracket {
		sb.WriteString("[")
		if t.BracketType.HasIndex {
			if t.BracketType.IndexType.IsNumber {
				sb.WriteString(fmt.Sprintf("%d", t.BracketType.IndexType.Number))
			} else {
				k.canonicalName(t.BracketType.IndexType.Type, sb)
			}
		}
		sb.WriteString("]")
		k.canonicalName(t.BracketType.ArrayType, sb)
		return
	}
	if t.SomeType.Name.Namespace != "" { // t.Name.String() unwrapped for efficiency
		sb.WriteString(t.SomeType.Name.Namespace)
		sb.WriteString(".")
	}
	sb.WriteString(t.SomeType.Name.Name)
	for i, arg := range t.SomeType.Arguments {
		if i == 0 {
			sb.WriteString("<")
		} else {
			sb.WriteString(",")
		}
		if arg.IsNumber {
			sb.WriteString(fmt.Sprintf("%d", arg.Number))
		} else {
			k.canonicalName(arg.Type, sb)
		}
	}
	if len(t.SomeType.Arguments) != 0 {
		sb.WriteString(">")
	}
}

func (k *Kernel) IsBit(tr tlast.TL2TypeRef) bool {
	return !tr.IsBracket && tr.SomeType.Name.Namespace == "" && tr.SomeType.Name.Name == "bit"
}

func (k *Kernel) resolveArgument(tr tlast.TL2TypeArgument, templateArguments []tlast.TL2TypeTemplate,
	lrc []tlast.TL2TypeArgument) (tlast.TL2TypeArgument, error) {
	if tr.IsNumber {
		return k.resolveArgumentImpl(tr, templateArguments, lrc)
	}
	before := tr
	was := k.CanonicalName(before.Type)
	tr, err := k.resolveArgumentImpl(tr, templateArguments, lrc)
	now := k.CanonicalName(before.Type)
	if was != now {
		panic(fmt.Sprintf("tl2pure: internal error, resolveArgument destroyed %s original value %s due to golang aliasing", now, was))
	}
	return tr, err
}

func (k *Kernel) resolveArgumentImpl(tr tlast.TL2TypeArgument, templateArguments []tlast.TL2TypeTemplate,
	lrc []tlast.TL2TypeArgument) (tlast.TL2TypeArgument, error) {
	// numbers resolve to numbers
	if tr.IsNumber {
		return tr, nil
	}
	if tr.Type.IsBracket {
		bracketType := *tr.Type.BracketType
		if bracketType.HasIndex {
			ic, err := k.resolveArgument(bracketType.IndexType, templateArguments, lrc)
			if err != nil {
				return tr, err
			}
			bracketType.IndexType = ic
		}
		ac, err := k.resolveType(bracketType.ArrayType, templateArguments, lrc)
		if err != nil {
			return tr, err
		}
		bracketType.ArrayType = ac
		tr.Type.BracketType = &bracketType
		return tr, nil
	}
	// names found in local arguments have priprity over global type names
	someType := tr.Type.SomeType
	if len(someType.Arguments) == 0 {
		if someType.Name.Namespace == "" {
			for i, targ := range templateArguments {
				if targ.Name == someType.Name.Name {
					return lrc[i], nil
				}
			}
			// probably ref to global type or a typo
		}
		return tr, nil
	}
	someType.Arguments = append([]tlast.TL2TypeArgument{}, someType.Arguments...) // preserve original
	for i, arg := range someType.Arguments {
		rt, err := k.resolveArgument(arg, templateArguments, lrc)
		if err != nil {
			return tr, err
		}
		someType.Arguments[i] = rt
	}
	tr.Type.SomeType = someType
	return tr, nil
}

func (k *Kernel) getInstance(tr tlast.TL2TypeRef) (*TypeInstanceRef, error) {
	canonicalName := k.CanonicalName(tr)
	if ref, ok := k.instances[canonicalName]; ok {
		return ref, nil
	}
	if tr.IsBracket {
		log.Printf("creating a bracket instance of type %s", canonicalName)
		ref := &TypeInstanceRef{}
		k.instances[canonicalName] = ref // storing pointer terminates recursion
		k.instancesOrdered = append(k.instancesOrdered, ref)

		elemBit := k.IsBit(tr.BracketType.ArrayType) // we must not call anything on TypeInstance during recursive resolution

		elemInstance, err := k.getInstance(tr.BracketType.ArrayType)
		if err != nil {
			return nil, err
		}
		if tr.BracketType.HasIndex {
			if tr.BracketType.IndexType.IsNumber {
				if elemBit {
					return nil, fmt.Errorf("type bit is not allowed as bracket element")
				}
				ref.ins = k.createTupleVector(canonicalName, true, tr.BracketType.IndexType.Number, elemInstance)
				// tuple
				return ref, nil
			}
			// map, bit is allowed as an element
			keyInstance, err := k.getInstance(tr.BracketType.IndexType.Type)
			if err != nil {
				return nil, err
			}
			if !keyInstance.ins.GoodForMapKey() {
				return nil, fmt.Errorf("type %s is not allowed as a map key (only 'bool', integers and 'string' allowed)", keyInstance.ins.CanonicalName())
			}
			ref.ins = k.createMap(canonicalName, keyInstance, elemInstance)
			return ref, nil
		}
		// vector
		if elemBit {
			return nil, fmt.Errorf("type bit is not allowed as bracket element")
		}
		ref.ins = k.createTupleVector(canonicalName, false, 0, elemInstance)
		return ref, nil
	}
	log.Printf("creating an instance of type %s", canonicalName)
	someType := tr.SomeType
	kt, ok := k.tips[someType.Name]
	if !ok {
		return nil, fmt.Errorf("type %s does not exist", someType.Name)
	}
	if _, ok := kt.instances[canonicalName]; ok {
		panic("type instance list contains duplicate")
	}
	ref := &TypeInstanceRef{}
	kt.instances[canonicalName] = ref

	k.instances[canonicalName] = ref // storing pointer terminates recursion
	k.instancesOrdered = append(k.instancesOrdered, ref)

	if len(someType.Arguments) != len(kt.tip.TemplateArguments) {
		return nil, fmt.Errorf("typeref to %s must have %d template arguments, has %d", someType.Name, len(kt.tip.TemplateArguments), len(someType.Arguments))
	}
	for i, arg := range someType.Arguments {
		targ := kt.tip.TemplateArguments[i]
		if targ.Category.IsUint32() != arg.IsNumber {
			return nil, fmt.Errorf("typeref %s argument %s category differ", someType.Name, targ.Name)
		}
		if arg.IsNumber {
			continue
		}
		// if some arguments are unused inside body, they will not be instantiated and checked
		_, err := k.getInstance(arg.Type)
		if err != nil {
			return nil, err
		}
	}

	// lrc2 := map[string]ResolvedArgument{} // internal context
	// for i, resolvedArg := range resolvedArgs {
	// targ := kt.tip.TemplateArguments[i]
	// if _, ok := lrc2[targ.Name]; ok {
	// return nil, fmt.Errorf("typeref %s template parameter %s name collision", ktr.Name, targ.Name)
	// }
	// lrc2[targ.Name] = resolvedArg
	// }
	var err error
	switch {
	case kt.tip.Type.IsUnionType:
		ref.ins, err = k.createUnion(canonicalName, kt.tip.Type.UnionType, kt.tip.TemplateArguments, someType.Arguments)
	case kt.tip.Type.IsAlias():
		ref.ins, err = k.createAlias(canonicalName, kt.tip.Type.TypeAlias, kt.tip.TemplateArguments, someType.Arguments)
	case kt.tip.Type.IsConstructorFields:
		ref.ins, err = k.createObject(canonicalName, kt.tip,
			true, kt.tip.Type.TypeAlias, kt.tip.Type.ConstructorFields,
			kt.tip.TemplateArguments, someType.Arguments,
			false, 0)
	default:
		return nil, fmt.Errorf("wrong type classification, internal error %s", canonicalName)
	}
	if err != nil {
		return nil, err
	}
	return ref, nil
}

func (k *Kernel) typeCheck(tip tlast.TL2TypeDefinition) error {
	if tip.IsUnionType {
		for _, v := range tip.UnionType.Variants {
			if err := k.typeCheckAliasFields(v.IsTypeAlias, v.TypeAlias, v.Fields); err != nil {
				return err
			}
		}
		return nil
	}
	if tip.IsAlias() {
		// does not work before resolving type, for example identity<t:type> = t;
		//	aliasBit := k.IsBit(tip.TypeAlias)
		//	if aliasBit {
		//		return fmt.Errorf("type bit is not allowed as a type alias")
		//	}
		return k.typeCheckTypeRef(tip.TypeAlias)
	}
	return k.typeCheckAliasFields(false, tlast.TL2TypeRef{}, tip.ConstructorFields)
}

func (k *Kernel) typeCheckAliasFields(isTypeAlias bool, typeAlias tlast.TL2TypeRef, fields []tlast.TL2Field) error {
	if isTypeAlias {
		return k.typeCheckTypeRef(typeAlias)
	}
	for _, f := range fields {
		if err := k.typeCheckTypeRef(f.Type); err != nil {
			return err
		}
	}
	return nil
}

func (k *Kernel) typeCheckTypeRef(tr tlast.TL2TypeRef) error {
	return nil
	// if tr.IsBracket {
	// 	// type check tr.BracketType.ArrayType
	// 	elemInstance, err := k.getInstance(tr.BracketType.ArrayType)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if tr.BracketType.IndexType != nil {
	// 		if tr.BracketType.IndexType.IsNumber {
	// 			if elemBit {
	// 				return nil, fmt.Errorf("type bit is not allowed as bracket element")
	// 			}
	// 			ref.ins = k.createTupleVector(canonicalName, true, tr.BracketType.IndexType.Number, elemInstance)
	// 			// tuple
	// 			return ref, nil
	// 		}
	// 		// map, bit is allowed as an element
	// 		keyInstance, err := k.getInstance(tr.BracketType.IndexType.Type)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		if !keyInstance.ins.GoodForMapKey() {
	// 			return nil, fmt.Errorf("type %s is not allowed as a map key (only 'bool', integers and 'string' allowed)", keyInstance.ins.CanonicalName())
	// 		}
	// 		ref.ins = k.createMap(canonicalName, keyInstance, elemInstance)
	// 		return ref, nil
	// 	}
	// 	// vector
	// 	if elemBit {
	// 		return nil, fmt.Errorf("type bit is not allowed as bracket element")
	// 	}
	// 	ref.ins = k.createTupleVector(canonicalName, false, 0, elemInstance)
	// 	return ref, nil
	// }
	// log.Printf("creating an instance of type %s", canonicalName)
	// someType := tr.SomeType
	// kt, ok := k.tips[someType.Name]
	// if !ok {
	// 	return nil, fmt.Errorf("type %s does not exist", someType.Name)
	// }
	// if _, ok := kt.instances[canonicalName]; ok {
	// 	panic("type instance list contains duplicate")
	// }
	// ref := &TypeInstanceRef{}
	// kt.instances[canonicalName] = ref

	// k.instances[canonicalName] = ref // storing pointer terminates recursion
	// k.instancesOrdered = append(k.instancesOrdered, ref)

	// if len(someType.Arguments) != len(kt.tip.TemplateArguments) {
	// 	return nil, fmt.Errorf("typeref to %s must have %d template arguments, has %d", someType.Name, len(kt.tip.TemplateArguments), len(someType.Arguments))
	// }
	// for i, arg := range someType.Arguments {
	// 	targ := kt.tip.TemplateArguments[i]
	// 	if targ.Category.IsUint32() != arg.IsNumber {
	// 		return nil, fmt.Errorf("typeref %s argument %s category differ", someType.Name, targ.Name)
	// 	}
	// 	if arg.IsNumber {
	// 		continue
	// 	}
	// 	// if some arguments are unused inside body, they will not be instantiated and checked
	// 	_, err := k.getInstance(arg.Type)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// // lrc2 := map[string]ResolvedArgument{} // internal context
	// // for i, resolvedArg := range resolvedArgs {
	// // targ := kt.tip.TemplateArguments[i]
	// // if _, ok := lrc2[targ.Name]; ok {
	// // return nil, fmt.Errorf("typeref %s template parameter %s name collision", ktr.Name, targ.Name)
	// // }
	// // lrc2[targ.Name] = resolvedArg
	// // }
	// var err error
	// switch {
	// // case kt.tip.Type.IsUnionType:
	// // ref.ins, err = k.createUnion(canonicalName, kt.tip.Type.UnionType, lrc2)
	// case kt.tip.Type.IsAlias():
	// 	ref.ins, err = k.createAlias(canonicalName, kt.tip.Type.TypeAlias, kt.tip.TemplateArguments, someType.Arguments)
	// case kt.tip.Type.IsConstructorFields:
	// 	ref.ins, err = k.createObject(canonicalName, kt.tip,
	// 		true, kt.tip.Type.TypeAlias, kt.tip.Type.ConstructorFields,
	// 		kt.tip.TemplateArguments, someType.Arguments,
	// 		false, 0)
	// default:
	// 	return nil, fmt.Errorf("wrong type classification, internal error %s", canonicalName)
	// }
	// if err != nil {
	// 	return nil, err
	// }
	// return ref, nil
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
		if comb.IsFunction {
			continue
		}
		kt := &KernelType{
			tip:       comb.TypeDecl,
			instances: map[string]*TypeInstanceRef{},
		}
		if err := k.addTip(kt); err != nil {
			return fmt.Errorf("error adding type %s: %w", comb.TypeDecl.Name, err)
		}
	}
	// type all declarations by comparing type ref with actual type definition
	for _, tip := range k.tipsOrdered {
		if err := k.typeCheck(tip.tip.Type); err != nil {
			return err
		}
	}
	// instantiate all top-level declarations
	for _, tip := range k.tipsOrdered {
		tr := tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: tip.tip.Name}}
		if len(tip.tip.TemplateArguments) == 0 {
			if _, err := k.getInstance(tr); err != nil {
				return fmt.Errorf("error adding type %s: %w", tip.tip.Name, err)
			}
		} else {
			for _, arg := range tip.tip.TemplateArguments {
				if arg.Category.IsType() {
					someType := tlast.TL2TypeApplication{Name: tlast.TL2TypeName{Name: "uint32"}}
					tr.SomeType.Arguments = append(tr.SomeType.Arguments,
						tlast.TL2TypeArgument{Type: tlast.TL2TypeRef{SomeType: someType}})
				} else {
					tr.SomeType.Arguments = append(tr.SomeType.Arguments,
						tlast.TL2TypeArgument{IsNumber: true, Number: 1})
				}
			}
			// this does nothing, we actually need to have a function with a subset of getInstance and all check
			if _, err := k.resolveType(tr, nil, nil); err != nil {
				return fmt.Errorf("error checking template %s: %w", tip.tip.Name, err)
			}
		}
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
