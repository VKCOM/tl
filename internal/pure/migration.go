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
		bb.WriteString(tip.canonicalName.String())
		if len(tip.combTL1) == 1 {
			if comb.IsFunction {
				// migrate function
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
				leftArgs, err := k.MigrationFields(bb, migrateTips, tip, comb, fieldsAfterReplace)
				if err != nil {
					return nil, err
				}
				bb.WriteString(" =>\n    ")
				if !k.IsTrueType(comb.FuncDecl) { // otherwise returns nothing
					if err := k.MigrationTypeRef(bb, migrateTips, tip, comb, comb.FuncDecl, leftArgs); err != nil {
						return nil, err
					}
				}
				bb.WriteString(";")
			} else {
				// migrate struct
				if comb.Construct.IDExplicit {
					_, _ = fmt.Fprintf(bb, "#%08x", comb.Construct.ID)
				}
				if err := k.MigrationTemplateArguments(bb, tip, comb); err != nil {
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
					if err := k.MigrationTypeRef(bb, migrateTips, tip, comb, fieldsAfterReplace[0].FieldType, comb.TemplateArguments); err != nil {
						return nil, err
					}
					bb.WriteString(";")
				} else {
					// migrate fields
					bb.WriteString(" = ")
					_, err := k.MigrationFields(bb, migrateTips, tip, comb, fieldsAfterReplace)
					if err != nil {
						return nil, err
					}
					bb.WriteString(";")
				}
			}
			bb.WriteString(comb.PR.End.FileContent()[comb.PR.End.Offset():comb.AllPR.End.Offset()])
		} else {
			// migrate union
			if err := k.MigrationTemplateArguments(bb, tip, comb); err != nil {
				return nil, err
			}
			bb.WriteString(" = ")
			variantNames, err := k.VariantNames(tip.combTL1)
			if err != nil {
				return nil, err
			}
			for i, comb := range tip.combTL1 {
				_, _ = fmt.Fprintf(bb, "\n    // tlgen:tl1name:%s", comb.Construct.Name.String())
				bb.WriteString("\n    | ")
				bb.WriteString(variantNames[i].String())
				bb.WriteString(" ")
				fieldsAfterReplace, _, err := k.replaceTL1Brackets(comb)
				if err != nil {
					return nil, err
				}
				if len(fieldsAfterReplace) == 1 && fieldsAfterReplace[0].FieldName == "" &&
					fieldsAfterReplace[0].Mask == nil {
					// migrate alias
					if err := k.MigrationTypeRef(bb, migrateTips, tip, comb, fieldsAfterReplace[0].FieldType, comb.TemplateArguments); err != nil {
						return nil, err
					}
				} else {
					// migrate fields
					_, err := k.MigrationFields(bb, migrateTips, tip, comb, fieldsAfterReplace)
					if err != nil {
						return nil, err
					}
				}
			}
			bb.WriteString(";")
		}
	}
	return allFiles, nil
}

func (k *Kernel) MigrationTemplateArguments(bb *bytes.Buffer, tip *KernelType, comb *tlast.Combinator) error {
	var targs []tlast.TemplateArgument
	for i, targ := range comb.TemplateArguments {
		if !tip.targs[i].usedAsNatVariable {
			targs = append(targs, targ)
		}
	}
	for i, targ := range targs {
		if i == 0 {
			bb.WriteString("<")
		} else {
			bb.WriteString(", ")
		}
		bb.WriteString(targ.FieldName)
		if targ.IsNat {
			bb.WriteString(":#")
		} else {
			bb.WriteString(":Type")
		}
		if i == len(targs)-1 {
			bb.WriteString(">")
		}
	}
	return nil
}

