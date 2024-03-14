// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"

	tls "github.com/vkcom/tl/internal/tlast/gentlo/tltls"
	"github.com/vkcom/tl/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: tlo2json <path-to-tlo-file>")
		os.Exit(1)
	}
	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error on reading file %s: %v", os.Args[1], err)
	}
	var v4 tls.SchemaV4
	if _, err := v4.ReadBoxed(buf); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error on reading schema: %v", err)
	}
	slices.SortFunc(v4.Constructors, func(a, b tls.Combinator) int {
		valAV4, okA := a.AsV4()
		if !okA {
			panic("invalid union interpretation for tls.combinator_v4: " + valAV4.String())
		}
		valBV4, okB := b.AsV4()
		if !okB {
			panic("invalid union interpretation for tls.combinator_v4: " + valAV4.String())
		}
		return strings.Compare(valAV4.Id, valBV4.Id)
	})
	out, err := v4.WriteJSON(nil)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error on creating json: %v", err)
	}
	jsonFile := strings.Replace(
		os.Args[1],
		".tlo",
		".json",
		1)
	err = os.WriteFile(
		jsonFile,
		utils.JsonPrettyPrint(out),
		0644)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error on writing to file %s: %v", jsonFile, err)
	}
}
