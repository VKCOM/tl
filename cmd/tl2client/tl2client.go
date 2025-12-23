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

	"github.com/TwiN/go-color"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vkcom/tl/internal/tl2pure"
	"github.com/vkcom/tl/internal/tlast"
	"github.com/vkcom/tl/internal/tlcodegen"
)

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
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
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "What should we buy at the market?\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] ", cursor, checked)
		s += color.InGreen(choice) + "\n"
	}
	// if m.cursor == len(m.choices)-1 {
	// for i := 0; i < 20; i++ {
	// s += "---Additional line\n"
	// }
	// }

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
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

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
