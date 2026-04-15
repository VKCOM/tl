// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package puregen

import "flag"

type OptionsPHP struct {
	AddFunctionBodies            bool
	FunctionsBodiesWhiteList     string
	IgnoreUnusedInFunctionsTypes bool
	AddRPCTypes                  bool
	AddFetchers                  bool
	AddSwitcher                  bool
	AddFetchersEchoComments      bool
	InplaceSimpleStructs         bool
	UseBuiltinDataProviders      bool
	AddTypeComments              bool

	AddMetaData    bool
	AddFactoryData bool
}

func (opt *OptionsPHP) Bind(f *flag.FlagSet) {
	// PHP
	f.BoolVar(&opt.AddFunctionBodies, "php-serialization-bodies", false,
		`whether to generate body to write/read generated structs and functions`)
	f.BoolVar(&opt.AddMetaData, "php-generate-meta", false,
		`whether to generate methods to get meta information about tl objects`)
	f.BoolVar(&opt.AddFactoryData, "php-generate-factory", false,
		`whether to generate factory of tl objects`)
	f.BoolVar(&opt.IgnoreUnusedInFunctionsTypes, "php-ignore-unused-types", true,
		`whether to not generate types without usages in functions (default:true)`)
	f.BoolVar(&opt.AddRPCTypes, "php-rpc-support", true,
		`whether to generate special rpc types (default:true)`)
	f.BoolVar(&opt.AddFetchers, "php-generate-fetchers", false,
		`whether to generate new fetchers for kphp compiler integration (requires --php-rpc-support=true)`)
	f.BoolVar(&opt.AddSwitcher, "php-generate-switcher", false,
		`whether to generate switcher for new functions for kphp compiler integration (requires --php-rpc-support=true)`)
	f.BoolVar(&opt.InplaceSimpleStructs, "php-inplace-simple-structs", true,
		`whether to avoid generation of structs with no arguments and only 1 field (default:true)`)
	f.BoolVar(&opt.UseBuiltinDataProviders, "php-use-builtin-data-providers", false,
		`whether to use builtin functions to store / fetch data instead of stream api`)
	f.BoolVar(&opt.AddFetchersEchoComments, "php-generate-fetchers-echo-comment", true,
		`whether to generate echo of usage for custom store/fetch`)
	f.StringVar(&opt.FunctionsBodiesWhiteList, "php-serialization-bodies-whitelist", "*",
		`comma-separated list of fully-qualified top-level namespaces (if have trailing '.'), to generate code for serialization function bodies. Empty means none, '*' means all (require --php-serialization-bodies=true)`)
	f.BoolVar(&opt.AddTypeComments, "php-add-type-comments", false,
		`whether to generate comment with type combinator for each type`)

	//f.BoolVar(&opt.CreateTLFilesWithAllTypesInReturn, "php-create-tl-files-with-all-types-in-return", false,
	//	`whether to create duplicates of passed tl files with all top level types in function return (option for testing)`)
	//f.BoolVar(&opt.CreateTLSplitedFilesForEachNamespace, "php-create-tl-splited-files-for-namespaces", false,
	//	`whether to create for each mentioned namespace separate file with all required dependencies (requires --php-create-tl-splited-files-for-namespaces-support-folder!="")`)
	//f.StringVar(&opt.CreateTLSplitedFilesForEachNamespaceFolder, "php-create-tl-splited-files-for-namespaces-folder", "",
	//	`folder to create splited files for such option (requires --php-create-tl-splited-files-for-namespaces-support=true)`)
}
