// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
package tlTasksFullFilledCron

import (
	"github.com/vkcom/tl/internal/tlcodegen/test/gen/schema/internal"
	"github.com/vkcom/tl/pkg/basictl"
)

var _ = basictl.NatWrite
var _ = internal.ErrorInvalidEnumTag

type TasksFullFilledCron struct {
	FieldsMask uint32
	A0         int32 // Conditional: item.FieldsMask.0
	A1         int32 // Conditional: item.FieldsMask.1
	A2         int32 // Conditional: item.FieldsMask.2
	A3         int32 // Conditional: item.FieldsMask.3
	A4         int32 // Conditional: item.FieldsMask.4
	A5         int32 // Conditional: item.FieldsMask.5
	A6         int32 // Conditional: item.FieldsMask.6
	A7         int32 // Conditional: item.FieldsMask.7
	A8         int32 // Conditional: item.FieldsMask.8
	A9         int32 // Conditional: item.FieldsMask.9
	A10        int32 // Conditional: item.FieldsMask.10
	A11        int32 // Conditional: item.FieldsMask.11
	A12        int32 // Conditional: item.FieldsMask.12
	A13        int32 // Conditional: item.FieldsMask.13
	A14        int32 // Conditional: item.FieldsMask.14
	A15        int32 // Conditional: item.FieldsMask.15
	A16        int32 // Conditional: item.FieldsMask.16
	A17        int32 // Conditional: item.FieldsMask.17
	A18        int32 // Conditional: item.FieldsMask.18
	A19        int32 // Conditional: item.FieldsMask.19
	A20        int32 // Conditional: item.FieldsMask.20
	A21        int32 // Conditional: item.FieldsMask.21
	A22        int32 // Conditional: item.FieldsMask.22
	A23        int32 // Conditional: item.FieldsMask.23
	A24        int32 // Conditional: item.FieldsMask.24
	A25        int32 // Conditional: item.FieldsMask.25
	A26        int32 // Conditional: item.FieldsMask.26
	A27        int32 // Conditional: item.FieldsMask.27
	A28        int32 // Conditional: item.FieldsMask.28
	A29        int32 // Conditional: item.FieldsMask.29
	A30        int32 // Conditional: item.FieldsMask.30
	A31        int32 // Conditional: item.FieldsMask.31
}

func (TasksFullFilledCron) TLName() string { return "tasks.fullFilledCron" }
func (TasksFullFilledCron) TLTag() uint32  { return 0xd4177d7e }

func (item *TasksFullFilledCron) SetA0(v int32) {
	item.A0 = v
	item.FieldsMask |= 1 << 0
}
func (item *TasksFullFilledCron) ClearA0() {
	item.A0 = 0
	item.FieldsMask &^= 1 << 0
}
func (item *TasksFullFilledCron) IsSetA0() bool { return item.FieldsMask&(1<<0) != 0 }

func (item *TasksFullFilledCron) SetA1(v int32) {
	item.A1 = v
	item.FieldsMask |= 1 << 1
}
func (item *TasksFullFilledCron) ClearA1() {
	item.A1 = 0
	item.FieldsMask &^= 1 << 1
}
func (item *TasksFullFilledCron) IsSetA1() bool { return item.FieldsMask&(1<<1) != 0 }

func (item *TasksFullFilledCron) SetA2(v int32) {
	item.A2 = v
	item.FieldsMask |= 1 << 2
}
func (item *TasksFullFilledCron) ClearA2() {
	item.A2 = 0
	item.FieldsMask &^= 1 << 2
}
func (item *TasksFullFilledCron) IsSetA2() bool { return item.FieldsMask&(1<<2) != 0 }

func (item *TasksFullFilledCron) SetA3(v int32) {
	item.A3 = v
	item.FieldsMask |= 1 << 3
}
func (item *TasksFullFilledCron) ClearA3() {
	item.A3 = 0
	item.FieldsMask &^= 1 << 3
}
func (item *TasksFullFilledCron) IsSetA3() bool { return item.FieldsMask&(1<<3) != 0 }

func (item *TasksFullFilledCron) SetA4(v int32) {
	item.A4 = v
	item.FieldsMask |= 1 << 4
}
func (item *TasksFullFilledCron) ClearA4() {
	item.A4 = 0
	item.FieldsMask &^= 1 << 4
}
func (item *TasksFullFilledCron) IsSetA4() bool { return item.FieldsMask&(1<<4) != 0 }

func (item *TasksFullFilledCron) SetA5(v int32) {
	item.A5 = v
	item.FieldsMask |= 1 << 5
}
func (item *TasksFullFilledCron) ClearA5() {
	item.A5 = 0
	item.FieldsMask &^= 1 << 5
}
func (item *TasksFullFilledCron) IsSetA5() bool { return item.FieldsMask&(1<<5) != 0 }

func (item *TasksFullFilledCron) SetA6(v int32) {
	item.A6 = v
	item.FieldsMask |= 1 << 6
}
func (item *TasksFullFilledCron) ClearA6() {
	item.A6 = 0
	item.FieldsMask &^= 1 << 6
}
func (item *TasksFullFilledCron) IsSetA6() bool { return item.FieldsMask&(1<<6) != 0 }

func (item *TasksFullFilledCron) SetA7(v int32) {
	item.A7 = v
	item.FieldsMask |= 1 << 7
}
func (item *TasksFullFilledCron) ClearA7() {
	item.A7 = 0
	item.FieldsMask &^= 1 << 7
}
func (item *TasksFullFilledCron) IsSetA7() bool { return item.FieldsMask&(1<<7) != 0 }

