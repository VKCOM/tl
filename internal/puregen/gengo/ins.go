// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"fmt"
	"slices"
	"strings"
	"sync/atomic"
)

// for golang cycle detection
type DirectImports struct {
	ns         map[*InternalNamespace]struct{}
	importSort bool
}

type Namespace struct {
	name  string
	types []*TypeRWWrapper
	decGo Deconflicter
}

type InsFile struct {
	ins      *InternalNamespace
	fileName string
}

type InternalNamespace struct {
	DebugID      int   // for identification in logs
	FloodCounter int64 // beware!

	Namespaces    map[string]struct{}
	DirectImports *DirectImports
	Types         []*TypeRWWrapper

	SubPath string
	Name    string
}

func (n *InternalNamespace) Prefix(directImports *DirectImports, in *InternalNamespace) string {
	if n == in {
		return ""
	}
	directImports.ns[n] = struct{}{}
	return n.Name + "."
}

func (n *InternalNamespace) ImportsSingleNamedNamespace() (empty bool, name string) {
	for nn := range n.Namespaces {
		if nn == "" {
			empty = true
			continue
		}
		if name != "" {
			return empty, ""
		}
		name = nn
	}
	return empty, name
}

var floodCounter atomic.Int64 // TODO - move somewhere

func (n *InternalNamespace) FindRecursiveImports(ri map[*InternalNamespace][]*InternalNamespace, replace *InternalNamespace) {
	for k := range ri {
		delete(ri, k)
	}
	fc := floodCounter.Add(1)
	for k := range n.DirectImports.ns {
		if k != replace {
			k.findRecursiveImports(fc, n, ri, replace, n)
		}
	}
	if replace != nil {
		for k := range replace.DirectImports.ns {
			if k != n {
				k.findRecursiveImports(fc, n, ri, replace, n)
			}
		}
	}
}

func (n *InternalNamespace) findRecursiveImports(floodCounter int64, parent *InternalNamespace, ri map[*InternalNamespace][]*InternalNamespace, replace *InternalNamespace, with *InternalNamespace) {
	if n.FloodCounter >= floodCounter {
		return // visited already
	}
	n.FloodCounter = floodCounter
	ri[n] = append(ri[n], parent)
	for k := range n.DirectImports.ns {
		kk := k.replace(replace, with)
		kk.findRecursiveImports(floodCounter, n, ri, replace, with)
	}
}

func (n InternalNamespace) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("id %d types: ", n.DebugID))
	s.WriteString(strings.Join(n.sortedElements(), ":"))
	s.WriteString("\nnamespaces: ")
	s.WriteString(strings.Join(n.sortedNamespaces(), ","))
	s.WriteString(fmt.Sprintf(" fc=%d\ndirect: ", n.FloodCounter))
	s.WriteString(strings.Join(n.sortedDirectElements(), ","))
	s.WriteString("\nrecursive: ")
	s.WriteString(strings.Join(n.sortedRecursiveElements(), ","))
	s.WriteString("\n")
	return s.String()
}

func (n *InternalNamespace) replace(replace *InternalNamespace, with *InternalNamespace) *InternalNamespace {
	if n == replace {
		return with
	}
	return n
}

// do not forget to remove from []*InternalNamespace after call
func (n *InternalNamespace) mergeFrom(from *InternalNamespace, internalNamespaces []*InternalNamespace) {
	into := n
	for _, t := range from.Types {
		into.Types = append(into.Types, t)
		t.ins = into
	}
	for nn := range from.Namespaces {
		n.Namespaces[nn] = struct{}{}
	}
	for nn := range from.DirectImports.ns {
		if nn != into {
			into.DirectImports.ns[nn] = struct{}{}
		}
	}
	into.DirectImports.importSort = into.DirectImports.importSort || from.DirectImports.importSort
	for _, ins := range internalNamespaces {
		if _, ok := ins.DirectImports.ns[from]; ok {
			if ins != into {
				ins.DirectImports.ns[into] = struct{}{}
			}
			// for k2 := range n.DirectImports {
			//	ins.DirectImports[k2] = struct{}{}
			// }
			delete(ins.DirectImports.ns, from)
		}
	}
}

func (n *InternalNamespace) sortedElements() []string {
	var elements []string
	for _, t := range n.Types {
		elements = append(elements, t.goGlobalName) // TODO - change to tlName.String() and fix problems
	}
	slices.Sort(elements)
	return elements
}

func (n *InternalNamespace) sortedRecursiveElements() []string {
	var elements []string
	ri := map[*InternalNamespace][]*InternalNamespace{}
	n.FindRecursiveImports(ri, nil)
	for r := range ri {
		var inside []string
		for _, t := range r.Types {
			inside = append(inside, t.goGlobalName)
		}
		slices.Sort(inside)
		elements = append(elements, strings.Join(inside, ":"))
	}
	slices.Sort(elements)
	return elements
}

func (n *InternalNamespace) sortedDirectElements() []string {
	var elements []string
	for r := range n.DirectImports.ns {
		var inside []string
		for _, t := range r.Types {
			inside = append(inside, t.goGlobalName)
		}
		slices.Sort(inside)
		elements = append(elements, strings.Join(inside, ":"))
	}
	slices.Sort(elements)
	return elements
}

func (n *InternalNamespace) sortedNamespaces() []string {
	var elements []string
	for n := range n.Namespaces {
		elements = append(elements, n)
	}
	slices.Sort(elements)
	return elements
}
