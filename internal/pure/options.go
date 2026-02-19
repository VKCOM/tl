// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"flag"
	"io"
)

type OptionsKernel struct {
	WarningsAreErrors bool

	TypesWhiteList string
	TL2WhiteList   string // if !empty, will generate also TL2 factory, meta, etc.

	RPCPreferTL2WhiteList string // RPC client code will prefer to send requests in TL2 format

	ErrorWriter io.Writer // all Errors and warnings should be redirected to this io.Writer, by default it is os.Stderr

	// TODO - remove after migration code stabilized
	TL2MigrationDevMode bool
	// TODO - quickly adapt new rules, remove these options
	NewBrackets bool // TODO - implement new brackets
}

func (opt *OptionsKernel) Bind(f *flag.FlagSet) {
	f.BoolVar(&opt.WarningsAreErrors, "Werror", false,
		"treat all warnings as errors")
	f.StringVar(&opt.TypesWhiteList, "typesWhiteList", "*",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate code. Empty means none, '*' means all")
	f.StringVar(&opt.TL2WhiteList, "tl2WhiteList", "",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to generate TL2 code. Empty means none, '*' means all")
	f.StringVar(&opt.RPCPreferTL2WhiteList, "rpcPreferTL2WhiteList", "",
		"comma-separated list of fully-qualified top-level types or namespaces (if have trailing '.'), to prefer sending RPC requests in TL2 format. Empty means none, '*' means all")

	f.BoolVar(&opt.TL2MigrationDevMode, "tl2migrationDevMode", false,
		"during migration, do not overwrite existing files")
	f.BoolVar(&opt.NewBrackets, "newBrackets", false,
		"generate vectors/tuples in pure kernel")
}
