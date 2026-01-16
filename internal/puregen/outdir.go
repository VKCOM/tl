// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package puregen

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/TwiN/go-color"
	"github.com/vkcom/tl/internal/utils"
)

const buildVersionFormat = `tlgen version: %s
schema url: %s
schema commit: %s
schema version: %d (%v)
`

type OutDir struct {
	Code map[string]string // fileName->Content, split by file names relative to output dir
}

func (gen *OutDir) AddCodeFile(filepathName string, code string) error {
	if _, ok := gen.Code[filepathName]; ok {
		return fmt.Errorf("generator %s: source file %q is generated twice", color.InRed("internal error"), filepathName)
	}
	if gen.Code == nil {
		gen.Code = map[string]string{}
	}
	gen.Code[filepathName] = code
	return nil
}

func (gen *OutDir) formatLint(opts *Options, code string, filename string) string {
	switch {
	case strings.HasSuffix(filename, ".go"):
		formattedCode, err := format.Source([]byte(code))
		if err != nil {
			// We generate code still, because it will be easy to debug when the wrong file is written out
			_, _ = fmt.Fprintf(opts.ErrorWriter, "generator %s: source file %q will not compile due to error: %v", color.InRed("internal error"), filename, err)
		} else {
			code = string(formattedCode)
		}
	case strings.HasSuffix(filename, ".h") ||
		strings.HasSuffix(filename, ".cpp"):
		code = strings.ReplaceAll(code, "\t", "  ")
	}
	return code
}

func (gen *OutDir) Write(opts *Options, markerFile string) error {
	if opts.Outdir == "" {
		return fmt.Errorf("--outdir should not be empty")
	}
	if err := os.Mkdir(opts.Outdir, 0755); err != nil && !os.IsExist(err) { // we thus require parent directory to exist
		return fmt.Errorf("error creating outdir %q: %w", opts.Outdir, err)
	}
	// We do not want to touch files which did not change at all.
	relativeFiles := map[string]bool{}
	var relativeDirs []string
	if err := gen.collectRelativePaths(opts.Outdir, "", relativeFiles, &relativeDirs); err != nil {
		return fmt.Errorf("error reading outdir content %q: %w", opts.Outdir, err)
	}
	if len(relativeFiles) != 0 && !relativeFiles[markerFile] {
		return fmt.Errorf("outdir %q not empty and has no %q marker file, please clean manually", opts.Outdir, markerFile)
	}
	markerContent := fmt.Sprintf(buildVersionFormat,
		strings.TrimSpace(utils.AppVersion()),
		strings.TrimSpace(opts.SchemaURL),
		strings.TrimSpace(opts.SchemaCommit),
		opts.SchemaTimestamp, time.Unix(int64(opts.SchemaTimestamp), 0).UTC())
	if err := gen.AddCodeFile(markerFile, markerContent); err != nil {
		return err
	}
	notTouched := 0
	written := 0
	deleted := 0
	for filepathName, code := range gen.Code {
		code = gen.formatLint(opts, code, filepathName)
		d := filepath.Join(opts.Outdir, filepath.Dir(filepathName))
		f := filepath.Join(opts.Outdir, filepathName)
		if !strings.HasPrefix(filepathName, "..") {
			// we allow relative paths outside gen folder for basictl*
			if err := os.MkdirAll(d, 0755); err != nil && !os.IsExist(err) {
				return fmt.Errorf("error creating dir %q: %w", d, err)
			}
		}
		if relativeFiles[filepathName] {
			delete(relativeFiles, filepathName)
			was, err := os.ReadFile(f)
			if err != nil {
				return fmt.Errorf("error reading previous file %q: %w", f, err)
			}
			if string(was) == code {
				notTouched++
				continue
			}
		}
		written++
		if err := os.WriteFile(f, []byte(code), 0644); err != nil {
			return fmt.Errorf("error writing file %q: %w", f, err)
		}
	}
	for filepathName := range relativeFiles {
		f := filepath.Join(opts.Outdir, filepathName)
		deleted++
		if err := os.Remove(f); err != nil {
			return fmt.Errorf("error deleting previous file %q: %w", f, err)
		}
	}
	for i := len(relativeDirs) - 1; i >= 0; i-- {
		f := filepath.Join(opts.Outdir, relativeDirs[i])
		_ = os.Remove(f) // non-empty dirs simply will not remove. This is good enough for us
	}
	// do not check Verbose
	fmt.Printf("%d target files did not change so were not touched, %d written, %d deleted\n", notTouched, written, deleted)
	return nil
}

func (gen *OutDir) collectRelativePaths(absDirName string, relDirName string, relativeFiles map[string]bool, relativeDirs *[]string) error {
	fis, err := os.ReadDir(absDirName)
	if err != nil {
		return err
	}
	for _, fi := range fis { // try all snapshots, loading the latest
		relFilename := filepath.Join(relDirName, fi.Name())
		absFilename := filepath.Join(absDirName, fi.Name())
		if fi.IsDir() {
			*relativeDirs = append(*relativeDirs, relFilename)
			if err = gen.collectRelativePaths(absFilename, relFilename, relativeFiles, relativeDirs); err != nil {
				return err
			}
			continue
		}
		relativeFiles[relFilename] = true
	}
	return nil
}
