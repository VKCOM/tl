package tl2pure

import (
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
)

type UIEditorUnion struct {
	value       *KernelValueUnion
	index       int
	prefix      string   // lowercase
	prefixRunes []string // lowercase
	names       []string
	lowerNames  []string
}

func (e *UIEditorUnion) SetValue(v *KernelValueUnion) {
	e.value = v
	e.index = v.index
	e.prefix = ""
	e.prefixRunes = e.prefixRunes[:0]
	e.names = e.names[:0]
	e.lowerNames = e.lowerNames[:0]
	for _, va := range v.instance.def.Variants {
		e.names = append(e.names, va.Name)
		e.lowerNames = append(e.lowerNames, strings.ToLower(va.Name))
	}
}

func (e *UIEditorUnion) UIWrite(sb *strings.Builder) {
	sb.WriteString(color.InBlackOverBlue(`"`))
	name := e.names[e.index]
	sb.WriteString(color.InBlackOverBlue(name[:len(e.prefix)]))
	sb.WriteString(name[len(e.prefix):])
	sb.WriteString(`"`)
}

func (e *UIEditorUnion) OnRune(msg string, model *UIModel) {
	newPrefixRunes := append(e.prefixRunes, strings.ToLower(msg))
	newPrefix := strings.Join(newPrefixRunes, "")
	if strings.HasPrefix(e.lowerNames[e.index], newPrefix) {
		e.prefixRunes = newPrefixRunes
		e.prefix = newPrefix
		return
	}
	for i, lowerName := range e.lowerNames {
		if strings.HasPrefix(lowerName, newPrefix) {
			e.index = i
			e.prefixRunes = newPrefixRunes
			e.prefix = newPrefix
			return
		}
	}
	model.LastError = fmt.Errorf("no variant with prefix %s", newPrefix)
}

func (e *UIEditorUnion) OnBackspace(model *UIModel) {
	if len(e.prefixRunes) == 0 {
		return
	}
	e.prefixRunes = e.prefixRunes[:len(e.prefixRunes)-1]
	e.prefix = strings.Join(e.prefixRunes, "")
}

func (e *UIEditorUnion) OnEnter(model *UIModel) {
	e.FinishOK()
	model.CurrentEditor = nil
}

func (e *UIEditorUnion) OnTab(model *UIModel) {
	e.FinishOK()
	model.CurrentEditor = nil
	model.Right()
	model.StartEdit()
}

func (e *UIEditorUnion) OnEscape(model *UIModel) {
	model.CurrentEditor = nil
}

func (e *UIEditorUnion) FinishOK() {
	e.value.index = e.index
}

func (e *UIEditorUnion) Value() KernelValue {
	return e.value
}
