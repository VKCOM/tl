// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"strings"

	"github.com/vkcom/tl/internal/tlast"
)

type Whitelist struct {
	filters []string
}

func NewWhiteList(filter string) Whitelist {
	var result Whitelist
	for _, str := range strings.Split(filter, ",") {
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		result.filters = append(result.filters, str)
	}
	return result
}

func (w *Whitelist) Empty() bool {
	return len(w.filters) == 0
}

func (w *Whitelist) HasNamespace(ns string) bool {
	inNameFilterElement := func(ns string, filter string) bool {
		if filter == "*" {
			return true
		}
		if strings.HasSuffix(filter, ".") {
			return ns == strings.TrimSuffix(filter, ".")
		}
		return false
	}

	for _, filter := range w.filters {
		if inNameFilterElement(ns, filter) {
			return true
		}
	}
	return false
}

func (w *Whitelist) HasName(name tlast.Name) bool {
	inNameFilterElement := func(name tlast.Name, filter string) bool {
		if filter == "*" {
			return true
		}
		if strings.HasSuffix(filter, ".") {
			return name.Namespace == strings.TrimSuffix(filter, ".")
		}
		return name.String() == filter
	}

	for _, filter := range w.filters {
		if inNameFilterElement(name, filter) {
			return true
		}
	}
	return false
}
