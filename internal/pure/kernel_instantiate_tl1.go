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
	"strconv"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
)

type LocalArg struct {
	wrongTypeErr error // we must add all field names to local context, because they must correctly shadow names outside, but we check the type
	arg          tlast.ArithmeticOrType
	natArgs      []ActualNatArg
}

// we remap TL1 type names into TL2 type names here.
// hopefully, this does not cause too many irregularities downstream
func (k *Kernel) replaceTL1BuiltinName(canonicalName string) string {
	canonicalName = strings.TrimPrefix(canonicalName, "%")
	switch canonicalName {
	case "int":
		canonicalName = "int32"
	case "long":
		canonicalName = "int64"
	case "#":
		canonicalName = "uint32"
	}
	return canonicalName
}

// top level types do not have bare/boxed in their names
func (k *Kernel) canonicalStringTL1(tr tlast.TypeRef, top bool) string {
	var s strings.Builder

	tName := tr.Type.String()
	switch tName {
	case "__vector", "__tuple":
		s.WriteString(tName)
	case "int":
		s.WriteString("int32")
	case "long":
		s.WriteString("int64")
	case "#":
		s.WriteString("uint32")
	default:
		kt, ok := k.tips[tName]
		if !ok {
			panic(fmt.Sprintf("canonical string tip %s not found", tName))
		}
		if kt.originTL2 {
			panic(fmt.Sprintf("canonical string tip %s not from TL1", tName))
		}
		if len(kt.combTL1) > 1 {
			if tr.Bare {
				panic(fmt.Sprintf("canonical string tip %s bare union", tName))
			}
			if tName == "Bool" {
				s.WriteString("bool")
			} else {
				kt.combTL1[0].TypeDecl.Name.WriteString(&s)
			}
		} else if len(kt.combTL1) == 1 {
			if tr.Bare || utils.ToLowerFirst(tr.Type.Name) == tr.Type.Name {
				kt.combTL1[0].Construct.Name.WriteString(&s)
			} else {
				if !top {
					s.WriteString("+")
				}
				kt.combTL1[0].Construct.Name.WriteString(&s)
			}
		} else {
			panic("all builtins are parsed from TL text, so must have exactly one constructor")
		}
	}
	if len(tr.Args) == 0 {
		switch s.String() { // TODO - fix this hack as early as possible
		case "int":
			return "int32"
		case "long":
			return "int64"
		}
		return s.String()
	}
	s.WriteByte('<')
	for i, a := range tr.Args {
		if i != 0 {
			s.WriteByte(',')
		}
		if a.IsArith {
			s.WriteString(strconv.FormatUint(uint64(a.Arith.Res), 10))
		} else if a.T.Type.String() == "*" {
			s.WriteString("*")
		} else {
			s.WriteString(k.canonicalStringTL1(a.T, false))
		}
	}
	s.WriteByte('>')
	return s.String()
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
				if actualArg.arg.IsArith || actualArg.arg.T.Type.String() == "*" {
					if tr.T.Bare {
						e1 := tr.T.PR.BeautifulError(fmt.Errorf("#-param reference %q cannot be bare", tr.T))
						//e2 := lt.PR.BeautifulError(fmt.Errorf("defined here"))
						//return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
						return tr, nil, e1
					}
					return actualArg.arg, actualArg.natArgs, nil
				}
				tName := k.replaceTL1BuiltinName(actualArg.arg.T.Type.String())
				kt, ok := k.tips[tName]
				if !ok {
					return tr, nil, fmt.Errorf("type %s does not exist", actualArg.arg.T.Type)
				}
				if kt.originTL2 {
					return tr, nil, fmt.Errorf("cannot reference TL2 type %s from TL1", actualArg.arg.T.Type)
				}
				if tr.T.Bare { // overwrite bare
					// TODO - look up type, check if it is union
					if len(kt.combTL1) > 1 {
						// TODO - better error. Does not reference call site
						//----- bare wrapping
						// bareWrapper {X:Type} a:%X = BareWrapper X;
						// bareWrapperTest a:(bareWrapper a.Color) = BareWrapperTest;
						e1 := tr.T.PR.BeautifulError(fmt.Errorf("field type %q is bare, so union %q cannot be passed", tr.T, actualArg.arg.T.Type))
						//e2 := lt.PR.BeautifulError(fmt.Errorf("defined here"))
						//return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
						return tr, nil, e1
					}
					// myUnionA = MyUnion;
					// myUnionB b:int = MyUnion;
					// wrapper {T:Type} a:%T = Wrapper T;
					// useWarpper xx:(wrapper MyUnion) = UseWrapper;
					actualArg.arg.T.Bare = true
					actualArg.arg.T.Type = kt.combTL1[0].Construct.Name // normalize
					// TODO - we must perform canonical conversion of %Int to int here
				}
				return actualArg.arg, actualArg.natArgs, nil
			}
		}
		// probably ref to global type or a typo
	}
	tName := k.replaceTL1BuiltinName(tr.T.Type.String())
	if tName != "__vector" && tName != "__tuple" {
		kt, ok := k.tips[tName]
		if !ok {
			return tr, nil, fmt.Errorf("type %s does not exist", tr.T.Type)
		}
		if kt.originTL2 {
			return tr, nil, fmt.Errorf("cannot reference TL2 type %s from TL1", tr.T.Type)
		}
		if kt.combTL1[0].IsFunction {
			return tr, nil, fmt.Errorf("cannot reference function %s", tr.T.Type)
		}
		if tr.T.Bare || utils.ToLowerFirst(tr.T.Type.Name) == tr.T.Type.Name {
			// TODO - look up type, check if it is union
			if len(kt.combTL1) > 1 {
				// TODO - better error. Does not reference call site
				//----- bare wrapping
				// bareWrapper {X:Type} a:%X = BareWrapper X;
				// bareWrapperTest a:(bareWrapper a.Color) = BareWrapperTest;
				e1 := tr.T.PR.BeautifulError(fmt.Errorf("reference to union %q cannot be bare", tr.T))
				//e2 := lt.PR.BeautifulError(fmt.Errorf("defined here"))
				//return nil, false, nil, HalfResolvedArgument{}, tlast.BeautifulError2(e1, e2)
				return tr, nil, e1
			}
			tr.T.Bare = true
			tr.T.Type = kt.combTL1[0].Construct.Name // normalize
		}
		//else {
		//	if len(kt.combTL1) == 1 {
		//		tr.T.Type = kt.combTL1[0].TypeDecl.Name // normalize
		//	}
		//}
	}
	// TODO - we must perform canonical conversion of %Int to int here
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

