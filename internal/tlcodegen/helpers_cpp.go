// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import "strings"

//const basictlPackage = "basictl"
//const basicCPPTLIOStreamsPath = basictlPackage + "/io_streams.h"
//const basicCPPTLIOThrowableStreamsPath = basictlPackage + "/io_throwable_streams.h"

func CppBasictlPackage(gen2 *Gen2) string {
	if gen2.options.BasicTLNamespace != "" {
		return gen2.options.BasicTLNamespace
	}
	return "basictl"
}

func CppBasicTLIOStreamsPath(gen2 *Gen2) string {
	return CppBasictlPackage(gen2) + "/io_streams.h"
}

func CppBasicTLIOThrowableStreamsPath(gen2 *Gen2) string {
	return CppBasictlPackage(gen2) + "/io_throwable_streams.h"
}

const basictlCppIncludeStart = "/** TLGEN: CPP INCLUDES */\n"
const basictlCppIncludeEnd = "/** TLGEN: CPP INCLUDES END */\n"

const NoNamespaceGroup = ""
const SpecialGroupPrefix = "__"

const CommonGroup = "common"
const IndependentTypes = CommonGroup
const GhostTypes = SpecialGroupPrefix + "ghosts"

const CppPrintGraphvizRepresentation = false
const CppPrintNamespaceDependencies = false

func cppIsSpecialNamespace(namespace string) bool {
	return namespace == NoNamespaceGroup || strings.HasPrefix(namespace, SpecialGroupPrefix) || namespace == CommonGroup
}
