// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
)

// Overwrites all files given to kernel.
// For each dir/file.tl containing combinator in a whitelist,
// if dir/file.tl2 does not exist, it is created.
// Then combinator is moved (with conversion) from dir/file.tl into dir/file.tl2
// Original file is left, even if it is empty, because user might wish to move
// remaining comments to the new file.
func (k *Kernel) Migration() error {
	typesMigrated := 0
	allFiles, err := k.migrationImpl(true, &typesMigrated)
	if err != nil {
		return err
	}
	notTouched := 0
	written := 0
	for name, bb := range allFiles {
		if k.opts.TL2MigrationDevMode {
			if strings.HasSuffix(name, ".tl") {
				name = strings.TrimSuffix(name, ".tl") + ".tl1m"
			} else if strings.HasSuffix(name, ".tl2") {
				name = strings.TrimSuffix(name, ".tl2") + ".tl2m"
			}
		}
		was, err := os.ReadFile(name)
		// especially useful for migration dev mode
		if err == nil && string(was) == bb.String() {
			notTouched++
			continue
		}
		written++
		if err := os.WriteFile(name, bb.Bytes(), 0644); err != nil {
			return fmt.Errorf("error writing file %q: %w", name, err)
		}
	}
	log.Printf("migration finished, %d types migrated, %d files written, %d files not touched", typesMigrated, written, notTouched)
	return nil
}

// migrate and format all files into single deterministic string
func (k *Kernel) MigrationForTests() (string, error) {
	allFiles, err := k.migrationImpl(false, nil)
	if err != nil {
		return "", err
	}
	var names []string
	for name := range allFiles {
		names = append(names, name)
	}
	result := strings.Builder{}
	for _, name := range names {
		bb := allFiles[name]
		_, _ = fmt.Fprintf(&result, "// %s:\n", name)
		result.Write(bb.Bytes())
	}
	return result.String(), nil
}

func (k *Kernel) migrationImpl(tryOpenEmptyTL2Files bool, typesMigrated *int) (map[string]*bytes.Buffer, error) {
	tl2WhiteList := NewWhiteList(k.opts.TL2WhiteList)

	migrateTips := map[*KernelType]struct{}{}
	migrateNames := map[tlast.Name]struct{}{}
outer:
	for _, tip := range k.tipsOrdered {
		if tip.originTL2 || tip.builtin {
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
			}
		}
	}
	if typesMigrated != nil {
		*typesMigrated = len(migrateTips)
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
		bb.WriteString(typ.SectionPR.Begin.FileContent()[typ.SectionPR.Begin.Offset():typ.SectionPR.End.Offset()])
		if _, ok := migrateNames[typ.Construct.Name]; !ok {
			bb.WriteString(typ.AllPR.Begin.FileContent()[typ.AllPR.Begin.Offset():typ.AllPR.End.Offset()])
			continue
		}
	}
	// we decide which #-arguments to keep
	//for _, tip := range k.tipsOrdered {
	//	if _, ok := migrateTips[tip]; !ok {
	//		continue
	//	}
	//}
	for _, tip := range k.tipsOrdered {
		if _, ok := migrateTips[tip]; !ok {
			continue
		}
		comb := tip.combTL1[0]
		tl2file := comb.AllPR.Begin.FileName() + "2"
		bb := getBB(tl2file, "")
		if bb.Len() == 0 {
			// if there is existing TL2 file with comments only, k.filesTL2 is empty, so read here
			was, err := os.ReadFile(tl2file)
			// if we cannot read it, we are OK with attempt to overwrite
			if err == nil {
				bb.Write(was)
			}
		}
		bb.WriteString(comb.AllPR.Begin.FileContent()[comb.AllPR.Begin.Offset():comb.PR.Begin.Offset()])
		if len(tip.combTL1) == 1 {
			cName := comb.Construct.Name
			if comb.IsFunction {
				// migrate function
				bb.WriteString(cName.String())
				if comb.Construct.IDExplicit {
					_, _ = fmt.Fprintf(bb, "#%08x", comb.Construct.ID)
				}
				if len(comb.TemplateArguments) != 0 {
					return nil, comb.Construct.NamePR.BeautifulError(errors.New("internal error: function with template arguments cannot be migrated"))
				}
				fieldsAfterReplace, _, err := k.replaceTL1Brackets(comb)
				if err != nil {
					return nil, err
				}
				if err := k.MigrationFields(bb, comb, fieldsAfterReplace); err != nil {
					return nil, err
				}
				bb.WriteString(" =>\n    ")
				if !k.IsTrueType(comb.FuncDecl) { // otherwise returns nothing
					if err := k.MigrationTypeRef(bb, comb, comb.FuncDecl); err != nil {
						return nil, err
					}
				}
				bb.WriteString(";")
			} else {
				// migrate struct
				bb.WriteString(cName.String())
				if comb.Construct.IDExplicit {
					_, _ = fmt.Fprintf(bb, "#%08x", comb.Construct.ID)
				}
				if err := k.MigrationTemplateArguments(bb, comb); err != nil {
					return nil, err
				}
				fieldsAfterReplace, _, err := k.replaceTL1Brackets(comb)
				if err != nil {
					return nil, err
				}
				if len(fieldsAfterReplace) == 1 && fieldsAfterReplace[0].FieldName == "" &&
					fieldsAfterReplace[0].Mask == nil {
					// migrate alias
					bb.WriteString(" <=> ")
					if err := k.MigrationTypeRef(bb, comb, fieldsAfterReplace[0].FieldType); err != nil {
						return nil, err
					}
					bb.WriteString(";")
				} else {
					// migrate fields
					bb.WriteString(" = ")
					if err := k.MigrationFields(bb, comb, fieldsAfterReplace); err != nil {
						return nil, err
					}
					bb.WriteString(";")
				}
			}
			bb.WriteString(comb.PR.End.FileContent()[comb.PR.End.Offset():comb.AllPR.End.Offset()])
		} else {
			// migrate union
			tName := comb.TypeDecl.Name
			bb.WriteString(tName.String())
			if err := k.MigrationTemplateArguments(bb, comb); err != nil {
				return nil, err
			}
			bb.WriteString(" = ")
			for _, comb := range tip.combTL1 {
				bb.WriteString("\n    | ")
				// TODO - legacy JSON name if differs from inferred
				bb.WriteString(comb.Construct.Name.String())
				bb.WriteString(" ")
				fieldsAfterReplace, _, err := k.replaceTL1Brackets(comb)
				if err != nil {
					return nil, err
				}
				if len(fieldsAfterReplace) == 1 && fieldsAfterReplace[0].FieldName == "" &&
					fieldsAfterReplace[0].Mask == nil {
					// migrate alias
					if err := k.MigrationTypeRef(bb, comb, fieldsAfterReplace[0].FieldType); err != nil {
						return nil, err
					}
				} else {
					// migrate fields
					if err := k.MigrationFields(bb, comb, fieldsAfterReplace); err != nil {
						return nil, err
					}
				}
			}
			bb.WriteString(";")
		}
	}
	return allFiles, nil
}

