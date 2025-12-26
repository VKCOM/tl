// Copyright 2025 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vkcom/tl/internal/tl2pure"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/tlcodegen"
)

type model struct {
	impl *tl2pure.UIModel
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// _, _ = fmt.Fprintf(os.Stderr, "keyMsg: %q\n", msg.String())
		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			//if m.cursor > 0 {
			//	m.cursor--
			//}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			//if m.cursor < len(m.choices)-1 {
			//	m.cursor++
			//}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			//_, ok := m.selected[m.cursor]
			//if ok {
			//	delete(m.selected, m.cursor)
			//} else {
			//	m.selected[m.cursor] = struct{}{}
			//}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	return m.impl.View()
}

func parseTL2File(file string) (tlast.TL2File, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return tlast.TL2File{}, fmt.Errorf("error reading schema file %q - %w", file, err)
	}
	dataStr := string(data)
	return tlast.ParseTL2File(dataStr, file, tlast.LexerOptions{LexerLanguage: tlast.TL2}, os.Stdout)
}

func main() {
	log.Printf("tlclient WIP version: %s", tlcodegen.TLGenVersion())

	var runUI bool
	flag.BoolVar(&runUI, "ui", false, "run in UI mode")
	flag.Parse()

	kernel := tl2pure.NewKernel()
	if len(flag.Args()) == 0 {
		log.Printf("tlclient requires 1 or more tl2 files")
		os.Exit(2)
	}
	for _, arg := range flag.Args() {
		f, err := parseTL2File(arg)
		if err != nil {
			log.Printf("%v", err)
			os.Exit(3)
		}
		kernel.AddFile(f)
	}
	err := kernel.Compile()
	if err != nil {
		log.Printf("%v", err)
		os.Exit(4)
	}
	//	tlt := kernel.TopLeveTypeInstances()
	tlt := kernel.AllTypeInstances()
	rnd := rand.New(rand.NewSource(1))
	for _, t := range tlt {
		log.Printf("type instance: %v", t.CanonicalName())
		for i := 0; i < 5; i++ {
			val := t.CreateValue()
			val.Random(rnd)
			tl2 := val.WriteTL2(nil, false, nil)
			js := val.WriteJSON(nil, nil)
			log.Printf(".   TL2: %x", tl2)
			log.Printf(".   JSON: %s", js)
		}
	}

	if !runUI {
		return
	}

	name := tlast.TL2TypeName{
		Namespace: "memcache",
		Name:      "set",
	}
	fun := kernel.GetFunctionInstance(name)
	if fun == nil {
		log.Printf("function %q not found", name)
		return
	}

	uiModel := &tl2pure.UIModel{
		Fun:           fun.CreateValueObject(),
		Path:          nil,
		CurrentEditor: nil,
		LastError:     nil,
	}

	p := tea.NewProgram(model{impl: uiModel})
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