func (k *Kernel) resolveMaskTL1(mask tlast.FieldMask, leftArgs []tlast.TemplateArgument,
	actualArgs []LocalArg) (ActualNatArg, error) {
	for i, targ := range leftArgs {
		if targ.FieldName == mask.MaskName {
			actualArg := actualArgs[i]
			if actualArg.wrongTypeErr != nil {
				return ActualNatArg{}, mask.PRName.BeautifulError(fmt.Errorf("reference to field %s impossible, must have # type", targ.FieldName))
			}
			if !targ.IsNat {
				return ActualNatArg{}, mask.PRName.BeautifulError(fmt.Errorf("fieldMask cannot reference Type-parameter %s", targ.FieldName))
			}
			if len(actualArg.natArgs) != 1 {
				return ActualNatArg{}, fmt.Errorf("internal error fieldMask cannot reference Type-parameter %s", targ.FieldName)
			}
			return actualArg.natArgs[0], nil
		}
	}
	return ActualNatArg{}, mask.PRName.BeautifulError(errors.New("fieldMask reference not found"))
}

func (k *Kernel) GetInstanceTL1(tr tlast.TypeRef) (TypeInstance, error) {
	ref, err := k.getInstanceTL1(tr, false)
	if err != nil {
		return nil, err
	}
	return ref.ins, nil
}

func (k *Kernel) getInstanceTL1(tr tlast.TypeRef, create bool) (*TypeInstanceRef, error) {
	canonicalName := k.canonicalStringTL1(tr, true)
	// canonicalName := k.replaceTL1BuiltinName(tr.String())
	if ref, ok := k.instances[canonicalName]; ok {
		return ref, nil
	}
	//if tr.Type.String() == "" {
	//	log.Printf("creating a bracket instance of type %s", canonicalName)
	//	// must store pointer before children GetInstance() terminates recursion
	//	// this instance stays not initialized in case of error, but kernel then is not consistent anyway
	//	ref := k.addInstance(canonicalName, k.brackets)
	//
	//	elemInstance, err := k.GetInstance(tr.BracketType.ArrayType)
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
	//		keyInstance, err := k.GetInstance(tr.BracketType.IndexType.Type)
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
	// must store pointer before children GetInstance() terminates recursion
	// this instance stays mpt initialized in case of error, but kernel then is not consistent anyway
	var err error
	tName := tr.Type.String()
	switch tName {
	case "__vector":
		ref := k.addInstance(canonicalName, k.brackets)
		leftArgs := []tlast.TemplateArgument{{FieldName: "t", IsNat: false}}
		ref.ins, err = k.createVectorTL1(canonicalName, tr, leftArgs, tr.Args)
		if err != nil {
			return nil, err
		}
		return ref, nil
	case "__tuple":
		ref := k.addInstance(canonicalName, k.brackets)
		leftArgs := []tlast.TemplateArgument{{FieldName: "n", IsNat: true}, {FieldName: "t", IsNat: false}}
		ref.ins, err = k.createTupleTL1(canonicalName, tr, leftArgs, tr.Args)
		if err != nil {
			return nil, err
		}
		return ref, nil
	}
	kt, ok := k.tips[tName]
	if !ok {
		return nil, fmt.Errorf("type %s does not exist", tr.Type)
	}
	// must store pointer before children GetInstance() terminates recursion
	// this instance stays not initialized in case of error, but kernel then is not consistent anyway
	ref := k.addInstance(canonicalName, kt)

	if kt.originTL2 {
		return nil, fmt.Errorf("TL1 combinator cannot reference TL2 combinator %s", tr.Type)
	}
	comb := kt.combTL1[0]
	if !comb.IsFunction {
		if len(comb.TemplateArguments) != len(tr.Args) {
			return nil, fmt.Errorf("typeref to %s must have %d template arguments, has %d", canonicalName, len(comb.TemplateArguments), len(tr.Args))
		}
		//switch {
		// TODO - union, etc.
		//default:
		ref.ins, err = k.createOrdinaryTypeTL1FromTL1(canonicalName, kt, tr, kt.combTL1, comb.TemplateArguments, tr.Args)
		//}
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

func (k *Kernel) createOrdinaryTypeTL1FromTL1(canonicalName string, tip *KernelType,
	resolvedType tlast.TypeRef, definition []*tlast.Combinator,
	leftArgs []tlast.TemplateArgument, actualArgs []tlast.ArithmeticOrType) (TypeInstance, error) {

	switch {
	case len(definition) > 1:
		return k.createUnionTL1FromTL1(canonicalName, tip, resolvedType, definition,
			leftArgs, actualArgs)
	case len(definition) == 1:
		return k.createStructTL1FromTL1(canonicalName, tip, resolvedType,
			definition[0].Fields,
			leftArgs, actualArgs,
			false, 0, nil)
	default:
		return nil, fmt.Errorf("wrong type classification, internal error %s", canonicalName)
	}
}
