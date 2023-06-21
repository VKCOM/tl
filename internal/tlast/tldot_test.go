// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlast

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDot(t *testing.T) {
	t.Run("", func(t *testing.T) {
		tl, err := ParseTL(`
int#a8509bda ? = Int;
//string#b5286e24 ? = String;
//user#eeb5b7ce 
//    id:int
//    name:string 
//    age:int
//= User;
---functions---
@any hren = Int;
`)
		require.NoError(t, err)
		tl.ToDot()
	})
}
