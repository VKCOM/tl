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
	bare = someType.Bare
	kt, ok := k.tips[tName]
	if !ok {
		return "", false, tr.PR.BeautifulError(fmt.Errorf("type or argument reference %s not found", tName))
	}
	// Special case, Bool is boxed in TL1, but bare in TL2, so we make it boxed always, from the
	// type resolution point of view.
	if someType.Name != kt.tl1BoxedName && kt.canonicalName.String() != "bool" {
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
	//}
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

// template instances in default namespace are moved into argument(s) namespace,
// if only 1 namespace (except default) in arguments
func (k *Kernel) getArgNamespace(rt tlast.TL2TypeRef) string {
	argNamespaces := map[string]struct{}{}
	k.collectArgsNamespaces(tlast.TL2TypeArgument{Type: rt}, argNamespaces)
	if len(argNamespaces) == 1 {
		for ns := range argNamespaces {
			return ns
		}
	}
	if rt.BracketType != nil {
		return ""
	}
	return rt.SomeType.Name.Namespace
}

func (k *Kernel) collectArgsNamespaces(rt tlast.TL2TypeArgument, argNamespaces map[string]struct{}) {
	// This is policy. You can adjust it, so more or less templates instantiations
	// are moved into namespace of arguments. Code should compile anyway.
	if rt.IsNumber || rt.Type.String() == "*" {
		return
	}
	if br := rt.Type.BracketType; br != nil {
		k.collectArgsNamespaces(tlast.TL2TypeArgument{Type: br.ArrayType}, argNamespaces)
		if br.HasIndex {
			k.collectArgsNamespaces(br.IndexType, argNamespaces)
		}
	} else {
		someType := rt.Type.SomeType
		if someType.Name.Namespace != "" {
			argNamespaces[someType.Name.Namespace] = struct{}{}
		}
		for _, arg := range someType.Arguments {
			k.collectArgsNamespaces(arg, argNamespaces)
		}
	}
}
