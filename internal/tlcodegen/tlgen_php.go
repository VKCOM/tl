package tlcodegen

import (
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

func (gen *Gen2) PhpGeneratecode() error {
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

		filepathName := wrapper.phpInfo.FileName
		if err := gen.addCodeFile(filepathName, code.String()); err != nil {
			return err
		}
	}
	return nil
}

func (gen *Gen2) PhpChoosePaths() {

}
