// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/vkcom/tl/internal/tlast"
)

// Overwrites all files given to kernel.
// For each dir/file.tl containing combinator in a whitelist,
// if dir/file.tl2 does not exist, it is created.
// Then combinator is moved (with conversion) from dir/file.tl into dir/file.tl2
// Original file is left, even if it is empty, because user might wish to move
// remaining comments to the new file.
func (k *Kernel) Migration() error {
	tl2WhiteList := NewWhiteList(k.opts.TL2WhiteList)
	if tl2WhiteList.HasNamespace("") {
		return fmt.Errorf("migration whitelist %q should not contain empty (global) namespace, it must be migrated manually", k.opts.TL2WhiteList)
	}
	migrateTips := map[*KernelType]struct{}{}
	migrateNames := map[tlast.Name]struct{}{}
outer:
	for _, tip := range k.tipsOrdered {
		if tip.originTL2 {
			continue
		}
		migrate := false
		for _, comb := range tip.combTL1 {
			if comb.Builtin {
				// skip warning here
				continue outer
			}
			for _, field := range comb.Fields {
				if field.Excl {
					e1 := field.PR.BeautifulError(fmt.Errorf("!X wrappers cannot be migrated"))
					e1.PrintWarning(k.opts.ErrorWriter, e1)
					continue outer
				}
			}
			if comb.IsFunction {
				for _, t := range comb.TemplateArguments {
					e1 := t.PR.BeautifulError(fmt.Errorf("function declaration %q with template arguments must be migrated manually", comb.Construct.Name.String()))
					e1.PrintWarning(k.opts.ErrorWriter, e1)
					continue outer
				}
			}
			if tl2WhiteList.HasName(comb.Construct.Name) {
				migrate = true
			}
			if !comb.IsFunction && tl2WhiteList.HasName(comb.TypeDecl.Name) {
				migrate = true
			}
		}
		if !migrate {
			continue
		}
		migrateTips[tip] = struct{}{}
		for _, comb := range tip.combTL1 {
			migrateNames[comb.Construct.Name] = struct{}{}
			if !comb.IsFunction {
				migrateNames[comb.TypeDecl.Name] = struct{}{}
				break
			}
		}
	}
	allFiles := map[string]*bytes.Buffer{}
	getBB := func(name string, initial string) *bytes.Buffer {
		if b, ok := allFiles[name]; ok {
			return b
		}
		b := bytes.NewBufferString(initial)
		allFiles[name] = b
		return b
	}
	// Note - tl2 files with comments only will be overwritten completely.
	// there is little values in fixing this.
	for _, typ := range k.filesTL2 {
		_ = getBB(typ.PR.Begin.FileName(), typ.PR.Begin.FileContent())
	}
	for _, typ := range k.filesTL1 {
		bb := getBB(typ.PR.Begin.FileName(), "")
		if _, ok := migrateNames[typ.Construct.Name]; !ok {
			bb.WriteString(typ.AllPR.Begin.FileContent()[typ.AllPR.Begin.Offset():typ.AllPR.End.Offset()])
			continue
		}
	}
	for _, tip := range k.tipsOrdered {
		if _, ok := migrateTips[tip]; !ok {
			continue
		}
		comb := tip.combTL1[0]
		tl2file := comb.AllPR.Begin.FileName() + "2"
		bb := getBB(tl2file, "")
		if bb.Len() == 0 {
			// if there is existing TL2 file with comments only, prevent overwriting it
			was, err := os.ReadFile(tl2file)
			// if we cannot read it, we are OK with it being overwritten
			if err == nil {
				bb.Write(was)
			}
		}
		bb.WriteString(comb.AllPR.Begin.FileContent()[comb.AllPR.Begin.Offset():comb.PR.Begin.Offset()])
		tName := comb.TypeDecl.Name
		cName := comb.Construct.Name
		if len(tip.combTL1) == 1 {
			bb.WriteString("// automatically migrated ") // TODO - remove this
			bb.WriteString(cName.String())
			if comb.IsFunction {
				// migrate function
			} else {
				// migrate struct
			}
			bb.WriteString(comb.PR.End.FileContent()[comb.PR.End.Offset():comb.AllPR.End.Offset()])
		} else {
			// migrate union
			bb.WriteString("// automatically migrated union ") // TODO - remove this
			bb.WriteString(tName.String())
		}
	}
	notTouched := 0
	written := 0
	for name, bb := range allFiles {
		was, err := os.ReadFile(name)
		if err != nil {
			return fmt.Errorf("error reading previous file %q: %w", name, err)
		}
		if string(was) == bb.String() {
			notTouched++
			continue
		}
		written++
		if err := os.WriteFile(name, bb.Bytes(), 0644); err != nil {
			return fmt.Errorf("error writing file %q: %w", name, err)
		}
	}
	log.Printf("migration finished, %d types migrated, %d files written, %d files not touched", len(migrateTips), written, notTouched)
	return nil
}