func (k *Kernel) IsTrueType(rt tlast.TypeRef) bool {
	return rt.Type.String() == "true" || rt.Type.String() == "True"
}

func (k *Kernel) MigrationTemplateArguments(bb *bytes.Buffer, comb *tlast.Combinator) error {
	var targs []tlast.TemplateArgument
	for _, targ := range comb.TemplateArguments {
		if !targ.IsNat || !targ.UsedAsNatVariable {
			targs = append(targs, targ)
		}
	}
	for i, targ := range targs {
		if i == 0 {
			bb.WriteString("<")
		} else {
			bb.WriteString(",")
		}
		bb.WriteString(targ.FieldName)
		if targ.IsNat {
			bb.WriteString(":#")
		} else {
			bb.WriteString(":Type")
		}
		if i == len(comb.TemplateArguments)-1 {
			bb.WriteString(">")
		}
	}
	return nil
}

func (k *Kernel) MigrationFields(bb *bytes.Buffer, comb *tlast.Combinator, fieldsAfterReplace []tlast.Field) error {
	for i, fieldDef := range fieldsAfterReplace {
		if fieldDef.FieldName == "" {
			return fieldDef.PR.BeautifulError(fmt.Errorf("internal error: anonymous field cannot be migrated"))
		}
		if i != 0 {
			bb.WriteString(" ")
		}
		bb.WriteString(fieldDef.FieldName)
		if fieldDef.Mask != nil && !k.IsTrueType(fieldDef.FieldType) { // bit cannot be optional
			bb.WriteString("?")
		}
		bb.WriteString(":")
		if err := k.MigrationTypeRef(bb, comb, fieldDef.FieldType); err != nil {
			return err
		}
	}
	return nil
}

func (k *Kernel) MigrationTypeRef(bb *bytes.Buffer, comb *tlast.Combinator, rt tlast.TypeRef) error {
	_, _ = fmt.Fprintf(bb, "%s", rt.String())
	return nil
}
