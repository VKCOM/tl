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
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/vkcom/tl/internal/tlast"
)

type migrationTL1RefsTL2Errors struct {
	totalErrors        int
	errNamespaces      map[string]struct{} // error per namespace greatly helps during migration
	migrationErrorList []*tlast.ParseError
}

func (m *migrationTL1RefsTL2Errors) addError(comb *tlast.Combinator, e1 *tlast.ParseError) {
	m.totalErrors++
	ns := comb.Construct.Name.Namespace // ignore types with 2 namespaces
	//if _, ok := m.errNamespaces[ns]; ok {
	//	return
	//}
	m.errNamespaces[ns] = struct{}{}
	//too much clutter, decided simple error is better
	//e2 := comb.Construct.NamePR.BeautifulError(errSeeHere)
	//m.migrationErrorList = append(m.migrationErrorList, tlast.BeautifulError2(e1, e2))
	m.migrationErrorList = append(m.migrationErrorList, e1)
}

// Overwrites all files given to kernel.
// For each dir/file.tl containing combinator in a whitelist,
// if dir/file.tl2 does not exist, it is created.
// Then combinator is moved (with conversion) from dir/file.tl into dir/file.tl2
// Original file is left, even if it is empty, because user might wish to move
// remaining comments to the new file.
func (k *Kernel) Migration() error {
	if err := k.Compile(); err != nil {
		return err
	}
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
				name = strings.TrimSuffix(name, ".tl") + "_migr.tl"
			} else if strings.HasSuffix(name, ".tl2") {
				name = strings.TrimSuffix(name, ".tl2") + "_migr.tl2"
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
	fmt.Printf("migration finished, %d types migrated, %d files written, %d files not touched\n", typesMigrated, written, notTouched)
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
	migrateTips := map[*KernelType]struct{}{}
	migrateNames := map[tlast.Name]struct{}{}
outer:
	for _, tip := range k.tipsOrdered {
		if tip.originTL2 || tip.builtin {
			continue
		}
		migrate := false
		for _, comb := range tip.combTL1 {
			if k.tl2WhiteList.HasName(comb.Construct.Name) {
				migrate = true
			}
			if !comb.IsFunction && k.tl2WhiteList.HasName(comb.TypeDecl.Name) {
				migrate = true
			}
		}
		if !migrate {
			continue
		}
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
	// check there will be no references to TL2 combinators from TL1 combinators
	refErrList, err := k.findTL1toTL2References(migrateTips)
	if err != nil {
		return nil, err
	}
	for _, err := range refErrList.migrationErrorList {
		err.PrintWarning(k.opts.ErrorWriter, nil)
	}
	if len(refErrList.migrationErrorList) != 0 { // do not need beautiful error here
		nss := slices.Sorted(maps.Keys(refErrList.errNamespaces))

		return nil, fmt.Errorf("migration failed with %d would be TL1->TL2 references in %d namespaces %s",
			refErrList.totalErrors, len(refErrList.errNamespaces), strings.Join(nss, ","))
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
				bb.WriteString(" ")
				fieldsAfterReplace, typesAfterReplace, _, err := k.replaceTL1Brackets(comb)
				if err != nil {
					return nil, err
				}
				leftArgs, err := k.MigrationFields(bb, migrateTips, tip, comb, fieldsAfterReplace, typesAfterReplace, false)
				if err != nil {
					return nil, err
				}
				bb.WriteString("\n    => ")
				if !k.IsTrueType(comb.FuncDecl) { // otherwise returns nothing
					fieldType := k.convertTypeRef(comb.FuncDecl)
					if err := k.MigrationTypeRef(bb, migrateTips, tip, comb, fieldType, leftArgs); err != nil {
						return nil, err
					}
				}
				bb.WriteString(";")
			} else {
				// migrate struct. we decided that migrating tags should be done manually, if ever needed
				// if comb.Construct.IDExplicit {
				//	_, _ = fmt.Fprintf(bb, "#%08x", comb.Construct.ID)
				// }
				if err := k.MigrationTemplateArguments(bb, tip, comb); err != nil {
					return nil, err
				}
				fieldsAfterReplace, typesAfterReplace, _, err := k.replaceTL1Brackets(comb)
				if err != nil {
					return nil, err
				}
				if len(fieldsAfterReplace) == 1 && fieldsAfterReplace[0].FieldName == "" &&
					fieldsAfterReplace[0].Mask == nil {
					// migrate alias
					bb.WriteString(" <=> ")
					if err := k.MigrationTypeRef(bb, migrateTips, tip, comb, typesAfterReplace[0], comb.TemplateArguments); err != nil {
						return nil, err
					}
					bb.WriteString(";")
				} else {
					// migrate fields
					bb.WriteString(" = ")
					_, err := k.MigrationFields(bb, migrateTips, tip, comb, fieldsAfterReplace, typesAfterReplace, false)
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
				// Ensure the same generation for both old JSON format and
				// the same golang type names to reduce diff and changes to projects.
				// After some time, remove this logic and simply use names which fit.
				// TODO - it appears to be too complex task for now.
				// We will return to this code later, for now all unions will have their
				// tl1name stored.
				//wouldBeName := tip.CanonicalName()
				//writeExplicitTL1Name := variantNames[i].Namespace != ""
				//if !writeExplicitTL1Name {
				//	wouldBeName.Name += variantNames[i].Name
				//	wouldBeName.Name = utils.ToLowerFirst(wouldBeName.Name)
				//	writeExplicitTL1Name = wouldBeName != comb.Construct.Name
				//	if writeExplicitTL1Name {
				//		wouldBeName = tip.CanonicalName()
				//		wouldBeName.Name = variantNames[i].Name + wouldBeName.Name
				//		writeExplicitTL1Name = wouldBeName != comb.Construct.Name
				//	}
				//}
				//if writeExplicitTL1Name {
				_, _ = fmt.Fprintf(bb, "\n    // tlgen:tl1name:%q", comb.Construct.Name.String())
				//}
				bb.WriteString("\n    | ")
				bb.WriteString(variantNames[i])
				bb.WriteString(" ")
				fieldsAfterReplace, typesAfterReplace, _, err := k.replaceTL1Brackets(comb)
				if err != nil {
					return nil, err
				}
				if len(fieldsAfterReplace) == 1 && fieldsAfterReplace[0].FieldName == "" &&
					fieldsAfterReplace[0].Mask == nil {
					// migrate alias
					if err := k.MigrationTypeRef(bb, migrateTips, tip, comb, typesAfterReplace[0], comb.TemplateArguments); err != nil {
						return nil, err
					}
				} else {
					// migrate fields
					_, err := k.MigrationFields(bb, migrateTips, tip, comb, fieldsAfterReplace, typesAfterReplace, true)
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

func (k *Kernel) findTL1toTL2References(migrateTips map[*KernelType]struct{}) (migrationTL1RefsTL2Errors, error) {
	// check there will be no references to TL2 combinators from TL1 combinators
	refErrList := migrationTL1RefsTL2Errors{errNamespaces: map[string]struct{}{}}
	for _, tip := range k.tipsOrdered {
		if _, ok := migrateTips[tip]; ok {
			continue
		}
		if tip.originTL2 {
			continue
		}
		// TL1 type and it is not migrating
		for _, typ := range tip.combTL1 {
			leftArgs := typ.TemplateArguments
			for _, fieldDef := range typ.Fields {
				if err := k.MigrationCheckTL2FromTL1Field(&refErrList, fieldDef, migrateTips, typ, leftArgs); err != nil {
					return refErrList, err
				}
				if fieldDef.FieldName != "" {
					leftArgs = append(leftArgs, tlast.TemplateArgument{
						FieldName: fieldDef.FieldName,
						IsNat:     true,
						PR:        fieldDef.PR,
					})
				}
			}
			if typ.IsFunction {
				if err := k.MigrationCheckTL2FromTL1Type(&refErrList, typ.FuncDecl, migrateTips, typ, leftArgs); err != nil {
					return refErrList, err
				}
			}
		}
	}
	return refErrList, nil
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

func (k *Kernel) MigrationFields(bb *bytes.Buffer, migrateTips map[*KernelType]struct{}, tip *KernelType, comb *tlast.Combinator,
	fieldsAfterReplace []tlast.Field, typesAfterReplace []tlast.TL2TypeRef, indent bool) ([]tlast.TemplateArgument, error) {
	leftArgs := comb.TemplateArguments
	for i, fieldDef := range fieldsAfterReplace {
		fieldType := typesAfterReplace[i]
		if fieldDef.FieldName == "" {
			return nil, fieldDef.PR.BeautifulError(fmt.Errorf("internal error: anonymous field cannot be migrated"))
		}
		if fieldDef.CommentBefore != "" {
			bb.WriteString(fieldDef.CommentBefore)
			bb.WriteString("\n    ")
			if indent {
				bb.WriteString("    ")
			}
		}
		//if i != 0 {
		//	bb.WriteString(" ")
		//}
		bb.WriteString(fieldDef.FieldName)
		if fieldDef.Mask != nil && k.IsTrueType2(fieldType) {
			bb.WriteString(":bit")
		} else {
			if fieldDef.Mask != nil {
				bb.WriteString("?")
			}
			bb.WriteString(":")
			if err := k.MigrationTypeRef(bb, migrateTips, tip, comb, fieldType, leftArgs); err != nil {
				return nil, err
			}
		}
		if fieldDef.CommentRight != "" {
			bb.WriteString(" ")
			bb.WriteString(fieldDef.CommentRight)
		}
		if fieldDef.NewlineRight {
			bb.WriteString("\n    ")
			if indent {
				bb.WriteString("    ")
			}
		} else {
			bb.WriteString(" ")
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
	tr tlast.TL2TypeRef, leftArgs []tlast.TemplateArgument) error {
	result, err := k.MigrationTypeRefImpl(migrateTips, tip, tr, leftArgs)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(bb, "%s", result.String())
	return nil
}

func (k *Kernel) MigrationTypeRefImpl(migrateTips map[*KernelType]struct{}, tip *KernelType,
	tr tlast.TL2TypeRef, leftArgs []tlast.TemplateArgument) (tlast.TL2TypeRef, error) {
	result, _, err := k.MigrationArgument(migrateTips, tip, tlast.TL2TypeArgument{Type: tr}, leftArgs, false)
	if err != nil {
		return tlast.TL2TypeRef{}, err
	}
	if result.IsNumber {
		return tlast.TL2TypeRef{}, fmt.Errorf("internal error during migration: number where type reference is required")
	}
	return result.Type, nil
}

func (k *Kernel) MigrationArgument(migrateTips map[*KernelType]struct{}, tip *KernelType,
	tra tlast.TL2TypeArgument, leftArgs []tlast.TemplateArgument, allowRemoved bool) (tlast.TL2TypeArgument, bool, error) {

	if tra.IsNumber {
		return tra, false, nil
	}
	if tra.Type.BracketType != nil {
		br := *tra.Type.BracketType
		tra.Type.BracketType = &br
		if br.HasIndex {
			result, removed, err := k.MigrationArgument(migrateTips, tip, br.IndexType, leftArgs, true)
			if err != nil {
				return tlast.TL2TypeArgument{}, false, err
			}
			if removed {
				br.HasIndex = false
				br.IndexType = tlast.TL2TypeArgument{}
			} else {
				br.IndexType = result
			}
		}
		result, err := k.MigrationTypeRefImpl(migrateTips, tip, br.ArrayType, leftArgs)
		if err != nil {
			return tlast.TL2TypeArgument{}, false, err
		}
		br.ArrayType = result
		return tra, false, nil
	}
	someType := tra.Type.SomeType

	if someType.Name.Namespace == "" {
		for i, targ := range leftArgs {
			if targ.FieldName == someType.Name.Name {
				for _, arg := range someType.Arguments {
					e1 := arg.PR.BeautifulError(fmt.Errorf("reference to template argument %s cannot have arguments", targ.FieldName))
					e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
					return tlast.TL2TypeArgument{}, false, tlast.BeautifulError2(e1, e2)
				}
				//result := tlast.TL2TypeArgument{Type: tlast.TL2TypeRef{SomeType: tlast.TL2TypeApplication{Name: tlast.TL2TypeName(tr.Type)}}}
				removed := i >= len(tip.targs) || tip.targs[i].usedAsNatVariable
				if removed {
					// reference to field or removed argument
					if !allowRemoved {
						e1 := someType.PR.BeautifulError(fmt.Errorf("reference to template argument  %s being removed during migration", targ.FieldName))
						e2 := targ.PR.BeautifulError(fmt.Errorf("declared here"))
						return tlast.TL2TypeArgument{}, false, tlast.BeautifulError2(e1, e2)
					}
				}
				return tra, removed, nil
			}
		}
	}
	tName := someType.Name.String()
	switch tName {
	case "Vector", "vector":
		if len(someType.Arguments) != 1 || someType.Arguments[0].IsNumber {
			return tlast.TL2TypeArgument{}, false, someType.PR.BeautifulError(errors.New("expected single type argument here"))
		}
		elemType, err := k.MigrationTypeRefImpl(migrateTips, tip, someType.Arguments[0].Type, leftArgs)
		if err != nil {
			return tlast.TL2TypeArgument{}, false, err
		}
		return tlast.TL2TypeArgument{Type: tlast.TL2TypeRef{BracketType: &tlast.TL2BracketType{ArrayType: elemType}}}, false, nil
	case "Tuple", "tuple":
		if len(someType.Arguments) != 2 || someType.Arguments[0].IsNumber {
			return tlast.TL2TypeArgument{}, false, someType.PR.BeautifulError(errors.New("expected type and nat arguments here"))
		}
		argType := someType.Arguments[0]
		argCount := someType.Arguments[1]
		//if tName == "__tuple" {
		//	argCount, argType = argType, argCount
		//}
		elemType, err := k.MigrationTypeRefImpl(migrateTips, tip, argType.Type, leftArgs)
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
		return tlast.TL2TypeArgument{}, false, fmt.Errorf("type %s does not exist", tName)
	}
	if kt.builtinWrappedCanonicalName != "" {
		tName = kt.builtinWrappedCanonicalName
		kt, ok = k.tips[tName]
		if !ok {
			return tlast.TL2TypeArgument{}, false, someType.PR.BeautifulError(fmt.Errorf("internal error: built-in wrapped type %s not found", tName))
		}
		//tr.T.Type = tlast.Name{Name: tName}
		//tr.T.Bare = false // not required
	}
	isDict, keyRT, elemRT, err := k.IsDictWrapper(tip, tra.Type)
	if err != nil {
		return tlast.TL2TypeArgument{}, false, err
	}
	if isDict {
		keyArg, _, err := k.MigrationArgument(migrateTips, tip, keyRT, leftArgs, false)
		if err != nil {
			return tlast.TL2TypeArgument{}, false, err
		}
		elemType, err := k.MigrationTypeRefImpl(migrateTips, tip, elemRT, leftArgs)
		if err != nil {
			return tlast.TL2TypeArgument{}, false, err
		}
		bracketType := tlast.TL2BracketType{
			ArrayType: elemType,
			IndexType: keyArg,
			HasIndex:  true,
		}
		return tlast.TL2TypeArgument{Type: tlast.TL2TypeRef{BracketType: &bracketType}}, false, nil
	}
	if kt.originTL2 {
		return tlast.TL2TypeArgument{}, false, fmt.Errorf("during migration, reference to TL2 type %s is found", tName)
	}
	_, migrateTip := migrateTips[kt]
	//result := tlast.TL2TypeApplication{Name: tlast.TL2TypeName(tr.Type)}
	//if migrateTip {
	result := tlast.TL2TypeApplication{Name: kt.canonicalName}
	//}
	if len(someType.Arguments) != len(kt.targs) {
		return tlast.TL2TypeArgument{}, false, fmt.Errorf("internal error during migration, reference to type %s has wrong # of arguments found", someType.Name)
	}
	for i, arg := range someType.Arguments {
		if migrateTip && kt.targs[i].usedAsNatVariable { // target type removes this arg
			continue
		}
		if arg.IsNumber {
			result.Arguments = append(result.Arguments, arg)
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

func (k *Kernel) MigrationCheckTL2FromTL1Field(refErrList *migrationTL1RefsTL2Errors, fieldDef tlast.Field, migrateTips map[*KernelType]struct{},
	comb *tlast.Combinator, leftArgs []tlast.TemplateArgument) error {
	if fieldDef.IsRepeated {
		for _, rep := range fieldDef.ScaleRepeat.Rep {
			if err := k.MigrationCheckTL2FromTL1Field(refErrList, rep, migrateTips, comb, leftArgs); err != nil {
				return err
			}
		}
		return nil
	}
	return k.MigrationCheckTL2FromTL1Argument(refErrList, tlast.ArithmeticOrType{T: fieldDef.FieldType}, migrateTips, comb, leftArgs)
}

func (k *Kernel) MigrationCheckTL2FromTL1Type(refErrList *migrationTL1RefsTL2Errors, tr tlast.TypeRef, migrateTips map[*KernelType]struct{},
	comb *tlast.Combinator, leftArgs []tlast.TemplateArgument) error {
	return k.MigrationCheckTL2FromTL1Argument(refErrList, tlast.ArithmeticOrType{T: tr}, migrateTips, comb, leftArgs)
}

func (k *Kernel) MigrationCheckTL2FromTL1Argument(refErrList *migrationTL1RefsTL2Errors, tra tlast.ArithmeticOrType, migrateTips map[*KernelType]struct{},
	comb *tlast.Combinator, leftArgs []tlast.TemplateArgument) error {
	if tra.IsArith {
		return nil
	}
	tr := tra.T

	if tr.Type.Namespace == "" {
		for _, targ := range leftArgs {
			if targ.FieldName == tr.Type.Name {
				return nil // no problem, reference to template argument
			}
		}
	}
	tName := tr.Type.String()
	kt, ok := k.tips[tName]
	if !ok {
		return tr.PR.BeautifulError(fmt.Errorf(" type %s does not exist", tName))
	}
	if kt.originTL2 {
		return tr.PR.BeautifulError(fmt.Errorf("prohibited reference from TL1 to TL2 type %s", tName))
	}
	if _, ok := migrateTips[kt]; ok {
		refErrList.addError(comb, tr.PR.BeautifulError(fmt.Errorf("reference from TL1 prevents migration of type %s", tName)))
	}
	for _, arg := range tr.Args {
		if err := k.MigrationCheckTL2FromTL1Argument(refErrList, arg, migrateTips, comb, leftArgs); err != nil {
			return err
		}
	}
	return nil
}
