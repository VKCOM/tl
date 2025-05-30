// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/tlcodegen"
)

func TestGen(t *testing.T) {
	outputDir := "output"
	defer func() { require.NoError(t, os.RemoveAll(outputDir)) }()
	data, err := os.ReadFile("./tls/goldmaster.tl")

	require.NoError(t, err)

	ast, err := tlast.ParseTL(string(data))
	if err != nil {
		t.Error(err)
	}

	gen, err := tlcodegen.GenerateCode(ast, tlcodegen.Gen2Options{
		ErrorWriter: io.Discard,
		Verbose:     true,
	})

	require.NoError(t, err)
	require.NoError(t, os.RemoveAll(outputDir))

	err = gen.WriteToDir(outputDir)

	require.NoError(t, err)
}