func (item *TasksFullFilledCron) SetA8(v int32) {
	item.A8 = v
	item.FieldsMask |= 1 << 8
}
func (item *TasksFullFilledCron) ClearA8() {
	item.A8 = 0
	item.FieldsMask &^= 1 << 8
}
func (item *TasksFullFilledCron) IsSetA8() bool { return item.FieldsMask&(1<<8) != 0 }

func (item *TasksFullFilledCron) SetA9(v int32) {
	item.A9 = v
	item.FieldsMask |= 1 << 9
}
func (item *TasksFullFilledCron) ClearA9() {
	item.A9 = 0
	item.FieldsMask &^= 1 << 9
}
func (item *TasksFullFilledCron) IsSetA9() bool { return item.FieldsMask&(1<<9) != 0 }

func (item *TasksFullFilledCron) SetA10(v int32) {
	item.A10 = v
	item.FieldsMask |= 1 << 10
}
func (item *TasksFullFilledCron) ClearA10() {
	item.A10 = 0
	item.FieldsMask &^= 1 << 10
}
func (item *TasksFullFilledCron) IsSetA10() bool { return item.FieldsMask&(1<<10) != 0 }

func (item *TasksFullFilledCron) SetA11(v int32) {
	item.A11 = v
	item.FieldsMask |= 1 << 11
}
func (item *TasksFullFilledCron) ClearA11() {
	item.A11 = 0
	item.FieldsMask &^= 1 << 11
}
func (item *TasksFullFilledCron) IsSetA11() bool { return item.FieldsMask&(1<<11) != 0 }

func (item *TasksFullFilledCron) SetA12(v int32) {
	item.A12 = v
	item.FieldsMask |= 1 << 12
}
func (item *TasksFullFilledCron) ClearA12() {
	item.A12 = 0
	item.FieldsMask &^= 1 << 12
}
func (item *TasksFullFilledCron) IsSetA12() bool { return item.FieldsMask&(1<<12) != 0 }

func (item *TasksFullFilledCron) SetA13(v int32) {
	item.A13 = v
	item.FieldsMask |= 1 << 13
}
func (item *TasksFullFilledCron) ClearA13() {
	item.A13 = 0
	item.FieldsMask &^= 1 << 13
}
func (item *TasksFullFilledCron) IsSetA13() bool { return item.FieldsMask&(1<<13) != 0 }

func (item *TasksFullFilledCron) SetA14(v int32) {
	item.A14 = v
	item.FieldsMask |= 1 << 14
}
func (item *TasksFullFilledCron) ClearA14() {
	item.A14 = 0
	item.FieldsMask &^= 1 << 14
}
func (item *TasksFullFilledCron) IsSetA14() bool { return item.FieldsMask&(1<<14) != 0 }

func (item *TasksFullFilledCron) SetA15(v int32) {
	item.A15 = v
	item.FieldsMask |= 1 << 15
}
func (item *TasksFullFilledCron) ClearA15() {
	item.A15 = 0
	item.FieldsMask &^= 1 << 15
}
func (item *TasksFullFilledCron) IsSetA15() bool { return item.FieldsMask&(1<<15) != 0 }

func (item *TasksFullFilledCron) SetA16(v int32) {
	item.A16 = v
	item.FieldsMask |= 1 << 16
}
func (item *TasksFullFilledCron) ClearA16() {
	item.A16 = 0
	item.FieldsMask &^= 1 << 16
}
func (item *TasksFullFilledCron) IsSetA16() bool { return item.FieldsMask&(1<<16) != 0 }

func (item *TasksFullFilledCron) SetA17(v int32) {
	item.A17 = v
	item.FieldsMask |= 1 << 17
}
func (item *TasksFullFilledCron) ClearA17() {
	item.A17 = 0
	item.FieldsMask &^= 1 << 17
}
func (item *TasksFullFilledCron) IsSetA17() bool { return item.FieldsMask&(1<<17) != 0 }

func (item *TasksFullFilledCron) SetA18(v int32) {
	item.A18 = v
	item.FieldsMask |= 1 << 18
}
func (item *TasksFullFilledCron) ClearA18() {
	item.A18 = 0
	item.FieldsMask &^= 1 << 18
}
func (item *TasksFullFilledCron) IsSetA18() bool { return item.FieldsMask&(1<<18) != 0 }

func (item *TasksFullFilledCron) SetA19(v int32) {
	item.A19 = v
	item.FieldsMask |= 1 << 19
}
func (item *TasksFullFilledCron) ClearA19() {
	item.A19 = 0
	item.FieldsMask &^= 1 << 19
}
func (item *TasksFullFilledCron) IsSetA19() bool { return item.FieldsMask&(1<<19) != 0 }

func (item *TasksFullFilledCron) SetA20(v int32) {
	item.A20 = v
	item.FieldsMask |= 1 << 20
}
func (item *TasksFullFilledCron) ClearA20() {
	item.A20 = 0
	item.FieldsMask &^= 1 << 20
}
func (item *TasksFullFilledCron) IsSetA20() bool { return item.FieldsMask&(1<<20) != 0 }

func (item *TasksFullFilledCron) SetA21(v int32) {
	item.A21 = v
	item.FieldsMask |= 1 << 21
}
func (item *TasksFullFilledCron) ClearA21() {
	item.A21 = 0
	item.FieldsMask &^= 1 << 21
}
func (item *TasksFullFilledCron) IsSetA21() bool { return item.FieldsMask&(1<<21) != 0 }

