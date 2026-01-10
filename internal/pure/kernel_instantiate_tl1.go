package pure

import (
	"fmt"
	"log"

	"github.com/vkcom/tl/internal/tlast"
)

type LocalArg struct {
	wrongTypeErr error // we must add all field names to local context, because they must correctly shadow names outside, but we check the type
	arg          tlast.ArithmeticOrType
	natArgs      []ActualNatArg
}

func (k *Kernel) resolveTypeTL1(tr tlast.TypeRef, leftArgs []tlast.TemplateArgument,
	actualArgs []LocalArg) (tlast.TypeRef, []ActualNatArg, error) {
	ac, natArgs, err := k.resolveArgumentTL1(tlast.ArithmeticOrType{IsArith: false, T: tr}, leftArgs, actualArgs)
	if err != nil {
		return tr, nil, err
	}
	if ac.IsArith {
		return tr, nil, fmt.Errorf("type argument %s resolved to a number %d", tr, ac.Arith.Res)
	}
	if ac.T.String() == "*" {
		return tr, nil, fmt.Errorf("type argument %s resolved to a nat argument %s", tr, ac.T)
	}
	return ac.T, natArgs, nil
}

func (k *Kernel) resolveArgumentTL1(tr tlast.ArithmeticOrType, leftArgs []tlast.TemplateArgument,
	actualArgs []LocalArg) (tlast.ArithmeticOrType, []ActualNatArg, error) {
	before := tr
	was := before.T.String()
	tr, natArgs, err := k.resolveArgumentTL1Impl(tr, leftArgs, actualArgs)
	after := before.T.String()
	if was != after {
		panic(fmt.Sprintf("tl2pure: internal error, resolveArgumentTL1 destroyed %s original value %s due to golang aliasing", after, was))
	}
	return tr, natArgs, err
}

func (k *Kernel) resolveArgumentTL1Impl(tr tlast.ArithmeticOrType, leftArgs []tlast.TemplateArgument,
	actualArgs []LocalArg) (tlast.ArithmeticOrType, []ActualNatArg, error) {
	if tr.IsArith {
		return tr, nil, nil
	}
	// names found in local arguments have priority over global type names
	if tr.T.Type.Namespace == "" {
		for i, targ := range leftArgs {
			if targ.FieldName == tr.T.Type.Name {
				if len(tr.T.Args) != 0 {
					return tr, nil, fmt.Errorf("reference to template argument %s cannot have arguments", targ.FieldName)
				}
				actualArg := actualArgs[i]
				if actualArg.wrongTypeErr != nil {
					return tr, nil, fmt.Errorf("reference to field %s impossible, must have # type", targ.FieldName)
				}
				return actualArg.arg, actualArg.natArgs, nil
			}
		}
		// probably ref to global type or a typo
	}
	tr.T.Args = append([]tlast.ArithmeticOrType{}, tr.T.Args...) // preserve original
	var natArgs []ActualNatArg
	for i, arg := range tr.T.Args {
		rt, natArgs2, err := k.resolveArgumentTL1(arg, leftArgs, actualArgs)
		if err != nil {
			return tr, nil, err
		}
		tr.T.Args[i] = rt
		natArgs = append(natArgs, natArgs2...)
	}
	return tr, natArgs, nil
}