func (k *Kernel) MigrationFields(bb *bytes.Buffer, migrateTips map[*KernelType]struct{}, tip *KernelType, comb *tlast.Combinator, fieldsAfterReplace []tlast.Field) ([]tlast.TemplateArgument, error) {
	leftArgs := comb.TemplateArguments
	for i, fieldDef := range fieldsAfterReplace {
		if fieldDef.FieldName == "" {
			return nil, fieldDef.PR.BeautifulError(fmt.Errorf("internal error: anonymous field cannot be migrated"))
		}
		if i != 0 {
			bb.WriteString(" ")
		}
		bb.WriteString(fieldDef.FieldName)
		if fieldDef.Mask != nil && !k.IsTrueType(fieldDef.FieldType) { // bit cannot be optional
			bb.WriteString("?")
		}
		bb.WriteString(":")
		if err := k.MigrationTypeRef(bb, migrateTips, tip, comb, fieldDef.FieldType, leftArgs); err != nil {
			return nil, err
		}
		if fieldDef.FieldName != "" {
			leftArgs = append(leftArgs, tlast.TemplateArgument{
				FieldName: fieldDef.FieldName,
				IsNat:     true,
				PR:        fieldDef.PR,
			})
		}
	}
	return leftArgs, nil
}

func (k *Kernel) MigrationTypeRef(bb *bytes.Buffer, migrateTips map[*KernelType]struct{}, tip *KernelType, comb *tlast.Combinator,
	tr tlast.TypeRef, leftArgs []tlast.TemplateArgument) error {
	result, err := k.MigrationTypeRefImpl(migrateTips, tip, tr, leftArgs)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(bb, "%s", result.String())
	return nil
}

func (k *Kernel) MigrationTypeRefImpl(migrateTips map[*KernelType]struct{}, tip *KernelType,
	tr tlast.TypeRef, leftArgs []tlast.TemplateArgument) (tlast.TL2TypeRef, error) {
	result, _, err := k.MigrationArgument(migrateTips, tip, tlast.ArithmeticOrType{T: tr}, leftArgs, false)
	if err != nil {
		return tlast.TL2TypeRef{}, err
	}
	if result.IsNumber {
		return tlast.TL2TypeRef{}, fmt.Errorf("internal error during migration: number where type reference is required")
	}
	return result.Type, nil
}

