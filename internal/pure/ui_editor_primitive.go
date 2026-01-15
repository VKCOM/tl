// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package pure

import (
	"strings"
)

type UIEditorPrimitive struct {
	value    KernelValuePrimitive
	str      string
	strRunes []string
}

func (e *UIEditorPrimitive) SetValue(v KernelValuePrimitive) {
	e.value = v
	e.str = string(v.WriteJSON(nil, nil))
	e.strRunes = e.strRunes[:0]
	for _, r := range e.str {
		e.strRunes = append(e.strRunes, string(r))
	}
}

func (e *UIEditorPrimitive) UIWrite(sb *strings.Builder, model *UIModel) {
	if e.str == "0" {
		model.WriteCursor(sb, "0")
	} else {
		sb.WriteString(e.str)
		model.WriteCursor(sb, " ")
	}
}

func (e *UIEditorPrimitive) OnRune(msg string, model *UIModel) {
	if e.str == "0" {
		e.strRunes = e.strRunes[:0]
	}
	e.strRunes = append(e.strRunes, msg)
	e.str = strings.Join(e.strRunes, "")
}

func (e *UIEditorPrimitive) OnBackspace(model *UIModel) {
	if len(e.strRunes) == 0 {
		return
	}
	e.strRunes = e.strRunes[:len(e.strRunes)-1]
	if len(e.strRunes) == 0 {
		e.strRunes = append(e.strRunes, "0")
	}
	e.str = strings.Join(e.strRunes, "")
}

func (e *UIEditorPrimitive) OnEnter(model *UIModel) {
	if err := e.value.SetFromEditor(e.str); err != nil {
		model.LastError = err
		return
	}
	model.CurrentEditor = nil
}

func (e *UIEditorPrimitive) OnTab(model *UIModel, side int) {
	if err := e.value.SetFromEditor(e.str); err != nil {
		model.LastError = err
		return
	}
	model.CurrentEditor = nil
	model.Move(side)
	model.StartEdit(0)
}

func (e *UIEditorPrimitive) OnEscape(model *UIModel) {
	model.CurrentEditor = nil
}

func (e *UIEditorPrimitive) Value() KernelValue {
	return e.value
}
