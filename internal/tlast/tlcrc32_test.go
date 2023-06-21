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

func TestCRC32TL(t *testing.T) {
	tests := []struct {
		tlText string
		tlTag  uint32
	}{
		{`---functions--- @any  get_arrays n:# a:n*[int] b:5*[int] = Tuple int 5;`, 0x90658cdb},
		{`---functions--- @any  get_arrays#12345678 n:# a:n*[int] b:5*[int] = Tuple int 5;`, 0x12345678},
	}
	for _, test := range tests {
		tl, err := ParseTL(test.tlText)
		require.NoError(t, err)
		require.Equal(t, test.tlTag, tl[0].Crc32())
	}
}
