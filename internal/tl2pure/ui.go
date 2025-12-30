package tl2pure

import (
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
)

type UIEditor interface {
	Value() KernelValue

	UIWrite(sb *strings.Builder)

	OnRune(msg string, model *UIModel)
	OnBackspace(model *UIModel)
	OnEnter(model *UIModel)
	OnTab(model *UIModel, side int) // -1 Shift-Tab
	OnEscape(model *UIModel)
}

type UIModel struct {
	Fun KernelValueObject

	Path []int // selected field on every hierarchy level, 0 means union constructor

	CurrentEditor UIEditor

	LastError error

	ShowLegend bool

	EditorUnion     UIEditorUnion
	EditorString    UIEditorString
	EditorPrimitive UIEditorPrimitive

	Width  int
	Height int
}

func (m *UIModel) colorButton(str string) string {
	return color.InCyan(str)
}

func (m *UIModel) colorButtonComment(str string) string {
	return str
}

func (m *UIModel) colorError(str string) string {
	return color.InRed(str)
}

func (m *UIModel) View() string {
	var sb strings.Builder
	sb.WriteString(m.colorButton("[F1]"))
	sb.WriteString(m.colorButtonComment("Legend"))
	sb.WriteString(m.colorButton("[Tab]"))
	sb.WriteString(m.colorButtonComment("Next"))
	sb.WriteString(m.colorButton("[Shift-Tab]"))
	sb.WriteString(m.colorButtonComment("Prev"))
	sb.WriteString(m.colorButton("[<-/->]"))
	sb.WriteString(m.colorButtonComment("Move"))
	sb.WriteString(m.colorButton("[Enter]"))
	sb.WriteString(m.colorButtonComment("Edit"))
	sb.WriteString(m.colorButton("[Esc]"))
	sb.WriteString(m.colorButtonComment("Cancel"))
	sb.WriteString(m.colorButton("[F5]"))
	sb.WriteString(m.colorButtonComment("Send"))
	//sb.WriteString("F8")
	//sb.WriteString(m.colorButtonComment("Extra"))
	//sb.WriteString(m.colorButton("[F2]"))
	//sb.WriteString(m.colorButtonComment("Show empty "))
	sb.WriteString("\n")
	sb.WriteString(m.Fun.instance.canonicalName)
	if m.Fun.instance.comb.FuncDecl.ID != nil {
		sb.WriteString("#")
		sb.WriteString(fmt.Sprintf("%08x", *m.Fun.instance.comb.FuncDecl.ID))
	}
	//sb.WriteString(" => ")
	//sb.WriteString(m.Fun.instance.resultType.Combinator().TypeDecl.Type.St.CanonicalName())
	sb.WriteString("    ")
	if m.LastError != nil {
		sb.WriteString(m.colorError(m.LastError.Error()))
	}
	sb.WriteString("\n")

	m.Fun.UIWrite(&sb, true, 0, m.Path, m)
	sb.WriteString("\n")
	//TODO - do not take into account color characters
	//var sb2 strings.Builder
	//m.Fun.UIWrite(&sb2, true, 0, m.Path, m)
	//str := sb2.String()
	//for m.Width > 0 && len(str) > m.Width {
	//	sb.WriteString(str[:m.Width])
	//	sb.WriteString("\n")
	//	str = str[m.Width:]
	//}
	//sb.WriteString(str)
	var bb ByteBuilder
	if m.ShowLegend {
		sb.WriteString(bb.PrintLegend())
		sb.WriteString("\n")
	}
	m.Fun.WriteTL2(&bb, false, false, 0, m)
	sb.WriteString(bb.Print())
	sb.WriteString("\n")
	sb.WriteString("Path: ")
	for i, pa := range m.Path {
		if i != 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("%d", pa))
	}
	return sb.String()
}

func (m *UIModel) Move(side int) {
	if len(m.Path) == 0 {
		panic("unexpected model state")
	}
	m.Path[len(m.Path)-1] += side
	childWantsSide := m.Fun.UIFixPath(-side, 0, m)
	if childWantsSide == 0 {
		return
	}
	m.Path = m.Path[:0]
	m.Fun.UIFixPath(childWantsSide, 0, m)
}

func (m *UIModel) Right() {
	m.Move(1)
}

func (m *UIModel) StartEdit(fromTab bool) {
	m.Fun.UIStartEdit(0, m, fromTab)
}

func (m *UIModel) SetCurrentEditor(e UIEditor) {
	if m.CurrentEditor != nil {
		panic("edit already started")
	}
	m.CurrentEditor = e
}

// [F1]Help [Tab]Next [^Enter]Send [Enter]Edit [Esc]Cancel  // [F5]Extra
// memcache.set # abcdef => bool (@read @readwrite) [last error]
// {"key":"str", "value":"vasya", "arr":[3, 5, 7]}