func (item *TasksFullFilledCron) SetA22(v int32) {
	item.A22 = v
	item.FieldsMask |= 1 << 22
}
func (item *TasksFullFilledCron) ClearA22() {
	item.A22 = 0
	item.FieldsMask &^= 1 << 22
}
func (item *TasksFullFilledCron) IsSetA22() bool { return item.FieldsMask&(1<<22) != 0 }

func (item *TasksFullFilledCron) SetA23(v int32) {
	item.A23 = v
	item.FieldsMask |= 1 << 23
}
func (item *TasksFullFilledCron) ClearA23() {
	item.A23 = 0
	item.FieldsMask &^= 1 << 23
}
func (item *TasksFullFilledCron) IsSetA23() bool { return item.FieldsMask&(1<<23) != 0 }

func (item *TasksFullFilledCron) SetA24(v int32) {
	item.A24 = v
	item.FieldsMask |= 1 << 24
}
func (item *TasksFullFilledCron) ClearA24() {
	item.A24 = 0
	item.FieldsMask &^= 1 << 24
}
func (item *TasksFullFilledCron) IsSetA24() bool { return item.FieldsMask&(1<<24) != 0 }

func (item *TasksFullFilledCron) SetA25(v int32) {
	item.A25 = v
	item.FieldsMask |= 1 << 25
}
func (item *TasksFullFilledCron) ClearA25() {
	item.A25 = 0
	item.FieldsMask &^= 1 << 25
}
func (item *TasksFullFilledCron) IsSetA25() bool { return item.FieldsMask&(1<<25) != 0 }

func (item *TasksFullFilledCron) SetA26(v int32) {
	item.A26 = v
	item.FieldsMask |= 1 << 26
}
func (item *TasksFullFilledCron) ClearA26() {
	item.A26 = 0
	item.FieldsMask &^= 1 << 26
}
func (item *TasksFullFilledCron) IsSetA26() bool { return item.FieldsMask&(1<<26) != 0 }

func (item *TasksFullFilledCron) SetA27(v int32) {
	item.A27 = v
	item.FieldsMask |= 1 << 27
}
func (item *TasksFullFilledCron) ClearA27() {
	item.A27 = 0
	item.FieldsMask &^= 1 << 27
}
func (item *TasksFullFilledCron) IsSetA27() bool { return item.FieldsMask&(1<<27) != 0 }

func (item *TasksFullFilledCron) SetA28(v int32) {
	item.A28 = v
	item.FieldsMask |= 1 << 28
}
func (item *TasksFullFilledCron) ClearA28() {
	item.A28 = 0
	item.FieldsMask &^= 1 << 28
}
func (item *TasksFullFilledCron) IsSetA28() bool { return item.FieldsMask&(1<<28) != 0 }

func (item *TasksFullFilledCron) SetA29(v int32) {
	item.A29 = v
	item.FieldsMask |= 1 << 29
}
func (item *TasksFullFilledCron) ClearA29() {
	item.A29 = 0
	item.FieldsMask &^= 1 << 29
}
func (item *TasksFullFilledCron) IsSetA29() bool { return item.FieldsMask&(1<<29) != 0 }

func (item *TasksFullFilledCron) SetA30(v int32) {
	item.A30 = v
	item.FieldsMask |= 1 << 30
}
func (item *TasksFullFilledCron) ClearA30() {
	item.A30 = 0
	item.FieldsMask &^= 1 << 30
}
func (item *TasksFullFilledCron) IsSetA30() bool { return item.FieldsMask&(1<<30) != 0 }

func (item *TasksFullFilledCron) SetA31(v int32) {
	item.A31 = v
	item.FieldsMask |= 1 << 31
}
func (item *TasksFullFilledCron) ClearA31() {
	item.A31 = 0
	item.FieldsMask &^= 1 << 31
}
func (item *TasksFullFilledCron) IsSetA31() bool { return item.FieldsMask&(1<<31) != 0 }

func (item *TasksFullFilledCron) Reset() {
	item.FieldsMask = 0
	item.A0 = 0
	item.A1 = 0
	item.A2 = 0
	item.A3 = 0
	item.A4 = 0
	item.A5 = 0
	item.A6 = 0
	item.A7 = 0
	item.A8 = 0
	item.A9 = 0
	item.A10 = 0
	item.A11 = 0
	item.A12 = 0
	item.A13 = 0
	item.A14 = 0
	item.A15 = 0
	item.A16 = 0
	item.A17 = 0
	item.A18 = 0
	item.A19 = 0
	item.A20 = 0
	item.A21 = 0
	item.A22 = 0
	item.A23 = 0
	item.A24 = 0
	item.A25 = 0
	item.A26 = 0
	item.A27 = 0
	item.A28 = 0
	item.A29 = 0
	item.A30 = 0
	item.A31 = 0
}

