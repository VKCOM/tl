package tl2pure

import (
	"strings"

	"github.com/TwiN/go-color"
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

func (e *UIEditorPrimitive) UIWrite(sb *strings.Builder) {
	sb.WriteString(color.InBlackOverBlue(e.str))
}

func (e *UIEditorPrimitive) OnRune(msg string, model *UIModel) {
	if e.str == "0" {
		e.strRunes = e.strRunes[:0]
		e.str = ""
	}
	e.strRunes = append(e.strRunes, msg)
	e.str = strings.Join(e.strRunes, "")
}

func (e *UIEditorPrimitive) OnBackspace(model *UIModel) {
	if len(e.strRunes) == 0 {
		return
	}
	e.strRunes = e.strRunes[:len(e.strRunes)-1]
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
	model.StartEdit(true)
}

func (e *UIEditorPrimitive) OnEscape(model *UIModel) {
	model.CurrentEditor = nil
}

func (e *UIEditorPrimitive) Value() KernelValue {
	return e.value
}
