package tlcodegen

import (
	"fmt"
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

	for _, wrapper := range gen.generatedTypesList {
		var code strings.Builder
		// add
		code.WriteString(PHPFileStart)
		// add copyright text
		code.WriteString(gen.copyrightText)

		wrapper.PHPGenerateCode(&code, true)

		fmt.Printf("TL[%[1]s] = Go {%[2]s, %[4]s} -> PHP {%[3]s, %[5]s}\n",
			wrapper.tlName.String(),
			wrapper.goGlobalName,
			wrapper.trw.PhpClassName(true),
			reflect.TypeOf(wrapper.trw),
			wrapper.trw.PhpTypeName(true),
		)

		//filepathName := wrapper.phpInfo.FileName
		//if err := gen.addCodeFile(filepathName, code.String()); err != nil {
		//	return err
		//}
	}
	return nil
}

func (gen *Gen2) PhpChoosePaths() {

}
