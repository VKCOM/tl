// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package puregen

import "flag"

type OptionsGo struct {
	BasicPackageNameFull   string // if empty, will be created
	TLPackageNameFull      string
	BasicRPCPath           string
	RawHandlerWhileList    string
	GenerateLegacyJsonRead bool
	UseCheckLengthSanity   bool

	SplitInternal bool
}

func (opt *OptionsGo) Bind(f *flag.FlagSet) {
	f.StringVar(&opt.BasicPackageNameFull, "basicPkgPath", "",
		"if empty, 'basictl' package will be generated in output dir, otherwise imports will be generated")
	f.StringVar(&opt.TLPackageNameFull, "pkgPath", "",
		"package path to be used inside generated code")
	f.StringVar(&opt.BasicRPCPath, "basicRPCPath", "",
		"path to rpc package")
	f.StringVar(&opt.RawHandlerWhileList, "rawHandlerWhiteList", "",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate RAW function handlers. Empty means none, '*' means all")
	f.BoolVar(&opt.GenerateLegacyJsonRead, "generateLegacyJsonRead", false,
		"whether to generate methods to read json in old way")
	f.BoolVar(&opt.UseCheckLengthSanity, "checkLengthSanity", true,
		"enable feature to generate code to check length sanity of arrays (default:true)")

	f.BoolVar(&opt.SplitInternal, "split-internal", false,
		"generated code will be split into independent packages (in a simple word: speeds up compilation)")
}
