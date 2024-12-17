package tlcodegen

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
)

type PhpClassMeta struct {
	UsedOnlyInInternal bool
	UsedInFunctions    bool
}

const (
	PHPFileStart             = "<?php\n"
	PHPRPCFunctionMock       = "__RPC_FUNCTION_MOCK"
	PHPRPCFunctionResultMock = "__RPC_FUNCTION_RESULT_MOCK"
	PHPRPCResponseMock       = "__RPC_RESPONSE_MOCK"
)

func (gen *Gen2) generateCodePHP(generateByteVersions []string) error {
	if err := gen.addCodeFile(BasicTlPathPhp, BasicTLCodePHP); err != nil {
		return err
	}

	// select files where to write code
	gen.PhpMarkAllInternalTypes()
	gen.PhpChoosePaths()
	if err := gen.PhpAdditionalFiles(); err != nil {
		return err
	}

	createdTypes := make(map[string]bool)

	for _, wrapper := range gen.generatedTypesList {
		if wrapper.tlName.String() == "rpcResponseOk" {
			print(wrapper.trw.PhpClassName(true, true))
		}
		if createdTypes[wrapper.trw.PhpClassName(true, true)] {
			continue
		}
		if !wrapper.PHPNeedsCode() {
			continue
		}
		err := phpGenerateCodeForWrapper(gen, wrapper, createdTypes, true, wrapper.PHPGenerateCode)
		if err != nil {
			return err
		}
	}
	return nil
}

func phpGenerateCodeForWrapper(gen *Gen2, wrapper *TypeRWWrapper, createdTypes map[string]bool, createInterfaceIfNeeded bool, codeGenerator func(code *strings.Builder, bytes bool) error) error {
	var code strings.Builder
	// add start symbol
	code.WriteString(PHPFileStart)
	code.WriteString("\n")
	// add copyright text
	code.WriteString(gen.copyrightText)
	code.WriteString(fmt.Sprintf("namespace VK\\%s;\n", strings.Join(wrapper.PHPTypePathElements(), "\\")))

	if err := codeGenerator(&code, true); err != nil {
		return err
	}

	if createInterfaceIfNeeded {
		if strct, isStruct := wrapper.trw.(*TypeRWStruct); isStruct {
			unionParent := strct.PhpConstructorNeedsUnion()
			if unionParent != nil && unionParent == wrapper {
				err := phpGenerateCodeForWrapper(gen, wrapper, createdTypes, false, func(code *strings.Builder, bytes bool) error {
					return PhpGenerateInterfaceCode(code, bytes, wrapper, []*TypeRWWrapper{wrapper})
				})
				if err != nil {
					return err
				}
			}
		}
	}

	fmt.Printf("TL[%[1]s] = Go {%[2]s, %[4]s} -> PHP {%[3]s, %[5]s}\n",
		wrapper.tlName.String(),
		wrapper.goGlobalName,
		wrapper.trw.PhpClassName(true, true),
		reflect.TypeOf(wrapper.trw),
		wrapper.trw.PhpTypeName(true, true),
	)

	filepathParts := []string{"VK"}
	//filepathParts = append(filepathParts, wrapper.PHPTypePathElements()...)
	path := fmt.Sprintf("%s.php", wrapper.trw.PhpClassName(true, createInterfaceIfNeeded))
	filepathParts = append(filepathParts, strings.Split(path, "\\")...)
	filepathName := filepath.Join(filepathParts...)
	if err := gen.addCodeFile(filepathName, code.String()); err != nil {
		return err
	}
	createdTypes[wrapper.trw.PhpClassName(true, createInterfaceIfNeeded)] = true
	return nil
}

func (gen *Gen2) PhpChoosePaths() {

}

func (gen *Gen2) PhpAdditionalFiles() error {
	if err := gen.addCodeFile(filepath.Join("VK", "TL", "RpcFunction.php"), fmt.Sprintf(RpcFunctionPHP, gen.copyrightText)); err != nil {
		return err
	}
	if err := gen.addCodeFile(filepath.Join("VK", "TL", "RpcResponse.php"), fmt.Sprintf(RpcResponsePHP, gen.copyrightText)); err != nil {
		return err
	}
	return nil
}

func (gen *Gen2) PhpMarkAllInternalTypes() {
	rpcResults := map[string]bool{
		"rpcResponseError":  true,
		"rpcResponseHeader": true,
		"rpcResponseOk":     true,
	}
	internalFunctions := make([]*TypeRWWrapper, 0)
	nonInternalFunctions := make([]*TypeRWWrapper, 0)
	for _, wrapper := range gen.generatedTypesList {
		if strct, isStrct := wrapper.trw.(*TypeRWStruct); isStrct && strct.ResultType != nil {
			if strct.wr.HasAnnotation("internal") {
				internalFunctions = append(internalFunctions, wrapper)
			} else {
				nonInternalFunctions = append(nonInternalFunctions, wrapper)
			}
		}
		// TODO: CHANGE SOMEHOW
		if rpcResults[wrapper.tlName.String()] {
			nonInternalFunctions = append(nonInternalFunctions, wrapper)
		}
	}
	internalReachable := PHPGetAllReachableTypes(internalFunctions)
	nonInternalReachable := PHPGetAllReachableTypes(nonInternalFunctions)

	for wrapper, _ := range internalReachable {
		if !nonInternalReachable[wrapper] {
			wrapper.phpInfo.UsedOnlyInInternal = true
		}
		wrapper.phpInfo.UsedInFunctions = true
	}

	for wrapper, _ := range nonInternalReachable {
		wrapper.phpInfo.UsedInFunctions = true
	}
}

func PHPGetAllReachableTypes(startTypes []*TypeRWWrapper) map[*TypeRWWrapper]bool {
	reachable := make(map[*TypeRWWrapper]bool)
	for _, startType := range startTypes {
		startType.PhpIterateReachableTypes(&reachable)
	}
	return reachable
}
