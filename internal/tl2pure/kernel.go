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
	// if tr.IsNumber {
	//	return k.resolveArgumentImpl(tr, templateArguments, lrc)
	// }
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
		tr.Category = tlast.TL2TypeCategoryNat
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
	if someType.Name.Namespace == "" {
		for i, targ := range templateArguments {
			if targ.Name == someType.Name.Name {
				if len(someType.Arguments) != 0 {
					return tr, fmt.Errorf("reference to template argument %s cannot have arguments", targ.Name)
				}
				return lrc[i], nil
			}
		}
		// probably ref to global type or a typo
	}
	//return tr, nil
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

func (k *Kernel) getInstance(tr tlast.TL2TypeRef) (*TypeInstanceRef, error) {
	canonicalName := k.CanonicalName(tr)
	if ref, ok := k.instances[canonicalName]; ok {
		return ref, nil
	}
	if tr.IsBracket {
		log.Printf("creating a bracket instance of type %s", canonicalName)
		// must store pointer before children getInstance() terminates recursion
		// this instance stays not initialized in case of error, but kernel then is not consistent anyway
		ref := k.addInstance(canonicalName, k.brackets)

		elemInstance, err := k.getInstance(tr.BracketType.ArrayType)
		if err != nil {
			return nil, err
		}
		if tr.BracketType.HasIndex {
			if tr.BracketType.IndexType.IsNumber {
				// tuple
				ref.ins = k.createTupleVector(canonicalName, true, tr.BracketType.IndexType.Number, elemInstance)
				return ref, nil
			}
			// map
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
		ref.ins = k.createTupleVector(canonicalName, false, 0, elemInstance)
		return ref, nil
	}
	log.Printf("creating an instance of type %s", canonicalName)
	// must store pointer before children getInstance() terminates recursion
	// this instance stays mpt initialized in case of error, but kernel then is not consistent anyway
	someType := tr.SomeType
	kt, ok := k.tips[someType.Name]
	if !ok {
		return nil, fmt.Errorf("type %s does not exist", someType.Name)
	}
	// must store pointer before children getInstance() terminates recursion
	// this instance stays not initialized in case of error, but kernel then is not consistent anyway
	ref := k.addInstance(canonicalName, kt)

	var err error
	if !kt.comb.IsFunction {
		ref.ins, err = k.createOrdinaryType(canonicalName, kt.comb.TypeDecl.Type, kt.comb.TypeDecl.TemplateArguments, someType.Arguments)
		if err != nil {
			return nil, err
		}
		return ref, nil
	}
	funcDecl := kt.comb.FuncDecl
	resultType, err := k.createOrdinaryType(canonicalName, funcDecl.ReturnType, nil, nil)
	if err != nil {
		return nil, err
	}
	ref.ins, err = k.createObject(canonicalName, true,
		tlast.TL2TypeRef{}, funcDecl.Arguments, nil, nil, false, 0,
		resultType)
	if err != nil {
		return nil, err
	}
	return ref, nil
}

// alias || fields || union
func (k *Kernel) createOrdinaryType(canonicalName string, definition tlast.TL2TypeDefinition,
	templateArguments []tlast.TL2TypeTemplate, lrc []tlast.TL2TypeArgument) (TypeInstance, error) {

	if len(lrc) != len(templateArguments) {
		return nil, fmt.Errorf("typeref to %s must have %d template arguments, has %d", canonicalName, len(templateArguments), len(lrc))
	}
	for i, targ := range templateArguments {
		arg := lrc[i]
		if targ.Category.IsUint32() != arg.IsNumber {
			return nil, fmt.Errorf("typeref %s argument %s category differ", canonicalName, targ.Name)
		}
		// if arg.IsNumber {
		//	continue
		// }
		// if some arguments are unused inside body, they will not be instantiated and checked, this is ok
		// _, err := k.getInstance(arg.Type)
		// if err != nil {
		//	return nil, err
		// }
	}

	// lrc2 := map[string]ResolvedArgument{} // internal context
	// for i, resolvedArg := range resolvedArgs {
	// targ := kt.typeDecl.TemplateArguments[i]
	// if _, ok := lrc2[targ.Name]; ok {
	// return nil, fmt.Errorf("typeref %s template parameter %s name collision", ktr.Name, targ.Name)
	// }
	// lrc2[targ.Name] = resolvedArg
	// }
	switch {
	case definition.IsUnionType:
		return k.createUnion(canonicalName, definition.UnionType, templateArguments, lrc)
	case definition.IsAlias():
		return k.createAlias(canonicalName, definition.TypeAlias, templateArguments, lrc)
	case definition.IsConstructorFields:
		return k.createObject(canonicalName,
			true, definition.TypeAlias, definition.ConstructorFields,
			templateArguments, lrc,
			false, 0, nil)
	default:
		return nil, fmt.Errorf("wrong type classification, internal error %s", canonicalName)
	}
}

func (k *Kernel) typeCheck(tip tlast.TL2TypeDefinition, leftArguments []tlast.TL2TypeTemplate) error {
	if tip.IsUnionType {
		for _, v := range tip.UnionType.Variants {
			if err := k.typeCheckAliasFields(v.IsTypeAlias, v.TypeAlias, v.Fields, leftArguments); err != nil {
				return err
			}
		}
		return nil
	}
	if tip.IsAlias() {
		// does not work before resolving type, for example identity<t:type> = t;
		//	aliasBit := k.IsBit(typeDecl.TypeAlias)
		//	if aliasBit {
		//		return fmt.Errorf("type bit is not allowed as a type alias")
		//	}
		return k.typeCheckTypeRef(tip.TypeAlias, leftArguments)
	}
	return k.typeCheckAliasFields(false, tlast.TL2TypeRef{},
		tip.ConstructorFields, leftArguments)
}

func (k *Kernel) typeCheckAliasFields(isTypeAlias bool, typeAlias tlast.TL2TypeRef,
	fields []tlast.TL2Field, leftArguments []tlast.TL2TypeTemplate) error {
	if isTypeAlias {
		cat, err := k.typeCheckArgument(tlast.TL2TypeArgument{Type: typeAlias}, leftArguments)
		if err != nil {
			return err
		}
		if !cat.IsType() {
			return fmt.Errorf("type alias cannot be number")
		}
		return nil
	}
	for _, f := range fields {
		cat, err := k.typeCheckArgument(tlast.TL2TypeArgument{Type: f.Type}, leftArguments)
		if err != nil {
			return err
		}
		if !cat.IsType() {
			return fmt.Errorf("field type %s cannot be number", f.Name)
		}
	}
	return nil
}

func (k *Kernel) typeCheckArgument(arg tlast.TL2TypeArgument, leftArguments []tlast.TL2TypeTemplate) (tlast.TL2TypeCategory, error) {
	if arg.IsNumber {
		return tlast.TL2TypeCategoryNat, nil
	}
	if !arg.Type.IsBracket && arg.Type.SomeType.Name.Namespace == "" {
		for _, la := range leftArguments {
			if arg.Type.SomeType.Name.Name == la.Name {
				if len(arg.Type.SomeType.Arguments) != 0 {
					return "", fmt.Errorf("reference to template argument %s cannot have arguments", la.Name)
				}
				return la.Category, nil
			}
		}
		// reference to global type
	}
	if err := k.typeCheckTypeRef(arg.Type, leftArguments); err != nil {
		return "", err
	}
	return tlast.TL2TypeCategoryType, nil
}

func (k *Kernel) typeCheckTypeRef(tr tlast.TL2TypeRef, leftArguments []tlast.TL2TypeTemplate) error {
	if tr.IsBracket {
		cat, err := k.typeCheckArgument(tlast.TL2TypeArgument{Type: tr.BracketType.ArrayType}, leftArguments)
		if err != nil {
			return err
		}
		if !cat.IsType() {
			return fmt.Errorf("bracket element type cannot be number")
		}
		if tr.BracketType.HasIndex {
			_, err = k.typeCheckArgument(tr.BracketType.IndexType, leftArguments)
			if err != nil {
				return err
			}
			// can be both type (map) and number (tuple)
		}
		return nil
	}
	someType := tr.SomeType
	kt, ok := k.tips[someType.Name]
	if !ok {
		return fmt.Errorf("type %s does not exist", someType.Name)
	}
	if kt.comb.IsFunction {
		return fmt.Errorf("cannot reference function %s", someType.Name)
	}
	if len(someType.Arguments) != len(kt.comb.TypeDecl.TemplateArguments) {
		return fmt.Errorf("typeref %s must have %d template arguments, has %d", someType.String(), len(kt.comb.TypeDecl.TemplateArguments), len(someType.Arguments))
	}
	for i, targ := range kt.comb.TypeDecl.TemplateArguments {
		cat, err := k.typeCheckArgument(someType.Arguments[i], leftArguments)
		if err != nil {
			return err
		}
		if targ.Category != cat {
			return fmt.Errorf("typeref %s argument %s wrong category, must be %s", someType.String(), targ.Name, targ.Category)
		}
	}
	return nil
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
