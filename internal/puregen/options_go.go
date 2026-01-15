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
	GenerateRPCCode        bool
	BasicRPCPath           string
	RawHandlerWhileList    string
	GenerateRandomCode     bool
	GenerateLegacyJsonRead bool
	UseCheckLengthSanity   bool

	SplitInternal bool
}

func (opt *OptionsGo) Bind(f *flag.FlagSet) {
	flag.StringVar(&opt.BasicPackageNameFull, "basicPkgPath", "",
		"if empty, 'basictl' package will be generated in output dir, otherwise imports will be generated")
	flag.StringVar(&opt.TLPackageNameFull, "pkgPath", "",
		"package path to be used inside generated code")
	flag.BoolVar(&opt.GenerateRPCCode, "generateRPCCode", false,
		"whether to generate *_server.go files")
	flag.StringVar(&opt.BasicRPCPath, "basicRPCPath", "",
		"path to rpc package")
	flag.StringVar(&opt.RawHandlerWhileList, "rawHandlerWhiteList", "",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate RAW function handlers. Empty means none, '*' means all")
	flag.BoolVar(&opt.GenerateRandomCode, "generateRandomCode", false,
		"whether to generate methods for random filling structs")
	flag.BoolVar(&opt.GenerateLegacyJsonRead, "generateLegacyJsonRead", false,
		"whether to generate methods to read json in old way")
	flag.BoolVar(&opt.UseCheckLengthSanity, "checkLengthSanity", true,
		"enable feature to generate code to check length sanity of arrays (default:true)")

	flag.BoolVar(&opt.SplitInternal, "split-internal", false,
		"generated code will be split into independent packages (in a simple word: speeds up compilation)")
}