/*
	func (k *Kernel) getTL1CanonicalString(tr tlast.TypeRef) string {
		var sb strings.Builder
		k.getTL1CanonicalStringImpl(&sb, tr)
		return sb.String()
	}

	func (k *Kernel) getTL1CanonicalStringImpl(sb *strings.Builder, tr tlast.TypeRef) {
		if tr.Type.Namespace != "" {
			sb.WriteString(tr.Type.Namespace)
			sb.WriteString(".")
		}
		sb.WriteString(tr.Type.Name)
		if len(tr.Args) == 0 {
			return
		}
		sb.WriteString("<")
		for i, arg := range tr.Args {
			if i != 0 {
				sb.WriteString(",")
			}
			if arg.IsArith {
				sb.Write(strconv.AppendUint(nil, uint64(arg.Arith.Res), 10))
				continue
			}
			if arg.IsTNatRef {
				sb.WriteString("*")
				continue
			}
			k.getTL1CanonicalStringImpl(sb, arg.T)
		}
		sb.WriteString(">")
	}
*/
func (k *Kernel) getInstanceTL1(tr tlast.TypeRef) (*TypeInstanceRef, error) {
	canonicalName := tr.String()
	// we remap TL1 type names into TL2 type names here.
	// hopefully, this does not cause too many irregularities downstream
	switch canonicalName {
	case "int":
		canonicalName = "int32"
	case "long":
		canonicalName = "int64"
	case "#":
		canonicalName = "uint32"
	}
	if ref, ok := k.instances[canonicalName]; ok {
		return ref, nil
	}
	//if tr.Type.String() == "" {
	//	log.Printf("creating a bracket instance of type %s", canonicalName)
	//	// must store pointer before children getInstance() terminates recursion
	//	// this instance stays not initialized in case of error, but kernel then is not consistent anyway
	//	ref := k.addInstance(canonicalName, k.brackets)
	//
	//	elemInstance, err := k.getInstance(tr.BracketType.ArrayType)
	//	if err != nil {
	//		return nil, err
	//	}
	//	if tr.BracketType.HasIndex {
	//		if tr.BracketType.IndexType.IsNumber {
	//			// tuple
	//			ref.ins = k.createArray(canonicalName, true, tr.BracketType.IndexType.Number, elemInstance)
	//			return ref, nil
	//		}
	//		// dict
	//		keyInstance, err := k.getInstance(tr.BracketType.IndexType.Type)
	//		if err != nil {
	//			return nil, err
	//		}
	//		if !keyInstance.ins.GoodForMapKey() {
	//			return nil, fmt.Errorf("type %s is not allowed as a map key (only 'bool', integers and 'string' allowed)", keyInstance.ins.CanonicalName())
	//		}
	//		ref.ins = k.createDict(canonicalName, keyInstance, elemInstance)
	//		return ref, nil
	//	}
	//	// vector
	//	ref.ins = k.createArray(canonicalName, false, 0, elemInstance)
	//	return ref, nil
	//}
	log.Printf("creating an instance of type %s", canonicalName)
	// must store pointer before children getInstance() terminates recursion
	// this instance stays mpt initialized in case of error, but kernel then is not consistent anyway
	kt, ok := k.tips[tr.Type.String()]
	if !ok {
		return nil, fmt.Errorf("type %s does not exist", tr.Type)
	}
	// must store pointer before children getInstance() terminates recursion
	// this instance stays not initialized in case of error, but kernel then is not consistent anyway
	ref := k.addInstance(canonicalName, kt)

	var err error
	if kt.originTL2 {
		return nil, fmt.Errorf("TL1 combinator cannot reference TL2 combinator %s", tr.Type)
	}
	comb := kt.combTL1[0]
	if !comb.IsFunction {
		if len(comb.TemplateArguments) != len(tr.Args) {
			return nil, fmt.Errorf("typeref to %s must have %d template arguments, has %d", canonicalName, len(comb.TemplateArguments), len(tr.Args))
		}
		ref.ins, err = k.createOrdinaryTypeTL1FromTL1(canonicalName, kt.combTL1, comb.TemplateArguments, tr.Args)
		if err != nil {
			return nil, err
		}
		return ref, nil
	}
	return nil, fmt.Errorf("TODO - function from TL1 not yet supported")
}

func (k *Kernel) createOrdinaryTypeTL1FromTL2(canonicalName string, definition []*tlast.Combinator,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.TL2TypeArgument) (TypeInstance, error) {

	switch {
	//case len(definition) > 1:
	//	return k.createUnion(canonicalName, definition.UnionType, leftArgs, actualArgs)
	//case definition[0].IsAlias():
	//	return k.createAlias(canonicalName, definition.TypeAlias, leftArgs, actualArgs)
	case len(definition) == 1:
		return k.createStructTL1FromTL2(canonicalName,
			definition[0].Fields,
			leftArgs, actualArgs,
			false, 0, nil)
	default:
		return nil, fmt.Errorf("wrong type classification, internal error %s", canonicalName)
	}
}

func (k *Kernel) createOrdinaryTypeTL1FromTL1(canonicalName string, definition []*tlast.Combinator,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (TypeInstance, error) {

	switch {
	//case len(definition) > 1:
	//	return k.createUnion(canonicalName, definition.UnionType, leftArgs, actualArgs)
	//case definition[0].IsAlias():
	//	return k.createAlias(canonicalName, definition.TypeAlias, leftArgs, actualArgs)
	case len(definition) == 1:
		return k.createStructTL1FromTL1(canonicalName,
			definition[0].Fields,
			leftArgs, actualArgs,
			false, 0, nil)
	default:
		return nil, fmt.Errorf("wrong type classification, internal error %s", canonicalName)
	}
}