func (k *Kernel) MigrationArgument(migrateTips map[*KernelType]struct{}, tip *KernelType,
	tra tlast.ArithmeticOrType, leftArgs []tlast.TemplateArgument, allowRemoved bool) (tlast.TL2TypeArgument, bool, error) {

	if tra.IsArith {
		return tlast.TL2TypeArgument{IsNumber: true, Number: tra.Arith.Res}, false, nil
	}
	tr := tra.T

	if tr.Type.Namespace == "" {
		for i, targ := range leftArgs {
			if targ.FieldName == tr.Type.Name {
				for _, arg := range tr.Args {
					e1 := arg.T.PR.BeautifulError(fmt.Errorf("reference to template argument %s cannot have arguments", targ.FieldName))
					e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
					return tlast.TL2TypeArgument{}, false, tlast.BeautifulError2(e1, e2)
				}
				result := tlast.TL2TypeArgument{Type: tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: tlast.TL2TypeName(tr.Type)}}}
				removed := i >= len(tip.targs) || tip.targs[i].usedAsNatVariable
				if removed {
					// reference to field or removed argument
					if !allowRemoved {
						e1 := tr.PR.BeautifulError(fmt.Errorf("reference to template argument  %s being removed during migration", targ.FieldName))
						e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
						return tlast.TL2TypeArgument{}, false, tlast.BeautifulError2(e1, e2)
					}
				}
				return result, removed, nil
			}
		}
	}
	tName := tr.Type.String()
	// TODO - dictionaries
	switch tName {
	case "__vector", "Vector", "vector":
		if len(tr.Args) != 1 || tr.Args[0].IsArith {
			return tlast.TL2TypeArgument{}, false, tr.PR.BeautifulError(errors.New("expected single type argument here"))
		}
		elemType, err := k.MigrationTypeRefImpl(migrateTips, tip, tr.Args[0].T, leftArgs)
		if err != nil {
			return tlast.TL2TypeArgument{}, false, err
		}
		return tlast.TL2TypeArgument{Type: tlast.TL2TypeRef{BracketType: &tlast.TL2BracketType{ArrayType: elemType}}}, false, nil
	case "__tuple", "Tuple", "tuple":
		if len(tr.Args) != 2 {
			return tlast.TL2TypeArgument{}, false, tr.PR.BeautifulError(errors.New("expected 2 arguments here"))
		}
		argType := tr.Args[0]
		argCount := tr.Args[1]
		if tName == "__tuple" {
			argCount, argType = argType, argCount
		}
		elemType, err := k.MigrationTypeRefImpl(migrateTips, tip, argType.T, leftArgs)
		if err != nil {
			return tlast.TL2TypeArgument{}, false, err
		}
		bracketType := tlast.TL2BracketType{ArrayType: elemType}
		indexType, removed, err := k.MigrationArgument(migrateTips, tip, argCount, leftArgs, true)
		if err != nil {
			return tlast.TL2TypeArgument{}, false, err
		}
		if !removed {
			bracketType.IndexType = indexType
			bracketType.HasIndex = true
		}
		return tlast.TL2TypeArgument{Type: tlast.TL2TypeRef{BracketType: &bracketType}}, false, nil
	}
	kt, ok := k.tips[tName]
	if !ok {
		return tlast.TL2TypeArgument{}, false, fmt.Errorf("type %s does not exist", tr.Type)
	}
	isDict, dictFieldT := k.IsDict(kt)
	if isDict {
		if len(tr.Args) != len(dictFieldT.combTL1[0].TemplateArguments) {
			return tlast.TL2TypeArgument{}, false, tr.PR.BeautifulError(fmt.Errorf("internal error during migration: expected %d arguments here", len(dictFieldT.combTL1[0].TemplateArguments)))
		}
		for _, targ := range tr.Args {
			if targ.IsArith {
				return tlast.TL2TypeArgument{}, false, targ.T.PR.BeautifulError(errors.New("internal error during migration: dictionary cannot be instantiated with number"))
			}
		}
		keyRT := dictFieldT.combTL1[0].Fields[0].FieldType
		valueRT := tr.Args[0].T
		if len(tr.Args) == 2 {
			keyRT = tr.Args[0].T
			valueRT = tr.Args[1].T
		}
		valueType, err := k.MigrationTypeRefImpl(migrateTips, tip, keyRT, leftArgs)
		if err != nil {
			return tlast.TL2TypeArgument{}, false, err
		}
		keyType, err := k.MigrationTypeRefImpl(migrateTips, tip, valueRT, leftArgs)
		if err != nil {
			return tlast.TL2TypeArgument{}, false, err
		}
		bracketType := tlast.TL2BracketType{
			ArrayType: valueType,
			IndexType: tlast.TL2TypeArgument{Type: keyType},
			HasIndex:  true,
		}
		return tlast.TL2TypeArgument{Type: tlast.TL2TypeRef{BracketType: &bracketType}}, false, nil
	}
	if kt.originTL2 {
		return tlast.TL2TypeArgument{}, false, fmt.Errorf("during migration, reference to TL2 type %s is found", tr.Type)
	}
	_, migrateTip := migrateTips[kt]
	//result := tlast.TL2TypeApplication{Name: tlast.TL2TypeName(tr.Type)}
	//if migrateTip {
	result := tlast.TL2TypeApplication{Name: tlast.TL2TypeName(kt.canonicalName)}
	//}
	if len(tr.Args) != len(kt.targs) {
		return tlast.TL2TypeArgument{}, false, fmt.Errorf("internal error during migration, reference to type %s has wrong # of arguments found", tr.Type)
	}
	for i, arg := range tr.Args {
		if migrateTip && kt.targs[i].usedAsNatVariable { // target type removes this arg
			continue
		}
		if arg.IsArith {
			result.Arguments = append(result.Arguments, tlast.TL2TypeArgument{IsNumber: true, Number: arg.Arith.Res})
			continue
		}
		indexType, _, err := k.MigrationArgument(migrateTips, tip, arg, leftArgs, false)
		if err != nil {
			return tlast.TL2TypeArgument{}, false, err
		}
		result.Arguments = append(result.Arguments, indexType)
	}
	return tlast.TL2TypeArgument{Type: tlast.TL2TypeRef{SomeType: result}}, false, nil
}
