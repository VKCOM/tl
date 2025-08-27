// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import (
	"fmt"
	"github.com/vkcom/tl/internal/utils"
	"path/filepath"
	"sort"
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
	PhpReadMethodCall(targetName string, bare bool, initIfDefault bool, args *TypeArgumentsTree, supportSuffix string) []string
	PhpWriteMethodCall(targetName string, bare bool, args *TypeArgumentsTree, supportSuffix string) []string
}

type PhpClassMeta struct {
	UsedOnlyInInternal bool
	UsedInFunctions    bool

	IsDuplicate bool
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
	createdTypes := make(map[string]int)
	wrappers := make([]*TypeRWWrapper, 0)

	duplicates := make(map[string][]*TypeRWWrapper)

	for _, wrapper := range gen.generatedTypesList {
		phpTypeName := wrapper.trw.PhpClassName(true, true)
		if i, ok := createdTypes[phpTypeName]; ok {
			if len(duplicates[phpTypeName]) == 0 {
				duplicates[phpTypeName] = append(duplicates[phpTypeName], wrappers[i])
			}
			duplicates[phpTypeName] = append(duplicates[phpTypeName], wrapper)
			continue
		}
		if !wrapper.PHPNeedsCode() {
			continue
		}
		createdTypes[phpTypeName] = len(wrappers)
		wrappers = append(wrappers, wrapper)
	}

	if false {
		duplicatedNames := utils.Keys(duplicates)
		sort.Strings(duplicatedNames)

		for _, name := range duplicatedNames {
			fmt.Printf("Duplicates for php type %q:\n", name)
			for i, wrapper := range duplicates[name] {
				if i != 0 {
					wrapper.phpInfo.IsDuplicate = true
				}
				fmt.Printf("\t%s\n", wrapper.goGlobalName)
			}
		}
	}

	return wrappers
}

func (gen *Gen2) PhpAdditionalFiles() error {
	if gen.options.AddFunctionBodies {
		if gen.options.UseBuiltinDataProviders {
			//if err := gen.addCodeFile(filepath.Join("VK", "TL", TLInterfacesPathPHP), TLInterfacesCodeWithoutStreamPHP); err != nil {
			//	return err
			//}
		} else {
			if err := gen.addCodeFile(filepath.Join("VK", "TL", BasicTlPathPHP), BasicTLCodePHP); err != nil {
				return err
			}
			if err := gen.addCodeFile(filepath.Join("VK", "TL", TLInterfacesPathPHP), TLInterfacesCodePHP); err != nil {
				return err
			}
		}
	}
	if gen.options.AddRPCTypes {
		if gen.options.UseBuiltinDataProviders {
			if err := gen.addCodeFile(filepath.Join("VK", "TL", "RpcFunction.php"), fmt.Sprintf(RpcFunctionWithFetchersPHP, gen.copyrightText)); err != nil {
				return err
			}
		} else {
			if err := gen.addCodeFile(filepath.Join("VK", "TL", "RpcFunction.php"), fmt.Sprintf(RpcFunctionPHP, gen.copyrightText)); err != nil {
				return err
			}
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

func phpAddMetaAndFactory(wr *TypeRWWrapper) bool {
	strct, iStruct := wr.trw.(*TypeRWStruct)
	if iStruct && strct.ResultType != nil {
		if strct.wr.origTL[0].OriginalDescriptor != nil &&
			strct.wr.origTL[0].OriginalDescriptor.OriginalDescriptor != nil {
			return len(strct.wr.origTL[0].OriginalDescriptor.OriginalDescriptor.TemplateArguments) == 0
		}
		return true
	}
	return false
}

func (gen *Gen2) phpCreateMeta() error {
	var code strings.Builder

	code.WriteString(fmt.Sprintf(`<?php

%snamespace VK\TL;

use VK\TL;

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
   * @return tl_item[]
   */
  function all_items_by_names() {
    return $this->tl_item_by_name;
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
		if phpAddMetaAndFactory(wr) {
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
	if err := gen.addCodeFile(filepath.Join("VK", "TL", "tl_meta.php"), code.String()); err != nil {
		return err
	}

	codeItem := strings.Builder{}
	codeItem.WriteString(fmt.Sprintf(`<?php

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
`, gen.copyrightText))

	if err := gen.addCodeFile(filepath.Join("VK", "TL", "tl_item.php"), codeItem.String()); err != nil {
		return err
	}

	return nil
}

func (gen *Gen2) phpCreateFactory() error {
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
%[3]s

%[2]s
class tl_factory {
  /** @var (callable(): RpcFunction)[] */ // TODO
  private $tl_factory_by_tag = [];

  /** @var (callable(): RpcFunction)[] */ // TODO
  private $tl_factory_by_name = [];

  /** @var (callable(RpcFunctionReturnResult): string|null)[] */
  private $tl_json_printer_for_response_by_name = [];

  /**
   * @param string $tl_name
   * @return RpcFunction|null
   */
  function tl_object_by_name($tl_name) {
    if (array_key_exists($tl_name, $this->tl_factory_by_name)) {
        return $this->tl_factory_by_name[$tl_name]();
    }
    return null;
  }

  /**
   * @param int $tl_tag
   * @return RpcFunction|null
   */
  function tl_object_by_tag($tl_tag) {
    if (array_key_exists($tl_tag, $this->tl_factory_by_tag)) {
        return $this->tl_factory_by_tag[$tl_tag]();
    }
    return null;
  }

  /**
   * @param string $tl_name
   * @param RpcFunctionReturnResult $result
   * @return string|null
   */
  function tl_json_print_response_by_name($tl_name, $result) {
    if (array_key_exists($tl_name, $this->tl_json_printer_for_response_by_name)) {
        return $this->tl_json_printer_for_response_by_name[$tl_name]($result);
    }
    return null;
  }

  function __construct() {`, gen.copyrightText, includes, includesOfRPC))

	for _, wr := range gen.PhpSelectTypesForGeneration() {
		if phpAddMetaAndFactory(wr) {
			code.WriteString(fmt.Sprintf(`
    /** @var $item%08[1]x (callable(): RpcFunction) */
    $item%08[1]x = function () { return new %[4]s(); };
    /** @var $item_json_print%08[1]x (callable(RpcFunctionReturnResult): string|null) */
    $item_json_print%08[1]x = function($result) {
       if ($result instanceof %[4]s_result) {
         return \JsonEncoder::encode($result);
       }
       return null;
    };
    $this->tl_factory_by_name["%[3]s"] = $item%08[1]x;
    $this->tl_factory_by_tag[0x%08[1]x] = $item%08[1]x;
    $this->tl_json_printer_for_response_by_name["%[3]s"] = $item_json_print%08[1]x;`,
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
	if err := gen.addCodeFile(filepath.Join("VK", "TL", "tl_factory.php"), code.String()); err != nil {
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
		} else {
			s += ", " + arg
		}
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
			result = append(result, shift+fmt.Sprintf(" * @param %[2]s $%[1]s", argNames[i], argTypes[i]))
		}
	}
	if returnType != "" {
		result = append(result, shift+" *")
		result = append(result, shift+" * @return "+returnType)
	}
	result = append(result, shift+" */")
	return strings.Join(result, "\n")
}

func phpFunctionArgumentsFormat(argNames []string) string {
	s := ""
	for i, name := range argNames {
		if i != 0 {
			s += ", "
		}
		s += "$"
		s += name
	}
	return s
}
