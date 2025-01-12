package tlcodegen

import (
	"fmt"
	"path/filepath"
	"strings"
)

const (
	PHPFileStart             = "<?php\n"
	PHPRPCFunctionMock       = "__RPC_FUNCTION_MOCK"
	PHPRPCFunctionResultMock = "__RPC_FUNCTION_RESULT_MOCK"
	PHPRPCResponseMock       = "__RPC_RESPONSE_MOCK"
)

type TypeRWPHPData interface {
	PhpClassName(withPath bool, bare bool) string
	PhpClassNameReplaced() bool
	PhpTypeName(withPath bool, bare bool) string
	PhpGenerateCode(code *strings.Builder, bytes bool) error
	// PhpDefaultInit return not null type initialization value
	PhpDefaultInit() string
	// PhpDefaultValue return default value for field of this type (can be null)
	PhpDefaultValue() string
	PhpIterateReachableTypes(reachableTypes *map[*TypeRWWrapper]bool)
	PhpReadMethodCall(targetName string, bare bool, args []string) []string
}

type PhpClassMeta struct {
	UsedOnlyInInternal bool
	UsedInFunctions    bool
}

func (gen *Gen2) generateCodePHP(generateByteVersions []string) error {
	if gen.options.AddFunctionBodies {
		if err := gen.addCodeFile(filepath.Join("VK", "TL", BasicTlPathPHP), BasicTLCodePHP); err != nil {
			return err
		}
		if err := gen.addCodeFile(filepath.Join("VK", "TL", TLInterfacesPathPHP), TLInterfacesCodePHP); err != nil {
			return err
		}
	}

	// select files where to write code
	gen.PhpMarkAllInternalTypes()
	gen.PhpChoosePaths()
	if err := gen.PhpAdditionalFiles(); err != nil {
		return err
	}

	createdTypes := make(map[string]bool)

	for _, wrapper := range gen.generatedTypesList {
		if createdTypes[wrapper.trw.PhpClassName(true, true)] {
			continue
		}
		if !wrapper.PHPNeedsCode() {
			continue
		}
		fmt.Println(fmt.Sprintf(
			"PHP{%[1]s} in GO{%[2]s}",
			wrapper.trw.PhpClassName(false, true),
			wrapper.goGlobalName,
			wrapper.NatParams,
			wrapper.origTL[0].TemplateArguments,
			wrapper.arguments),
		)
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
	if gen.options.AddMetaData {
		if err := gen.phpCreateMeta(); err != nil {
			return err
		}
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

	for wrapper := range internalReachable {
		if !nonInternalReachable[wrapper] {
			wrapper.phpInfo.UsedOnlyInInternal = true
		}
		wrapper.phpInfo.UsedInFunctions = true
	}

	for wrapper := range nonInternalReachable {
		wrapper.phpInfo.UsedInFunctions = true
	}
}

func (gen *Gen2) phpCreateMeta() error {
	var code strings.Builder

	code.WriteString(`<?php

namespace VK\TL;

use VK\TL;


class tl_item {
  /** @var int */
  public $tag = 0;

  /** @var int */
  public $annotations = 0;

  /** @var string **/
  public $tl_name = '';

  /** 
   * @param int $tag
   * @param int $annotations
   * @param string $tl_name
   */
  function __construct($tag, $annotations, $tl_name) {
    $this->tag = $tag;
    $this->annotations = $annotations;
    $this->tl_name = $tl_name;
  }
}

class tl_meta {
  /** @var tl_item[] */
  private $tl_item_by_tag = null;

  /** @var tl_item[] */
  private $tl_item_by_name = null;

  /**
   * @param string $tl_name
   * @return tl_item|null
   */
  function tl_item_by_name($tl_name) {
    if (array_key_exists($tl_name, $this->tl_item_by_name)) {
        return $this->tl_item_by_name[$tl_name];
    }
    return null;
  }

  /**
   * @param int $tl_name
   * @return tl_item|null
   */
  function tl_item_by_tag($tl_tag) {
    if (array_key_exists($tl_tag, $this->tl_item_by_tag)) {
        return $this->tl_item_by_tag[$tl_tag];
    }
    return null;
  }

  function __construct() {`)

	createdTypes := make(map[string]bool)

	for _, wr := range gen.generatedTypesList {
		if createdTypes[wr.trw.PhpClassName(true, true)] {
			continue
		}
		if !wr.PHPNeedsCode() {
			continue
		}
		if _, iStruct := wr.trw.(*TypeRWStruct); iStruct && len(wr.origTL[0].TemplateArguments) == 0 {
			code.WriteString(fmt.Sprintf(`
    $item%08[1]x = new tl_item(0x%08[1]x, 0x%[2]x, "%[3]s");
    $this->tl_item_by_name["%[3]s"] = $item%08[1]x;
    $this->tl_item_by_tag[0x%08[1]x] = $item%08[1]x;`,
				wr.tlTag,
				wr.AnnotationsMask(),
				wr.tlName.String(),
			))
		}
	}

	code.WriteString(`
  }
}
`)
	if err := gen.addCodeFile(filepath.Join("VK", "TL", "meta.php"), code.String()); err != nil {
		return err
	}
	return nil
}

func PHPGetAllReachableTypes(startTypes []*TypeRWWrapper) map[*TypeRWWrapper]bool {
	reachable := make(map[*TypeRWWrapper]bool)
	for _, startType := range startTypes {
		startType.PhpIterateReachableTypes(&reachable)
	}
	return reachable
}

func PHPSpecialMembersTypes(wrapper *TypeRWWrapper) string {
	if wrapper.tlName.String() == PHPRPCFunctionMock {
		return "TL\\RpcFunction"
	}
	if wrapper.tlName.String() == PHPRPCFunctionResultMock {
		return "TL\\RpcFunctionReturnResult"
	}
	if wrapper.tlName.String() == PHPRPCResponseMock {
		return "TL\\RpcResponse"
	}
	return ""
}

func phpFormatArgs(args []string) string {
	s := ""
	for _, arg := range args {
		s += ", " + arg
	}
	return s
}