func (item *TasksFullFilledCron) Read(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatRead(w, &item.FieldsMask); err != nil {
		return w, err
	}
	if item.FieldsMask&(1<<0) != 0 {
		if w, err = basictl.IntRead(w, &item.A0); err != nil {
			return w, err
		}
	} else {
		item.A0 = 0
	}
	if item.FieldsMask&(1<<1) != 0 {
		if w, err = basictl.IntRead(w, &item.A1); err != nil {
			return w, err
		}
	} else {
		item.A1 = 0
	}
	if item.FieldsMask&(1<<2) != 0 {
		if w, err = basictl.IntRead(w, &item.A2); err != nil {
			return w, err
		}
	} else {
		item.A2 = 0
	}
	if item.FieldsMask&(1<<3) != 0 {
		if w, err = basictl.IntRead(w, &item.A3); err != nil {
			return w, err
		}
	} else {
		item.A3 = 0
	}
	if item.FieldsMask&(1<<4) != 0 {
		if w, err = basictl.IntRead(w, &item.A4); err != nil {
			return w, err
		}
	} else {
		item.A4 = 0
	}
	if item.FieldsMask&(1<<5) != 0 {
		if w, err = basictl.IntRead(w, &item.A5); err != nil {
			return w, err
		}
	} else {
		item.A5 = 0
	}
	if item.FieldsMask&(1<<6) != 0 {
		if w, err = basictl.IntRead(w, &item.A6); err != nil {
			return w, err
		}
	} else {
		item.A6 = 0
	}
	if item.FieldsMask&(1<<7) != 0 {
		if w, err = basictl.IntRead(w, &item.A7); err != nil {
			return w, err
		}
	} else {
		item.A7 = 0
	}
	if item.FieldsMask&(1<<8) != 0 {
		if w, err = basictl.IntRead(w, &item.A8); err != nil {
			return w, err
		}
	} else {
		item.A8 = 0
	}
	if item.FieldsMask&(1<<9) != 0 {
		if w, err = basictl.IntRead(w, &item.A9); err != nil {
			return w, err
		}
	} else {
		item.A9 = 0
	}
	if item.FieldsMask&(1<<10) != 0 {
		if w, err = basictl.IntRead(w, &item.A10); err != nil {
			return w, err
		}
	} else {
		item.A10 = 0
	}
	if item.FieldsMask&(1<<11) != 0 {
		if w, err = basictl.IntRead(w, &item.A11); err != nil {
			return w, err
		}
	} else {
		item.A11 = 0
	}
	if item.FieldsMask&(1<<12) != 0 {
		if w, err = basictl.IntRead(w, &item.A12); err != nil {
			return w, err
		}
	} else {
		item.A12 = 0
	}
	if item.FieldsMask&(1<<13) != 0 {
		if w, err = basictl.IntRead(w, &item.A13); err != nil {
			return w, err
		}
	} else {
		item.A13 = 0
	}
	if item.FieldsMask&(1<<14) != 0 {
		if w, err = basictl.IntRead(w, &item.A14); err != nil {
			return w, err
		}
	} else {
		item.A14 = 0
	}
	if item.FieldsMask&(1<<15) != 0 {
		if w, err = basictl.IntRead(w, &item.A15); err != nil {
			return w, err
		}
	} else {
		item.A15 = 0
	}
	if item.FieldsMask&(1<<16) != 0 {
		if w, err = basictl.IntRead(w, &item.A16); err != nil {
			return w, err
		}
	} else {
		item.A16 = 0
	}
	if item.FieldsMask&(1<<17) != 0 {
		if w, err = basictl.IntRead(w, &item.A17); err != nil {
			return w, err
		}
	} else {
		item.A17 = 0
	}
	if item.FieldsMask&(1<<18) != 0 {
		if w, err = basictl.IntRead(w, &item.A18); err != nil {
			return w, err
		}
	} else {
		item.A18 = 0
	}
	if item.FieldsMask&(1<<19) != 0 {
		if w, err = basictl.IntRead(w, &item.A19); err != nil {
			return w, err
		}
	} else {
		item.A19 = 0
	}
	if item.FieldsMask&(1<<20) != 0 {
		if w, err = basictl.IntRead(w, &item.A20); err != nil {
			return w, err
		}
	} else {
		item.A20 = 0
	}
	if item.FieldsMask&(1<<21) != 0 {
		if w, err = basictl.IntRead(w, &item.A21); err != nil {
			return w, err
		}
	} else {
		item.A21 = 0
	}
	if item.FieldsMask&(1<<22) != 0 {
		if w, err = basictl.IntRead(w, &item.A22); err != nil {
			return w, err
		}
	} else {
		item.A22 = 0
	}
	if item.FieldsMask&(1<<23) != 0 {
		if w, err = basictl.IntRead(w, &item.A23); err != nil {
			return w, err
		}
	} else {
		item.A23 = 0
	}
	if item.FieldsMask&(1<<24) != 0 {
		if w, err = basictl.IntRead(w, &item.A24); err != nil {
			return w, err
		}
	} else {
		item.A24 = 0
	}
	if item.FieldsMask&(1<<25) != 0 {
		if w, err = basictl.IntRead(w, &item.A25); err != nil {
			return w, err
		}
	} else {
		item.A25 = 0
	}
	if item.FieldsMask&(1<<26) != 0 {
		if w, err = basictl.IntRead(w, &item.A26); err != nil {
			return w, err
		}
	} else {
		item.A26 = 0
	}
	if item.FieldsMask&(1<<27) != 0 {
		if w, err = basictl.IntRead(w, &item.A27); err != nil {
			return w, err
		}
	} else {
		item.A27 = 0
	}
	if item.FieldsMask&(1<<28) != 0 {
		if w, err = basictl.IntRead(w, &item.A28); err != nil {
			return w, err
		}
	} else {
		item.A28 = 0
	}
	if item.FieldsMask&(1<<29) != 0 {
		if w, err = basictl.IntRead(w, &item.A29); err != nil {
			return w, err
		}
	} else {
		item.A29 = 0
	}
	if item.FieldsMask&(1<<30) != 0 {
		if w, err = basictl.IntRead(w, &item.A30); err != nil {
			return w, err
		}
	} else {
		item.A30 = 0
	}
	if item.FieldsMask&(1<<31) != 0 {
		if w, err = basictl.IntRead(w, &item.A31); err != nil {
			return w, err
		}
	} else {
		item.A31 = 0
	}
	return w, nil
}

