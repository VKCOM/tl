package tl2pure

import (
	"strings"

	"github.com/TwiN/go-color"
)

type UIEditorString struct {
	value    *KernelValueString
	str      string
	strRunes []string
}

func (e *UIEditorString) SetValue(v *KernelValueString) {
	e.value = v
	e.str = v.value
	e.strRunes = e.strRunes[:0]
	for _, r := range e.str {
		e.strRunes = append(e.strRunes, string(r))
	}
}

func (e *UIEditorString) UIWrite(sb *strings.Builder) {
	sb.WriteString(color.InBlackOverBlue(`"`))
	sb.WriteString(color.InBlackOverBlue(e.str))
	sb.WriteString(`"`)
}

func (e *UIEditorString) OnRune(msg string, model *UIModel) {
	e.strRunes = append(e.strRunes, msg)
	e.str = strings.Join(e.strRunes, "")
}

func (e *UIEditorString) OnBackspace(model *UIModel) {
	if len(e.strRunes) == 0 {
		return
	}
	e.strRunes = e.strRunes[:len(e.strRunes)-1]
	e.str = strings.Join(e.strRunes, "")
}

func (e *UIEditorString) OnEnter(model *UIModel) {
	e.FinishOK()
	model.CurrentEditor = nil
}

func (e *UIEditorString) OnTab(model *UIModel, side int) {
	e.FinishOK()
	model.CurrentEditor = nil
	model.Move(side)
	model.StartEdit(0)
}

func (e *UIEditorString) OnEscape(model *UIModel) {
	model.CurrentEditor = nil
}

func (e *UIEditorString) FinishOK() {
	e.value.value = e.str
}

func (e *UIEditorString) Value() KernelValue {
	return e.value
}
