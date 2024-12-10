package tlcodegen

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
)

type PhpClassMeta struct {
	FileName  string
	GroupName string

	ClassName string
}

const (
	PHPFileStart = "<?php\n"
	//PHPFileEnd   = "?>\n"
)

func (gen *Gen2) generateCodePHP(generateByteVersions []string) error {
	if err := gen.addCodeFile(BasicTlPathPhp, BasicTLCodePHP); err != nil {
		return err
	}

	// select files where to write code
	gen.PhpChoosePaths()

	createdTypes := make(map[string]bool)

	for _, wrapper := range gen.generatedTypesList {
		if wrapper.PHPTypePath() == "" {
			continue
		}
		if wrapper.PHPIsPrimitiveType() {
			continue
		}
		if createdTypes[wrapper.trw.PhpClassName(true)] {
			continue
		}
		if strct, isStrct := wrapper.trw.(*TypeRWStruct); isStrct && strct.ResultType != nil && strct.wr.HasAnnotation("internal") {
			continue
		}
		var code strings.Builder
		// add start symbol
		code.WriteString(PHPFileStart)
		code.WriteString("\n")
		// add copyright text
		code.WriteString(gen.copyrightText)
		code.WriteString(fmt.Sprintf("namespace VK\\%s;\n", strings.Join(wrapper.PHPTypePathElements(), "\\")))

		if err := wrapper.PHPGenerateCode(&code, true); err != nil {
			return err
		}

		fmt.Printf("TL[%[1]s] = Go {%[2]s, %[4]s} -> PHP {%[3]s, %[5]s}\n",
			wrapper.tlName.String(),
			wrapper.goGlobalName,
			wrapper.trw.PhpClassName(true),
			reflect.TypeOf(wrapper.trw),
			wrapper.trw.PhpTypeName(true),
		)

		fmt.Printf("Core[%s] = %s, %s\n", wrapper.goGlobalName, wrapper.PHPGenCoreType().goGlobalName, reflect.TypeOf(wrapper.PHPGenCoreType().trw))

		filepathParts := []string{"VK"}
		filepathParts = append(filepathParts, wrapper.PHPTypePathElements()...)
		filepathParts = append(filepathParts, fmt.Sprintf("%s.php", wrapper.trw.PhpClassName(false)))
		filepathName := filepath.Join(filepathParts...)
		if err := gen.addCodeFile(filepathName, code.String()); err != nil {
			return err
		}
		createdTypes[wrapper.trw.PhpClassName(true)] = true
	}
	return nil
}

func (gen *Gen2) PhpChoosePaths() {

}
