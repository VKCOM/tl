// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
	"github.com/vkcom/tl/internal/tlast"
)

type Whitelist struct {
	whiltelistName string
	filters        []string
	used           []bool
	asterisk       bool // it is always considered unused
}

func NewWhiteList(whiltelistName string, filter string) Whitelist {
	result := Whitelist{whiltelistName: whiltelistName}
	for _, str := range strings.Split(filter, ",") {
		str = strings.TrimSpace(str)
		if str == "" {
			continue
		}
		if str == "*" {
			result.asterisk = true
			continue
		}
		result.filters = append(result.filters, str)
		result.used = append(result.used, false)
	}
	return result
}

func (w *Whitelist) Empty() bool {
	return len(w.filters) == 0
}

func (w *Whitelist) HasName2(name tlast.TL2TypeName) bool {
	return w.HasName(tlast.Name(name))
}

func (w *Whitelist) HasName(name tlast.Name) bool {
	if w.asterisk {
		return true
	}
	inNameFilterElement := func(name tlast.Name, filter string) bool {
		if strings.HasSuffix(filter, ".") {
			return name.Namespace == strings.TrimSuffix(filter, ".")
		}
		return name.String() == filter
	}

	for i, filter := range w.filters {
		if inNameFilterElement(name, filter) {
			w.used[i] = true
			return true
		}
	}
	return false
}

func (w *Whitelist) UnusedWarning() error {
	var unusedFilters []string
	for i, used := range w.used {
		if !used {
			unusedFilters = append(unusedFilters, w.filters[i])
		}
	}
	if len(unusedFilters) != 0 {
		return fmt.Errorf("%s: unused filters in %s whitelist: %s", color.InYellow("warning"), w.whiltelistName, strings.Join(unusedFilters, ", "))
	}
	return nil
}
