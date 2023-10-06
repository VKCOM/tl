// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/exp/slices"
)

type pair struct {
	canonical, common string
}

// walkDeterministic recursively walks through roots directories
// and returns all regular files with extension fileExt
//
// also copied from projects/vktl/internal/walkdeterministic.go
func walkDeterministic(fileExt string, root ...string) ([]string, error) {
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
	slices.SortFunc(pairs, func(a, b pair) int {
		return strings.Compare(a.canonical, b.canonical)
	})
	res := make([]string, 0, len(pairs))
	for _, p := range pairs {
		res = append(res, p.common)
	}
	return res, nil
}

func SHA256All(files []string) (string, error) {
	stream := sha256.New()
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return err.Error(), err
		}
		stream.Write(data)
	}
	sum := stream.Sum(nil)
	return hex.EncodeToString(sum), nil
}

func main() {
	if len(os.Args) < 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: sha256sum dir")
		os.Exit(1)
	}
	files, err := walkDeterministic(".go", os.Args[1:]...)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	sum, err := SHA256All(files)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	fmt.Printf("%s", sum)
	os.Exit(0)
}
