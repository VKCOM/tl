// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package puregen

import (
	"strconv"
)

// During recursive generation, we store wrappers to type when they are needed, so that
// we can generate actual types later, when all references to wrappers are set
// also wrapper stores common information

type Deconflicter struct {
	usedNames map[string]bool
}

//func (d *Deconflicter) hasConflict(s string) bool {
//	_, ok := d.usedNames[s]
//	return ok
//}

func (d *Deconflicter) DeconflictName(s string) string {
	if d.usedNames == nil {
		d.usedNames = map[string]bool{}
	}
	var suffix string
	for i := 0; d.usedNames[s+suffix]; i++ {
		suffix = strconv.Itoa(i)
	}
	s += suffix
	d.usedNames[s] = true
	return s
}

func (d *Deconflicter) FillGolangIdentifies() {
	d.DeconflictName("Write")
	d.DeconflictName("Read")
	d.DeconflictName("WriteTL2")
	d.DeconflictName("ReadTL2")
}