func (item *TasksFullFilledCron) WriteGeneral(w []byte) (_ []byte, err error) {
	return item.Write(w), nil
}

func (item *TasksFullFilledCron) Write(w []byte) []byte {
	w = basictl.NatWrite(w, item.FieldsMask)
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.IntWrite(w, item.A0)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.IntWrite(w, item.A1)
	}
	if item.FieldsMask&(1<<2) != 0 {
		w = basictl.IntWrite(w, item.A2)
	}
	if item.FieldsMask&(1<<3) != 0 {
		w = basictl.IntWrite(w, item.A3)
	}
	if item.FieldsMask&(1<<4) != 0 {
		w = basictl.IntWrite(w, item.A4)
	}
	if item.FieldsMask&(1<<5) != 0 {
		w = basictl.IntWrite(w, item.A5)
	}
	if item.FieldsMask&(1<<6) != 0 {
		w = basictl.IntWrite(w, item.A6)
	}
	if item.FieldsMask&(1<<7) != 0 {
		w = basictl.IntWrite(w, item.A7)
	}
	if item.FieldsMask&(1<<8) != 0 {
		w = basictl.IntWrite(w, item.A8)
	}
	if item.FieldsMask&(1<<9) != 0 {
		w = basictl.IntWrite(w, item.A9)
	}
	if item.FieldsMask&(1<<10) != 0 {
		w = basictl.IntWrite(w, item.A10)
	}
	if item.FieldsMask&(1<<11) != 0 {
		w = basictl.IntWrite(w, item.A11)
	}
	if item.FieldsMask&(1<<12) != 0 {
		w = basictl.IntWrite(w, item.A12)
	}
	if item.FieldsMask&(1<<13) != 0 {
		w = basictl.IntWrite(w, item.A13)
	}
	if item.FieldsMask&(1<<14) != 0 {
		w = basictl.IntWrite(w, item.A14)
	}
	if item.FieldsMask&(1<<15) != 0 {
		w = basictl.IntWrite(w, item.A15)
	}
	if item.FieldsMask&(1<<16) != 0 {
		w = basictl.IntWrite(w, item.A16)
	}
	if item.FieldsMask&(1<<17) != 0 {
		w = basictl.IntWrite(w, item.A17)
	}
	if item.FieldsMask&(1<<18) != 0 {
		w = basictl.IntWrite(w, item.A18)
	}
	if item.FieldsMask&(1<<19) != 0 {
		w = basictl.IntWrite(w, item.A19)
	}
	if item.FieldsMask&(1<<20) != 0 {
		w = basictl.IntWrite(w, item.A20)
	}
	if item.FieldsMask&(1<<21) != 0 {
		w = basictl.IntWrite(w, item.A21)
	}
	if item.FieldsMask&(1<<22) != 0 {
		w = basictl.IntWrite(w, item.A22)
	}
	if item.FieldsMask&(1<<23) != 0 {
		w = basictl.IntWrite(w, item.A23)
	}
	if item.FieldsMask&(1<<24) != 0 {
		w = basictl.IntWrite(w, item.A24)
	}
	if item.FieldsMask&(1<<25) != 0 {
		w = basictl.IntWrite(w, item.A25)
	}
	if item.FieldsMask&(1<<26) != 0 {
		w = basictl.IntWrite(w, item.A26)
	}
	if item.FieldsMask&(1<<27) != 0 {
		w = basictl.IntWrite(w, item.A27)
	}
	if item.FieldsMask&(1<<28) != 0 {
		w = basictl.IntWrite(w, item.A28)
	}
	if item.FieldsMask&(1<<29) != 0 {
		w = basictl.IntWrite(w, item.A29)
	}
	if item.FieldsMask&(1<<30) != 0 {
		w = basictl.IntWrite(w, item.A30)
	}
	if item.FieldsMask&(1<<31) != 0 {
		w = basictl.IntWrite(w, item.A31)
	}
	return w
}

func (item *TasksFullFilledCron) ReadBoxed(w []byte) (_ []byte, err error) {
	if w, err = basictl.NatReadExactTag(w, 0xd4177d7e); err != nil {
		return w, err
	}
	return item.Read(w)
}

func (item *TasksFullFilledCron) WriteBoxedGeneral(w []byte) (_ []byte, err error) {
	return item.WriteBoxed(w), nil
}

func (item *TasksFullFilledCron) WriteBoxed(w []byte) []byte {
	w = basictl.NatWrite(w, 0xd4177d7e)
	return item.Write(w)
}

func (item TasksFullFilledCron) String() string {
	return string(item.WriteJSON(nil))
}

