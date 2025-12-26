package tl2pure

import (
	"fmt"
	"log"

	"github.com/vkcom/tl/internal/tlast"
)

func (k *Kernel) getInstance(tr tlast.TL2TypeRef) (*TypeInstanceRef, error) {
	canonicalName := tr.String()
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
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument) (TypeInstance, error) {

	if len(actualArgs) != len(leftArgs) {
		return nil, fmt.Errorf("typeref to %s must have %d template arguments, has %d", canonicalName, len(leftArgs), len(actualArgs))
	}
	for i, targ := range leftArgs {
		arg := actualArgs[i]
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
		return k.createUnion(canonicalName, definition.UnionType, leftArgs, actualArgs)
	case definition.IsAlias():
		return k.createAlias(canonicalName, definition.TypeAlias, leftArgs, actualArgs)
	case definition.IsConstructorFields:
		return k.createObject(canonicalName,
			true, definition.TypeAlias, definition.ConstructorFields,
			leftArgs, actualArgs,
			false, 0, nil)
	default:
		return nil, fmt.Errorf("wrong type classification, internal error %s", canonicalName)
	}
}
