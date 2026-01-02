package tl2pure

import (
	"fmt"
	"runtime/debug"
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

func TLGenVersion() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}
	return ""
}

func (m *UIModel) View() string {
	var sb strings.Builder
	version := ""
	if info, ok := debug.ReadBuildInfo(); ok {
		version = info.Main.Version
	}
	sb.WriteString(fmt.Sprintf("tl2client %s: ", version))
	sb.WriteString(m.colorButton("[F1]"))
	sb.WriteString(m.colorButtonComment("Help"))
	sb.WriteString("\n")
	var bb ByteBuilder
	if m.ShowLegend {
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
		sb.WriteString(m.colorButton("[F8]"))
		sb.WriteString(m.colorButtonComment("Delete"))
		sb.WriteString(m.colorButton("[F5]"))
		sb.WriteString(m.colorButtonComment("Send"))
		sb.WriteString("\n")
		sb.WriteString("TL2 legend: ")
		sb.WriteString(bb.PrintLegend())
		sb.WriteString("\n")
	}
	//sb.WriteString(m.colorButtonComment("Extra"))
	sb.WriteString(m.Fun.instance.canonicalName)
	if m.Fun.instance.comb.FuncDecl.Magic != 0 {
		sb.WriteString("#")
		sb.WriteString(fmt.Sprintf("%08x", m.Fun.instance.comb.FuncDecl.Magic))
	}
	sb.WriteString("    ")
	if m.LastError != nil {
		sb.WriteString(m.colorError(m.LastError.Error()))
	}
	//sb.WriteString(" => ")
	//sb.WriteString(m.Fun.instance.resultType.Combinator().TypeDecl.Type.St.CanonicalName())
	sb.WriteString("\n")

	m.Fun.UIWrite(&sb, true, 0, m)
	sb.WriteString("\n")

	m.Fun.WriteTL2(&bb, false, true, 0, m)
	sb.WriteString(bb.Print())
	sb.WriteString("\n")
	sb.WriteString("Path: ")
	for i, pa := range m.Path {
		if i != 0 {
			sb.WriteString(",")
		}
		sb.WriteString(fmt.Sprintf("%d", pa))
	}
	sb.WriteString(fmt.Sprintf(" cursor: %d %d", bb.cursorStart, bb.cursorFinish))
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

func (m *UIModel) StartEdit(createMode int) {
	m.Fun.UIStartEdit(0, m, createMode)
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