func (item *TasksFullFilledCron) ReadJSON(legacyTypeNames bool, in *basictl.JsonLexer) error {
	var propFieldsMaskPresented bool
	var propA0Presented bool
	var propA1Presented bool
	var propA2Presented bool
	var propA3Presented bool
	var propA4Presented bool
	var propA5Presented bool
	var propA6Presented bool
	var propA7Presented bool
	var propA8Presented bool
	var propA9Presented bool
	var propA10Presented bool
	var propA11Presented bool
	var propA12Presented bool
	var propA13Presented bool
	var propA14Presented bool
	var propA15Presented bool
	var propA16Presented bool
	var propA17Presented bool
	var propA18Presented bool
	var propA19Presented bool
	var propA20Presented bool
	var propA21Presented bool
	var propA22Presented bool
	var propA23Presented bool
	var propA24Presented bool
	var propA25Presented bool
	var propA26Presented bool
	var propA27Presented bool
	var propA28Presented bool
	var propA29Presented bool
	var propA30Presented bool
	var propA31Presented bool

	if in != nil {
		in.Delim('{')
		if !in.Ok() {
			return in.Error()
		}
		for !in.IsDelim('}') {
			key := in.UnsafeFieldName(true)
			in.WantColon()
			switch key {
			case "fields_mask":
				if propFieldsMaskPresented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "fields_mask")
				}
				if err := internal.Json2ReadUint32(in, &item.FieldsMask); err != nil {
					return err
				}
				propFieldsMaskPresented = true
			case "a0":
				if propA0Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a0")
				}
				if err := internal.Json2ReadInt32(in, &item.A0); err != nil {
					return err
				}
				propA0Presented = true
			case "a1":
				if propA1Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a1")
				}
				if err := internal.Json2ReadInt32(in, &item.A1); err != nil {
					return err
				}
				propA1Presented = true
			case "a2":
				if propA2Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a2")
				}
				if err := internal.Json2ReadInt32(in, &item.A2); err != nil {
					return err
				}
				propA2Presented = true
			case "a3":
				if propA3Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a3")
				}
				if err := internal.Json2ReadInt32(in, &item.A3); err != nil {
					return err
				}
				propA3Presented = true
			case "a4":
				if propA4Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a4")
				}
				if err := internal.Json2ReadInt32(in, &item.A4); err != nil {
					return err
				}
				propA4Presented = true
			case "a5":
				if propA5Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a5")
				}
				if err := internal.Json2ReadInt32(in, &item.A5); err != nil {
					return err
				}
				propA5Presented = true
			case "a6":
				if propA6Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a6")
				}
				if err := internal.Json2ReadInt32(in, &item.A6); err != nil {
					return err
				}
				propA6Presented = true
			case "a7":
				if propA7Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a7")
				}
				if err := internal.Json2ReadInt32(in, &item.A7); err != nil {
					return err
				}
				propA7Presented = true
			case "a8":
				if propA8Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a8")
				}
				if err := internal.Json2ReadInt32(in, &item.A8); err != nil {
					return err
				}
				propA8Presented = true
			case "a9":
				if propA9Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a9")
				}
				if err := internal.Json2ReadInt32(in, &item.A9); err != nil {
					return err
				}
				propA9Presented = true
			case "a10":
				if propA10Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a10")
				}
				if err := internal.Json2ReadInt32(in, &item.A10); err != nil {
					return err
				}
				propA10Presented = true
			case "a11":
				if propA11Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a11")
				}
				if err := internal.Json2ReadInt32(in, &item.A11); err != nil {
					return err
				}
				propA11Presented = true
			case "a12":
				if propA12Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a12")
				}
				if err := internal.Json2ReadInt32(in, &item.A12); err != nil {
					return err
				}
				propA12Presented = true
			case "a13":
				if propA13Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a13")
				}
				if err := internal.Json2ReadInt32(in, &item.A13); err != nil {
					return err
				}
				propA13Presented = true
			case "a14":
				if propA14Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a14")
				}
				if err := internal.Json2ReadInt32(in, &item.A14); err != nil {
					return err
				}
				propA14Presented = true
			case "a15":
				if propA15Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a15")
				}
				if err := internal.Json2ReadInt32(in, &item.A15); err != nil {
					return err
				}
				propA15Presented = true
			case "a16":
				if propA16Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a16")
				}
				if err := internal.Json2ReadInt32(in, &item.A16); err != nil {
					return err
				}
				propA16Presented = true
			case "a17":
				if propA17Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a17")
				}
				if err := internal.Json2ReadInt32(in, &item.A17); err != nil {
					return err
				}
				propA17Presented = true
			case "a18":
				if propA18Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a18")
				}
				if err := internal.Json2ReadInt32(in, &item.A18); err != nil {
					return err
				}
				propA18Presented = true
			case "a19":
				if propA19Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a19")
				}
				if err := internal.Json2ReadInt32(in, &item.A19); err != nil {
					return err
				}
				propA19Presented = true
			case "a20":
				if propA20Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a20")
				}
				if err := internal.Json2ReadInt32(in, &item.A20); err != nil {
					return err
				}
				propA20Presented = true
			case "a21":
				if propA21Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a21")
				}
				if err := internal.Json2ReadInt32(in, &item.A21); err != nil {
					return err
				}
				propA21Presented = true
			case "a22":
				if propA22Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a22")
				}
				if err := internal.Json2ReadInt32(in, &item.A22); err != nil {
					return err
				}
				propA22Presented = true
			case "a23":
				if propA23Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a23")
				}
				if err := internal.Json2ReadInt32(in, &item.A23); err != nil {
					return err
				}
				propA23Presented = true
			case "a24":
				if propA24Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a24")
				}
				if err := internal.Json2ReadInt32(in, &item.A24); err != nil {
					return err
				}
				propA24Presented = true
			case "a25":
				if propA25Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a25")
				}
				if err := internal.Json2ReadInt32(in, &item.A25); err != nil {
					return err
				}
				propA25Presented = true
			case "a26":
				if propA26Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a26")
				}
				if err := internal.Json2ReadInt32(in, &item.A26); err != nil {
					return err
				}
				propA26Presented = true
			case "a27":
				if propA27Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a27")
				}
				if err := internal.Json2ReadInt32(in, &item.A27); err != nil {
					return err
				}
				propA27Presented = true
			case "a28":
				if propA28Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a28")
				}
				if err := internal.Json2ReadInt32(in, &item.A28); err != nil {
					return err
				}
				propA28Presented = true
			case "a29":
				if propA29Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a29")
				}
				if err := internal.Json2ReadInt32(in, &item.A29); err != nil {
					return err
				}
				propA29Presented = true
			case "a30":
				if propA30Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a30")
				}
				if err := internal.Json2ReadInt32(in, &item.A30); err != nil {
					return err
				}
				propA30Presented = true
			case "a31":
				if propA31Presented {
					return internal.ErrorInvalidJSONWithDuplicatingKeys("tasks.fullFilledCron", "a31")
				}
				if err := internal.Json2ReadInt32(in, &item.A31); err != nil {
					return err
				}
				propA31Presented = true
			default:
				return internal.ErrorInvalidJSONExcessElement("tasks.fullFilledCron", key)
			}
			in.WantComma()
		}
		in.Delim('}')
		if !in.Ok() {
			return in.Error()
		}
	}
	if !propFieldsMaskPresented {
		item.FieldsMask = 0
	}
	if !propA0Presented {
		item.A0 = 0
	}
	if !propA1Presented {
		item.A1 = 0
	}
	if !propA2Presented {
		item.A2 = 0
	}
	if !propA3Presented {
		item.A3 = 0
	}
	if !propA4Presented {
		item.A4 = 0
	}
	if !propA5Presented {
		item.A5 = 0
	}
	if !propA6Presented {
		item.A6 = 0
	}
	if !propA7Presented {
		item.A7 = 0
	}
	if !propA8Presented {
		item.A8 = 0
	}
	if !propA9Presented {
		item.A9 = 0
	}
	if !propA10Presented {
		item.A10 = 0
	}
	if !propA11Presented {
		item.A11 = 0
	}
	if !propA12Presented {
		item.A12 = 0
	}
	if !propA13Presented {
		item.A13 = 0
	}
	if !propA14Presented {
		item.A14 = 0
	}
	if !propA15Presented {
		item.A15 = 0
	}
	if !propA16Presented {
		item.A16 = 0
	}
	if !propA17Presented {
		item.A17 = 0
	}
	if !propA18Presented {
		item.A18 = 0
	}
	if !propA19Presented {
		item.A19 = 0
	}
	if !propA20Presented {
		item.A20 = 0
	}
	if !propA21Presented {
		item.A21 = 0
	}
	if !propA22Presented {
		item.A22 = 0
	}
	if !propA23Presented {
		item.A23 = 0
	}
	if !propA24Presented {
		item.A24 = 0
	}
	if !propA25Presented {
		item.A25 = 0
	}
	if !propA26Presented {
		item.A26 = 0
	}
	if !propA27Presented {
		item.A27 = 0
	}
	if !propA28Presented {
		item.A28 = 0
	}
	if !propA29Presented {
		item.A29 = 0
	}
	if !propA30Presented {
		item.A30 = 0
	}
	if !propA31Presented {
		item.A31 = 0
	}
	if propA0Presented {
		item.FieldsMask |= 1 << 0
	}
	if propA1Presented {
		item.FieldsMask |= 1 << 1
	}
	if propA2Presented {
		item.FieldsMask |= 1 << 2
	}
	if propA3Presented {
		item.FieldsMask |= 1 << 3
	}
	if propA4Presented {
		item.FieldsMask |= 1 << 4
	}
	if propA5Presented {
		item.FieldsMask |= 1 << 5
	}
	if propA6Presented {
		item.FieldsMask |= 1 << 6
	}
	if propA7Presented {
		item.FieldsMask |= 1 << 7
	}
	if propA8Presented {
		item.FieldsMask |= 1 << 8
	}
	if propA9Presented {
		item.FieldsMask |= 1 << 9
	}
	if propA10Presented {
		item.FieldsMask |= 1 << 10
	}
	if propA11Presented {
		item.FieldsMask |= 1 << 11
	}
	if propA12Presented {
		item.FieldsMask |= 1 << 12
	}
	if propA13Presented {
		item.FieldsMask |= 1 << 13
	}
	if propA14Presented {
		item.FieldsMask |= 1 << 14
	}
	if propA15Presented {
		item.FieldsMask |= 1 << 15
	}
	if propA16Presented {
		item.FieldsMask |= 1 << 16
	}
	if propA17Presented {
		item.FieldsMask |= 1 << 17
	}
	if propA18Presented {
		item.FieldsMask |= 1 << 18
	}
	if propA19Presented {
		item.FieldsMask |= 1 << 19
	}
	if propA20Presented {
		item.FieldsMask |= 1 << 20
	}
	if propA21Presented {
		item.FieldsMask |= 1 << 21
	}
	if propA22Presented {
		item.FieldsMask |= 1 << 22
	}
	if propA23Presented {
		item.FieldsMask |= 1 << 23
	}
	if propA24Presented {
		item.FieldsMask |= 1 << 24
	}
	if propA25Presented {
		item.FieldsMask |= 1 << 25
	}
	if propA26Presented {
		item.FieldsMask |= 1 << 26
	}
	if propA27Presented {
		item.FieldsMask |= 1 << 27
	}
	if propA28Presented {
		item.FieldsMask |= 1 << 28
	}
	if propA29Presented {
		item.FieldsMask |= 1 << 29
	}
	if propA30Presented {
		item.FieldsMask |= 1 << 30
	}
	if propA31Presented {
		item.FieldsMask |= 1 << 31
	}
	return nil
}

