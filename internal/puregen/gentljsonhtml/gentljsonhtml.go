// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gentljsonhtml

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/vkcom/tl/internal/pure"
	"github.com/vkcom/tl/internal/puregen"
	"github.com/vkcom/tl/internal/utils"
)

func Generate(kernel *pure.Kernel, options *puregen.Options) error {
	if err := kernel.Compile(); err != nil {
		return err
	}
	if options.Kernel.Verbose {
		fmt.Printf("generating tljson.html documentation...\n")
	}
	if options.Outfile == "" {
		return fmt.Errorf("--outfile should not be empty")
	}
	sortedInstances := kernel.AllTypeInstances()
	slices.SortStableFunc(sortedInstances, func(a, b pure.TypeInstance) int {
		return cmp.Compare(a.CanonicalName(), b.CanonicalName())
	})

	buf := tlJSON(kernel, options, sortedInstances, utils.AppVersion())
	if err := os.WriteFile(options.Outfile, []byte(buf), 0644); err != nil {
		return fmt.Errorf("error writing tljson file %q: %w", options.Outfile, err)
	}
	return nil
}

func JSONHelpString(kernel *pure.Kernel, ins pure.TypeInstance) string {
	if ins.KernelType() == nil {
		return ins.CanonicalName()
	}
	var s strings.Builder
	s.WriteString(ins.KernelType().CanonicalName().String())
	rt := ins.Common().ResolvedType()
	//if br := rt.BracketType; br != nil {
	//	if br.HasIndex {
	//
	//	}
	//}
	if len(rt.SomeType.Arguments) == 0 {
		return s.String()
	}
	s.WriteByte('<')
	for i, a := range rt.SomeType.Arguments {
		// fieldName := t.origTL[0].TemplateArguments[i].FieldName // arguments must be the same for all union elements
		if i != 0 {
			s.WriteByte(',')
		}
		if a.IsNumber {
			s.WriteString(strconv.FormatUint(uint64(a.Number), 10))
		} else if a.Type.String() == "*" {
			s.WriteString("#") // TODO - write fieldName here if special argument to function is set
		} else {
			ref, _, err := kernel.GetInstance(a.Type)
			if err != nil {
				panic(fmt.Errorf("internal error: cannot get type of argument %s: %w", a.Type, err))
			}
			s.WriteString(JSONHelpString(kernel, ref))
		}
	}
	s.WriteByte('>')
	return s.String()
}

func JSONHelpFullType(kernel *pure.Kernel, ins pure.TypeInstance, bare bool, fields []pure.Field, natArgs []pure.ActualNatArg) string {
	result := helpString2(kernel, ins, bare, fields, &natArgs)
	if len(natArgs) != 0 {
		panic("JSONHelpFullType should consume all arguments")
	}
	return result
}

func JSONHelpNatArg(ins pure.TypeInstance, fields []pure.Field, natArg pure.ActualNatArg) string {
	if natArg.IsNumber() {
		return fmt.Sprintf("%d", natArg.Number())
	}
	if natArg.IsField() {
		return fields[natArg.FieldIndex()].Name()
	}
	return natArg.Name()
}

func helpString2(kernel *pure.Kernel, ins pure.TypeInstance, bare bool, fields []pure.Field, natArgs *[]pure.ActualNatArg) string {
	var s strings.Builder
	if ins.KernelType() == nil {
		return ins.CanonicalName()
	}
	rt := ins.Common().ResolvedType()
	//if br := rt.BracketType; br != nil {
	//	if br.HasIndex {
	//		if br.IndexType.IsNumber || br.IndexType.Type.String() == "*" {
	//			head = "BuiltinTuple"
	//			w.resolvedT2GoNameArg(&b, br.IndexType, insideNamespace)
	//			//if br.IndexType.IsNumber {
	//			//	b.WriteString(strconv.FormatUint(uint64(br.IndexType.Number), 10))
	//			//}
	//		} else {
	//			head = "BuiltinDict"
	//			w.resolvedT2GoNameArg(&b, br.IndexType, insideNamespace)
	//		}
	//	} else {
	//		head = "BuiltinVector"
	//	}
	//	w.resolvedT2GoNameArg(&b, tlast.TL2TypeArgument{Type: br.ArrayType}, insideNamespace)
	//} else {
	//	head = canonicalGoName(w.goCanonicalName, insideNamespace)
	//	for _, arg := range rt.SomeType.Arguments {
	//		w.resolvedT2GoNameArg(&b, arg, insideNamespace)
	//	}
	//}

	s.WriteString(ins.KernelType().CanonicalName().String())
	if len(rt.SomeType.Arguments) == 0 {
		return s.String()
	}
	s.WriteString("<")
	for i, a := range rt.SomeType.Arguments {
		if i != 0 {
			s.WriteString(",")
		}
		if a.IsNumber {
			s.WriteString(fmt.Sprintf("%d", a.Number))
		} else if a.Type.String() == "*" {
			natArg := (*natArgs)[0]
			*natArgs = (*natArgs)[1:]
			if natArg.IsField() {
				s.WriteString(fields[natArg.FieldIndex()].Name())
			} else {
				s.WriteString(natArg.Name())
			}
		} else {
			ref, fieldBare, err := kernel.GetInstance(a.Type)
			if err != nil {
				panic(fmt.Errorf("internal error: cannot get type of argument %s: %w", a.Type, err))
			}
			s.WriteString(helpString2(kernel, ref, fieldBare, fields, natArgs))
		}
	}
	s.WriteString(">")
	return s.String()
}

func IsEmptyStruct(ins pure.TypeInstance) bool {
	structElement, ok := ins.(*pure.TypeInstanceStruct)
	if !ok {
		return false
	}
	return len(structElement.Fields()) == 0
}
