// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package build

import (
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

var (
	// Build* заполняются при сборке go build -ldflags
	buildTimestamp  string
	machine         string
	commit          string
	commitTimestamp string
	version         string
	number          string
	branchName      string
	name            string

	info string // combination of above

	appName               string
	commitTimestampUint32 uint32

	buildTimeFormatted   string
	buildTimestampUint32 uint32
)

func Time() string {
	return buildTimeFormatted
}

func Timestamp() uint32 {
	return buildTimestampUint32
}

func Machine() string {
	return machine
}

func Commit() string {
	return commit
}

// UNIX timestampt seconds, so stable in any TZ
func CommitTimestamp() uint32 {
	return commitTimestampUint32
}

func Version() string {
	return version
}

func Number() string {
	return number
}

func Name() string {
	return name
}

func BranchName() string {
	return branchName
}

func Info() string {
	return info
}

func init() {
	appName = path.Base(os.Args[0])
	ts, _ := strconv.ParseUint(commitTimestamp, 10, 32)
	commitTimestampUint32 = uint32(ts)

	if buildTimestamp != "" {
		ts, _ = strconv.ParseUint(buildTimestamp, 10, 32)
		buildTimestampUint32 = uint32(ts)
		buildTimeFormatted = time.Unix(int64(ts), 0).Format("2006-01-02T15:04:05-0700")
	}

	info = fmt.Sprintf("%s compiled at %s by %s after %s on %s build %s", appName, buildTimeFormatted, runtime.Version(), version, machine, number)
}

func AppName() string { // TODO - remember during build?
	return appName
}

func FlagParseShowVersionHelpWithTail(set *flag.FlagSet, args []string) {
	help := false
	version := false
	set.BoolVar(&help, `h`, false, `show this help`)
	set.BoolVar(&help, `help`, false, `show this help`)
	set.BoolVar(&version, `v`, false, `show version`)
	set.BoolVar(&version, `version`, false, `show version`)

	err := set.Parse(args)
	if err != nil {
		os.Exit(2) // enforce ExitOnError policy
	}

	if version {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", Info())
		os.Exit(0)
	}
	if help {
		if set.Usage != nil {
			set.Usage()
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", set.Name())
			set.PrintDefaults()
		}
		os.Exit(0)
	}
}

func FlagSetParseShowVersionHelp(set *flag.FlagSet, args []string) {
	FlagParseShowVersionHelpWithTail(set, args)
	if len(set.Args()) != 0 {
		_, _ = fmt.Fprintf(os.Stderr, "Unexpected command line argument - %q, check command line for typos\n", set.Args()[0])
		os.Exit(1)
	}
}

// Fatals if additional parameters passed. Protection against 'kittenhosue ch-addr=x -c=y' when dash is forgotten
func FlagParseShowVersionHelp() {
	FlagSetParseShowVersionHelp(flag.CommandLine, os.Args[1:])
}
