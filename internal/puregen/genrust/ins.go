// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package genrust

import (
	"github.com/VKCOM/tl/internal/puregen"
)

// for tracking imports.
type DirectImports struct {
}

type Namespace struct {
	name  string
	types []*TypeRWWrapper
	decGo puregen.Deconflicter
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
