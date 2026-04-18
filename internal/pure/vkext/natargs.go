// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package vkext

import (
	"github.com/VKCOM/tl/internal/pure"
)

//
//type natArgs struct {
//	stack  []uint32
//	offset int // my arguments start at offset
//}
//
//func (a *natArgs) MyArgs() []uint32 {
//	return a.stack[a.offset:]
//}
//
//func (a *natArgs) Pop() natArgs {
//	return natArgs{stack: a.stack[:a.offset], offset: a.offset}
//}

func formatNatArg(myNatArgs []uint32, arg pure.ActualNatArg) uint32 {
	if arg.IsNumber() {
		return arg.Number()
	}
	if arg.IsField() {
		panic("general formatNatArg cannot reference fields")
	}
	return myNatArgs[arg.FieldIndex()]
}

func formatNatArgs(natArgsStack []uint32, myNatArgs []uint32, natArgs []pure.ActualNatArg) []uint32 {
	for _, arg := range natArgs {
		natArgsStack = append(natArgsStack, formatNatArg(myNatArgs, arg))
	}
	return natArgsStack
}
