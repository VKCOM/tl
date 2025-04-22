// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package tlcodegen

import "strings"

const basicCPPTLIOStreamsPath = "basictl/io_streams.h"
const basicCPPTLIOThrowableStreamsPath = "basictl/io_throwable_streams.h"

const basictlPackage = "basictl"
const basictlCppIncludeStart = "/** TLGEN: CPP INCLUDES */\n"
const basictlCppIncludeEnd = "/** TLGEN: CPP INCLUDES END */\n"

const NoNamespaceGroup = ""
const SpecialGroupPrefix = "__"

const CommonGroup = SpecialGroupPrefix + "common_namespace"
const IndependentTypes = CommonGroup
const GhostTypes = SpecialGroupPrefix + "ghosts"

const CppPrintGraphvizRepresentation = false
const CppPrintNamespaceDependencies = false

func cppIsSpecialNamespace(namespace string) bool {
	return namespace == NoNamespaceGroup || strings.HasPrefix(namespace, SpecialGroupPrefix)
}
