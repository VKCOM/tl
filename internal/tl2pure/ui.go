package tl2pure

import (
	"fmt"
	"strings"

	"github.com/TwiN/go-color"
)

type UIEditor interface {
	WriteUI(sb *strings.Builder)
}

type UIModel struct {
	Fun KernelValueObject

	Path []int // selected field on every hierarchy level, 0 means union constructor

	CurrentEditor UIEditor

	LastError error
}

func (m UIModel) View() string {
	var sb strings.Builder
	sb.WriteString("[F1]")
	sb.WriteString(color.InBlackOverCyan("Help  "))
	sb.WriteString("[Tab]")
	sb.WriteString(color.InBlackOverCyan("Next  "))
	sb.WriteString("[^Enter]")
	sb.WriteString(color.InBlackOverCyan("Send  "))
	//sb.WriteString("F8")
	//sb.WriteString(color.InBlackOverCyan("Extra"))
	sb.WriteString("[Enter]")
	sb.WriteString(color.InBlackOverCyan("Edit  "))
	sb.WriteString("[Esc]")
	sb.WriteString(color.InBlackOverCyan("Cancel"))
	sb.WriteString("\n")
	sb.WriteString(m.Fun.instance.canonicalName)
	if m.Fun.instance.comb.FuncDecl.ID != nil {
		sb.WriteString("#")
		sb.WriteString(fmt.Sprintf("%08x", *m.Fun.instance.comb.FuncDecl.ID))
	}
	sb.WriteString(" => ")
	sb.WriteString(m.Fun.instance.resultType.CanonicalName())
	sb.WriteString("    ")
	if m.LastError != nil {
		sb.WriteString(color.InRed(m.LastError.Error()))
	}
	sb.WriteString("\n")
	m.Fun.WriteUI(&sb, true, 0, m.Path, &m)
	return sb.String()
}

// [F1]Help [Tab]Next [^Enter]Send [Enter]Edit [Esc]Cancel  // [F5]Extra
// memcache.set # abcdef => bool (@read @readwrite) [last error]
// {"key":"str", "value":"vasya", "arr":[3, 5, 7]}
