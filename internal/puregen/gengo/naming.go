// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package gengo

import (
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/utils"
)

func canonicalGoName(name tlast.Name, insideNamespace string) string {
	if name.Namespace == insideNamespace {
		return utils.CNameToCamelName(name.Name)
	}
	return utils.CNameToCamelName(name.Namespace) + utils.CNameToCamelName(name.Name)
}

func canonicalGoName2(name tlast.TL2TypeName, insideNamespace string) string {
	return canonicalGoName(tlast.Name(name), insideNamespace)
}
