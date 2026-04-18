// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genphp

import (
	"fmt"
)

// check that brackets cannot be function return type

type TypeRWBrackets struct {
	wr          *TypeRWWrapper
	vectorLike  bool   // # [T], because # has no reference name
	dynamicSize bool   // with passed nat param
	size        uint32 // if !dynamicSize
	element     Field

	dictLike       bool // for now, can be true only if vectorLike is true. But should work for dynamicSize tuples, so TODO
	dictKeyString  bool
	dictKeyField   Field
	dictValueField Field
}

func phpDictionaryElement(wr *TypeRWWrapper) Field {
	if !phpIsDictionary(wr) {
		panic(fmt.Sprintf("not a dict: %s", wr.TLName()))
	}
	structElement, _ := wr.trw.(*TypeRWStruct)
	return structElement.Fields[1]
}

func phpIsDictionary(wr *TypeRWWrapper) bool {
	return PHPIsDict(wr.pureType.KernelType())
}
