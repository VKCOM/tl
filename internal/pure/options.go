// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import "flag"

type OptionsKernel struct {
	WarningsAreErrors bool

	TypesWhiteList string
	GenerateTL2    bool
	TL2WhiteList   string
}

func (opt *OptionsKernel) Bind(f *flag.FlagSet) {
	f.BoolVar(&opt.WarningsAreErrors, "Werror", false,
		"treat all warnings as errors")
	f.StringVar(&opt.TypesWhiteList, "typesWhiteList", "*",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate code. Empty means none, '*' means all")
	f.BoolVar(&opt.GenerateTL2, "tl2-generate", false,
		"generate code for tl2 methods (currently work only for golang)")
	f.StringVar(&opt.TL2WhiteList, "tl2WhiteList", "*",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate TL2 code. Empty means none, '*' means all")
}
