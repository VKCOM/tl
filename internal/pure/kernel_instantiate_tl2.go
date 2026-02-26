// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
)

// TODO - make this simple formatiing, do not look up here, look up only in resolveType
// top level types do not have bare/boxed in their names, instead bare is returned from function
func (k *Kernel) canonicalString(tr tlast.TL2TypeRef, top bool) (_ string, bare bool, _ error) {
	var s strings.Builder

	if br := tr.BracketType; br != nil {
		s.WriteString("[")
		if br.HasIndex {
			if br.IndexType.IsNumber {
				s.WriteString(strconv.FormatUint(uint64(br.IndexType.Number), 10))
			} else if br.IndexType.Type.String() == "*" {
				s.WriteString("*")
			} else {
				str, _, err := k.canonicalString(br.IndexType.Type, false)
				if err != nil {
					return "", false, err
				}
				s.WriteString(str)
			}
		}
		s.WriteString("]")
		str, _, err := k.canonicalString(br.ArrayType, false)
		if err != nil {
			return "", false, err
		}
		s.WriteString(str)
		return s.String(), true, nil
	}
	someType := tr.SomeType
	tName := someType.Name.String()
	kt, ok := k.tips[tName]
	if !ok {
		return "", false, tr.PR.BeautifulError(fmt.Errorf("type or argument reference %s not found", tName))
	}
	bare = someType.Bare
	if someType.Name != kt.tl1BoxedName {
		bare = true
	}
	if bare && !kt.CanBeBare() {
		// TODO - we could simply treat % as "bare if possible", which would allow writing it basically everywhere
		e1 := tr.PR.BeautifulError(fmt.Errorf("type reference to %s cannot be bare", tName))
		if kt.originTL2 {
			// TODO - beautiful
			return "", false, e1
		}
		e2 := kt.combTL1[0].TypeDecl.NamePR.BeautifulError(errSeeHere)
		return "", false, tlast.BeautifulError2(e1, e2)
	}
	if !bare && !kt.CanBeBoxed() { // TODO - impossible?
		e1 := tr.PR.BeautifulError(fmt.Errorf("type reference to %s cannot be boxed", tName))
		if kt.originTL2 {
			// TODO - beautiful
			return "", false, e1
		}
		e2 := kt.combTL1[0].Construct.NamePR.BeautifulError(errSeeHere)
		return "", false, tlast.BeautifulError2(e1, e2)
	}
	if !top && !bare && kt.CanBeBare() {
		s.WriteString("+")
	}
	s.WriteString(kt.canonicalName.String())
	if len(someType.Arguments) == 0 {
		return s.String(), bare, nil
	}
	s.WriteByte('<')
	for i, a := range someType.Arguments {
		if i != 0 {
			s.WriteByte(',')
		}
		if a.IsNumber {
			s.WriteString(strconv.FormatUint(uint64(a.Number), 10))
		} else if a.Type.String() == "*" {
			s.WriteString("*")
		} else {
			str, _, err := k.canonicalString(a.Type, false)
			if err != nil {
				return "", false, err
			}
			s.WriteString(str)
		}
	}
	s.WriteByte('>')
	return s.String(), bare, nil
}

// alias || fields || union
func (k *Kernel) createOrdinaryType(canonicalName string, tip *KernelType, tr tlast.TL2TypeRef,
	tlName tlast.TL2TypeName, tlTag uint32,
	definition tlast.TL2TypeDefinition,
	leftArgs []tlast.TL2TypeTemplate, actualArgs []tlast.TL2TypeArgument) (TypeInstance, error) {

	switch {
	case definition.IsAlias():
		return k.createAlias(canonicalName, tip, tr, definition.TypeAlias, leftArgs, actualArgs)
	case definition.StructType.IsUnionType:
		return k.createUnion(canonicalName, tip, tr, tlTag, definition.StructType.UnionType, leftArgs)
	default:
		return k.createStruct(canonicalName, tip, tr,
			tlName, tlTag,
			true, definition.TypeAlias, definition.StructType.ConstructorFields,
			leftArgs, false, 0, nil, false)
	}
}
