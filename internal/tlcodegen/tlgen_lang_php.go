// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

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
	PhpReadMethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree) []string
	PhpWriteMethodCall(targetName string, bare bool, args *TypeArgumentsTree) []string
}

type PhpClassMeta struct {
	UsedOnlyInInternal bool
	UsedInFunctions    bool
}

func (gen *Gen2) generateCodePHP(bytesWhiteList []string) error {
	// select files where to write code
	gen.PhpMarkAllInternalTypes()
	gen.PhpChoosePaths()

	if err := gen.PhpAdditionalFiles(); err != nil {
		return err
	}

	for _, wrapper := range gen.PhpSelectTypesForGeneration() {
		err := phpGenerateCodeForWrapper(gen, wrapper, true, wrapper.PHPGenerateCode)
		if err != nil {
			return err
		}
	}
	return nil
}

func phpGenerateCodeForWrapper(gen *Gen2, wrapper *TypeRWWrapper, createInterfaceIfNeeded bool, codeGenerator func(code *strings.Builder, bytes bool) error) error {
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
				err := phpGenerateCodeForWrapper(gen, wrapper, false, func(code *strings.Builder, bytes bool) error {
					return PhpGenerateInterfaceCode(code, bytes, wrapper, []*TypeRWWrapper{wrapper})
				})
				if err != nil {
					return err
				}
			}
		}
	}

	filepathName := filepath.Join(wrapper.PHPFilePath(createInterfaceIfNeeded)...)
	if err := gen.addCodeFile(filepathName, code.String()); err != nil {
		return err
	}
	return nil
}

func (gen *Gen2) PhpChoosePaths() {

}

func (gen *Gen2) PhpSelectTypesForGeneration() []*TypeRWWrapper {
	createdTypes := make(map[string]bool)
	wrappers := make([]*TypeRWWrapper, 0)

	for _, wrapper := range gen.generatedTypesList {
		if createdTypes[wrapper.trw.PhpClassName(true, true)] {
			continue
		}
		if !wrapper.PHPNeedsCode() {
			continue
		}
		createdTypes[wrapper.trw.PhpClassName(true, true)] = true
		wrappers = append(wrappers, wrapper)
	}
	return wrappers
}