// This method is general version of WriteJSON, use it instead!
func (item *TasksFullFilledCron) WriteJSONGeneral(tctx *basictl.JSONWriteContext, w []byte) (_ []byte, err error) {
	return item.WriteJSONOpt(tctx, w), nil
}

func (item *TasksFullFilledCron) WriteJSON(w []byte) []byte {
	tctx := basictl.JSONWriteContext{}
	return item.WriteJSONOpt(&tctx, w)
}
func (item *TasksFullFilledCron) WriteJSONOpt(tctx *basictl.JSONWriteContext, w []byte) []byte {
	w = append(w, '{')
	backupIndexFieldsMask := len(w)
	w = basictl.JSONAddCommaIfNeeded(w)
	w = append(w, `"fields_mask":`...)
	w = basictl.JSONWriteUint32(w, item.FieldsMask)
	if (item.FieldsMask != 0) == false {
		w = w[:backupIndexFieldsMask]
	}
	if item.FieldsMask&(1<<0) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a0":`...)
		w = basictl.JSONWriteInt32(w, item.A0)
	}
	if item.FieldsMask&(1<<1) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a1":`...)
		w = basictl.JSONWriteInt32(w, item.A1)
	}
	if item.FieldsMask&(1<<2) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a2":`...)
		w = basictl.JSONWriteInt32(w, item.A2)
	}
	if item.FieldsMask&(1<<3) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a3":`...)
		w = basictl.JSONWriteInt32(w, item.A3)
	}
	if item.FieldsMask&(1<<4) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a4":`...)
		w = basictl.JSONWriteInt32(w, item.A4)
	}
	if item.FieldsMask&(1<<5) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a5":`...)
		w = basictl.JSONWriteInt32(w, item.A5)
	}
	if item.FieldsMask&(1<<6) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a6":`...)
		w = basictl.JSONWriteInt32(w, item.A6)
	}
	if item.FieldsMask&(1<<7) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a7":`...)
		w = basictl.JSONWriteInt32(w, item.A7)
	}
	if item.FieldsMask&(1<<8) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a8":`...)
		w = basictl.JSONWriteInt32(w, item.A8)
	}
	if item.FieldsMask&(1<<9) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a9":`...)
		w = basictl.JSONWriteInt32(w, item.A9)
	}
	if item.FieldsMask&(1<<10) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a10":`...)
		w = basictl.JSONWriteInt32(w, item.A10)
	}
	if item.FieldsMask&(1<<11) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a11":`...)
		w = basictl.JSONWriteInt32(w, item.A11)
	}
	if item.FieldsMask&(1<<12) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a12":`...)
		w = basictl.JSONWriteInt32(w, item.A12)
	}
	if item.FieldsMask&(1<<13) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a13":`...)
		w = basictl.JSONWriteInt32(w, item.A13)
	}
	if item.FieldsMask&(1<<14) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a14":`...)
		w = basictl.JSONWriteInt32(w, item.A14)
	}
	if item.FieldsMask&(1<<15) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a15":`...)
		w = basictl.JSONWriteInt32(w, item.A15)
	}
	if item.FieldsMask&(1<<16) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a16":`...)
		w = basictl.JSONWriteInt32(w, item.A16)
	}
	if item.FieldsMask&(1<<17) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a17":`...)
		w = basictl.JSONWriteInt32(w, item.A17)
	}
	if item.FieldsMask&(1<<18) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a18":`...)
		w = basictl.JSONWriteInt32(w, item.A18)
	}
	if item.FieldsMask&(1<<19) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a19":`...)
		w = basictl.JSONWriteInt32(w, item.A19)
	}
	if item.FieldsMask&(1<<20) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a20":`...)
		w = basictl.JSONWriteInt32(w, item.A20)
	}
	if item.FieldsMask&(1<<21) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a21":`...)
		w = basictl.JSONWriteInt32(w, item.A21)
	}
	if item.FieldsMask&(1<<22) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a22":`...)
		w = basictl.JSONWriteInt32(w, item.A22)
	}
	if item.FieldsMask&(1<<23) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a23":`...)
		w = basictl.JSONWriteInt32(w, item.A23)
	}
	if item.FieldsMask&(1<<24) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a24":`...)
		w = basictl.JSONWriteInt32(w, item.A24)
	}
	if item.FieldsMask&(1<<25) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a25":`...)
		w = basictl.JSONWriteInt32(w, item.A25)
	}
	if item.FieldsMask&(1<<26) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a26":`...)
		w = basictl.JSONWriteInt32(w, item.A26)
	}
	if item.FieldsMask&(1<<27) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a27":`...)
		w = basictl.JSONWriteInt32(w, item.A27)
	}
	if item.FieldsMask&(1<<28) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a28":`...)
		w = basictl.JSONWriteInt32(w, item.A28)
	}
	if item.FieldsMask&(1<<29) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a29":`...)
		w = basictl.JSONWriteInt32(w, item.A29)
	}
	if item.FieldsMask&(1<<30) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a30":`...)
		w = basictl.JSONWriteInt32(w, item.A30)
	}
	if item.FieldsMask&(1<<31) != 0 {
		w = basictl.JSONAddCommaIfNeeded(w)
		w = append(w, `"a31":`...)
		w = basictl.JSONWriteInt32(w, item.A31)
	}
	return append(w, '}')
}

func (item *TasksFullFilledCron) MarshalJSON() ([]byte, error) {
	return item.WriteJSON(nil), nil
}

func (item *TasksFullFilledCron) UnmarshalJSON(b []byte) error {
	if err := item.ReadJSON(true, &basictl.JsonLexer{Data: b}); err != nil {
		return internal.ErrorInvalidJSON("tasks.fullFilledCron", err.Error())
	}
	return nil
}
