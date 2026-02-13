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
		if strings.Contains(info.Main.Version, "dirty") || strings.Contains(info.Main.Version, "-") {
			// during development, we regenerate goldmaster files, and we do not want them to contain dirty version
			return "(devel)"
		}
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

// finds value for tags in comment with template {tag}:"{value}"
func ExtractTLGenTag(comment string, tag string) (found bool, value string) {
	index := strings.Index(comment, tag)
	if index == -1 {
		return
	}
	comment = comment[index+len(tag):]
	index = strings.Index(comment, ":")
	if index == -1 {
		return
	}
	comment = comment[index+1:]
	index = strings.Index(comment, "\"")
	if index == -1 {
		return
	}
	comment = comment[index+1:]
	index = strings.Index(comment, "\"")
	if index == -1 {
		return
	}
	found = true
	value = comment[:index]
	return
}
