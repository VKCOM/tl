// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package utils

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type pair struct {
	canonical, common string
}

// WalkDeterministic recursively walks through roots directories
// and returns all regular files with extension fileExt
//
// also copied to projects/vktl/cmd/sha256sum/sha256sum.go
func WalkDeterministic(fileExt string, root ...string) ([]string, error) {
	var pairs []pair
	for _, r := range root {
		err := filepath.Walk(r, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			if !strings.HasSuffix(path, fileExt) {
				return nil
			}
			pairs = append(pairs, pair{
				canonical: strings.ReplaceAll(path, string(os.PathSeparator), "/"),
				common:    path})
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	sort.SliceStable(pairs, func(i, j int) bool { return pairs[i].canonical < pairs[j].canonical })
	res := make([]string, 0, len(pairs))
	for _, p := range pairs {
		res = append(res, p.common)
	}
	return res, nil
}