func (gen *Gen2) PhpAdditionalFiles() error {
	if gen.options.AddFunctionBodies {
		if err := gen.addCodeFile(filepath.Join("VK", "TL", BasicTlPathPHP), BasicTLCodePHP); err != nil {
			return err
		}
		if err := gen.addCodeFile(filepath.Join("VK", "TL", TLInterfacesPathPHP), TLInterfacesCodePHP); err != nil {
			return err
		}
	}
	if gen.options.AddRPCTypes {
		if err := gen.addCodeFile(filepath.Join("VK", "TL", "RpcFunction.php"), fmt.Sprintf(RpcFunctionPHP, gen.copyrightText)); err != nil {
			return err
		}
		if err := gen.addCodeFile(filepath.Join("VK", "TL", "RpcResponse.php"), fmt.Sprintf(RpcResponsePHP, gen.copyrightText)); err != nil {
			return err
		}
		if gen.options.AddFetchers {
			if err := gen.addCodeFile(filepath.Join("RPCFunctionFetcher.php"), fmt.Sprintf(RpcFunctionFetchersPHP, gen.copyrightText)); err != nil {
				return err
			}
		}
	}
	if gen.options.AddMetaData {
		if err := gen.phpCreateMeta(); err != nil {
			return err
		}
	}
	if gen.options.AddFactoryData {
		if err := gen.phpCreateFactory(); err != nil {
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
		if gen.options.AddRPCTypes && rpcResults[wrapper.tlName.String()] {
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

	code.WriteString(fmt.Sprintf(`<?php

%snamespace VK\TL;

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
  private $tl_item_by_tag = [];

  /** @var tl_item[] */
  private $tl_item_by_name = [];

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
   * @param int $tl_tag
   * @return tl_item|null
   */
  function tl_item_by_tag($tl_tag) {
    if (array_key_exists($tl_tag, $this->tl_item_by_tag)) {
        return $this->tl_item_by_tag[$tl_tag];
    }
    return null;
  }

  function __construct() {`, gen.copyrightText))

	for _, wr := range gen.PhpSelectTypesForGeneration() {
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

func (gen *Gen2) phpCreateFactory() error {
	addFactory := func(wr *TypeRWWrapper) bool {
		_, iStruct := wr.trw.(*TypeRWStruct)
		return iStruct && len(wr.origTL[0].TemplateArguments) == 0 && wr.PHPUnionParent() == nil
	}

	var code strings.Builder

	includes := ""

	for _, wr := range gen.PhpSelectTypesForGeneration() {
		//if addFactory(wr) {
		includes += fmt.Sprintf("include \"%s\";\n", filepath.Join(wr.PHPFilePath(true)[2:]...))
		//}
	}

	includesOfRPC := ""
	if gen.options.AddRPCTypes {
		includesOfRPC = `
include "RpcFunction.php";
include "RpcResponse.php";`
	}

	code.WriteString(fmt.Sprintf(`<?php

%[1]snamespace VK\TL;

use VK\TL;

include "tl_interfaces.php";%[3]s

%[2]s
class tl_factory {
  /** @var mixed[] */ // TODO
  private $tl_factory_by_tag = [];

  /** @var mixed[] */ // TODO
  private $tl_factory_by_name = [];

  /**
   * @param string $tl_name
   * @return TL\TL_Object|null
   */
  function tl_object_by_name($tl_name) {
    if (array_key_exists($tl_name, $this->tl_factory_by_name)) {
        return $this->tl_factory_by_name[$tl_name]();
    }
    return null;
  }

  /**
   * @param int $tl_tag
   * @return TL\TL_Object|null
   */
  function tl_object_by_tag($tl_tag) {
    if (array_key_exists($tl_tag, $this->tl_factory_by_tag)) {
        return $this->tl_factory_by_tag[$tl_tag]();
    }
    return null;
  }

  function __construct() {`, gen.copyrightText, includes, includesOfRPC))

	for _, wr := range gen.PhpSelectTypesForGeneration() {
		if addFactory(wr) {
			code.WriteString(fmt.Sprintf(`
    $item%08[1]x = function () { return new %[4]s(); };
    $this->tl_factory_by_name["%[3]s"] = $item%08[1]x;
    $this->tl_factory_by_tag[0x%08[1]x] = $item%08[1]x;`,
				wr.tlTag,
				wr.AnnotationsMask(),
				wr.tlName.String(),
				wr.trw.PhpClassName(true, true),
			))
		}
	}

	code.WriteString(`
  }
}
`)
	if err := gen.addCodeFile(filepath.Join("VK", "TL", "factory.php"), code.String()); err != nil {
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

func phpFormatArgs(args []string, isFirst bool) string {
	s := ""
	for i, arg := range args {
		if isFirst && i == 0 {
			s += arg
		}
		s += ", " + arg
	}
	return s
}

func phpFunctionCommentFormat(argNames []string, argTypes []string, returnType string, shift string) string {
	if len(argNames) != len(argTypes) {
		return ""
	}
	result := make([]string, 0)
	result = append(result, shift+"/**")
	if len(argNames) == 0 {
		result = append(result, shift+" * @kphp-inline")
	} else {
		for i := range argNames {
			result = append(result, shift+fmt.Sprintf(" * @param $%[1]s %[2]s", argNames[i], argTypes[i]))
		}
	}
	if returnType != "" {
		result = append(result, shift+" *")
		result = append(result, shift+" * @return "+returnType)
	}
	result = append(result, shift+" */")
	return strings.Join(result, "\n")
}
