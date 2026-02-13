// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gentljsonhtml

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/vkcom/tl/internal/pure"
	"github.com/vkcom/tl/internal/puregen"
	"github.com/vkcom/tl/internal/tlast"
)

func Generate(kernel *pure.Kernel, options *puregen.Options) error {
	if options.Verbose {
		log.Print("generating file with combinators in canonical form,,,")
	}
	if options.Outfile == "" {
		return fmt.Errorf("--outfile should not be empty")
	}
	var buf bytes.Buffer
	tlast.TL(kernel.TL1()).WriteGenerate2TL(&buf)
	if err := os.WriteFile(options.Outfile, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("error writing canonical form file %q: %w", options.Outfile, err)
	}
	return nil
}

func JSONHelpString(ins pure.TypeInstance) string {
	return ins.CanonicalName()
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
	// TODO implement
	//if len(w.origTL) > 1 {
	//	if bare {
	//		panic("helpString2 of bare union")
	//	}
	//	s.WriteString(w.origTL[0].TypeDecl.Name.String())
	//} else {
	//	if bare {
	//		s.WriteString(w.origTL[0].Construct.Name.String())
	//	} else {
	//		s.WriteString(w.origTL[0].TypeDecl.Name.String())
	//	}
	//}
	s.WriteString(ins.KernelType().CanonicalName().String())
	rt := ins.Common().ResolvedType()
	if len(rt.Args) == 0 {
		return s.String()
	}
	s.WriteString("<")
	for i, a := range rt.Args {
		if i != 0 {
			s.WriteString(",")
		}
		if a.IsArith {
			s.WriteString(fmt.Sprintf("%d", a.Arith.Res))
		} else if a.T.String() == "*" {
			natArg := (*natArgs)[0]
			*natArgs = (*natArgs)[1:]
			if natArg.IsField() {
				s.WriteString(fields[natArg.FieldIndex()].Name())
			} else {
				s.WriteString(natArg.Name())
			}
		} else {
			ref, fieldBare, err := kernel.GetInstanceTL1(a.T)
			if err != nil {
				panic(fmt.Errorf("internal error: cannot get type of argument %s: %w", a.T, err))
			}
			s.WriteString(helpString2(kernel, ref, fieldBare, fields, natArgs))
		}
	}
	s.WriteString(">")
	return s.String()
}
