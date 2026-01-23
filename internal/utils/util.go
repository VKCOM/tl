// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package utils

import (
	"bytes"
	"encoding/json"
	"runtime/debug"
	"strings"
)

func JsonPrettyPrint(in []byte) []byte {
	var out bytes.Buffer
	err := json.Indent(&out, in, "", "  ")
	if err != nil {
		return in
	}
	return out.Bytes()
}

func AppVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}
	return ""
}

func DoLint(commentRight string) bool {
	if len(commentRight) < 2 {
		return true
	}
	for _, f := range strings.Fields(commentRight[2:]) {
		if f == "tlgen:nolint" {
			return false
		}
	}
	return true
}
