// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Package checks holds the individual backward-compatibility scenarios for the TL1
// compatibility linter. Each scenario lives in its own file and is exposed as a Check value
// with a common signature, so scenarios can be added, read and reviewed in isolation.
//
// The package is self-contained (depends only on pure and tlast) and knows nothing about the
// command-line wiring; the runner in the parent compat package collects Check values from All
// and executes them.
package checks

import (
	"github.com/VKCOM/tl/internal/pure"
	"github.com/VKCOM/tl/internal/tlast"
)

// Schema is an indexed, comparison-friendly view of one schema version's TL1 combinators.
// It is built once per version and shared by all scenario checks.
type Schema struct {
	// Types groups non-builtin, non-function combinators by their type name (right of '='),
	// so all constructors of a union land together.
	Types      map[tlast.Name][]*tlast.Combinator
	TypesOrder []tlast.Name // declaration order, for deterministic error reporting

	// Functions indexes functions by their constructor name (left of '=').
	Functions      map[tlast.Name]*tlast.Combinator
	FunctionsOrder []tlast.Name

	// Force disables the MatchedCombinators cache: when true, matches are recomputed on every
	// call. Set it if a Schema is mutated after construction (the cache assumes the two schemas
	// being compared stay immutable for the duration of a check run).
	Force bool

	// matched caches MatchedCombinators results, keyed by the cur schema pointer. Every scenario
	// pairs the same two immutable schemas, so this is computed once and reused across checks.
	matched map[*Schema][]CombinatorPair
}

// NewSchema extracts a Schema from a kernel's TL1 combinators. It mirrors the grouping the old
// linter did in tlcodegen.extractTypes.
func NewSchema(k *pure.Kernel) *Schema {
	s := &Schema{
		Types:     make(map[tlast.Name][]*tlast.Combinator),
		Functions: make(map[tlast.Name]*tlast.Combinator),
	}
	for _, combinator := range k.TL1() {
		if combinator.Builtin {
			continue
		}
		if combinator.IsFunction {
			name := combinator.Construct.Name
			if _, ok := s.Functions[name]; !ok {
				s.FunctionsOrder = append(s.FunctionsOrder, name)
			}
			s.Functions[name] = combinator
			continue
		}
		name := combinator.TypeDecl.Name
		if s.Types[name] == nil {
			s.TypesOrder = append(s.TypesOrder, name)
		}
		s.Types[name] = append(s.Types[name], combinator)
	}
	return s
}

// CombinatorPair is a single combinator that exists in both schema versions, matched by name.
type CombinatorPair struct {
	Prev *tlast.Combinator
	Cur  *tlast.Combinator
}

// MatchedCombinators returns every combinator present in both prev (the receiver) and cur,
// matched by name: constructors are matched within their type by constructor name, functions by
// their constructor name. The result follows prev's declaration order, so errors are reported
// deterministically.
//
// Combinators that were removed in cur are intentionally skipped: reporting their disappearance
// is the job of the constructor-removed and function-removed scenarios, not of field-level ones.
//
// The result is cached per cur schema (see Schema.matched); set Schema.Force to bypass the cache.
func (prev *Schema) MatchedCombinators(cur *Schema) []CombinatorPair {
	if !prev.Force {
		if pairs, ok := prev.matched[cur]; ok {
			return pairs
		}
	}
	pairs := prev.matchCombinators(cur)
	if !prev.Force {
		if prev.matched == nil {
			prev.matched = make(map[*Schema][]CombinatorPair)
		}
		prev.matched[cur] = pairs
	}
	return pairs
}

// matchCombinators computes the matched pairs; MatchedCombinators wraps it with caching.
func (prev *Schema) matchCombinators(cur *Schema) []CombinatorPair {
	var pairs []CombinatorPair
	for _, typeName := range prev.TypesOrder {
		curByConstructor := make(map[tlast.Name]*tlast.Combinator, len(cur.Types[typeName]))
		for _, c := range cur.Types[typeName] {
			curByConstructor[c.Construct.Name] = c
		}
		for _, prevConstructor := range prev.Types[typeName] {
			if curConstructor, ok := curByConstructor[prevConstructor.Construct.Name]; ok {
				pairs = append(pairs, CombinatorPair{Prev: prevConstructor, Cur: curConstructor})
			}
		}
	}
	for _, functionName := range prev.FunctionsOrder {
		if curFunction, ok := cur.Functions[functionName]; ok {
			pairs = append(pairs, CombinatorPair{Prev: prev.Functions[functionName], Cur: curFunction})
		}
	}
	return pairs
}
