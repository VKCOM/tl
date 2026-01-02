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
	"time"

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
	case tl2pure.UIBlinkMsg:
		m.impl.Blink = msg
	case tea.WindowSizeMsg:
		m.impl.Width = msg.Width
		m.impl.Height = msg.Height
	case tea.KeyMsg:
		m.impl.LastError = nil
		if msg.Type == tea.KeyBackspace {
			if m.impl.CurrentEditor != nil {
				m.impl.CurrentEditor.OnBackspace(m.impl)
				return m, nil
			}
		}
		if msg.Type == tea.KeyF1 {
			m.impl.ShowLegend = !m.impl.ShowLegend
			return m, nil
		}
		if msg.Type == tea.KeyF8 {
			if m.impl.CurrentEditor == nil {
				m.impl.Fun.UIKey(0, m.impl, false, true, false, false)
				return m, nil
			}
		}
		if msg.Type == tea.KeyTab {
			if m.impl.CurrentEditor != nil {
				m.impl.CurrentEditor.OnTab(m.impl, 1)
				return m, nil
			}
			m.impl.Move(1)
			m.impl.StartEdit(0)
			return m, nil
		}
		if msg.Type == tea.KeyShiftTab {
			if m.impl.CurrentEditor != nil {
				m.impl.CurrentEditor.OnTab(m.impl, -1)
				return m, nil
			}
			m.impl.Move(-1)
			m.impl.StartEdit(0)
			return m, nil
		}
		if msg.Type == tea.KeyEnter {
			if m.impl.CurrentEditor != nil {
				m.impl.CurrentEditor.OnEnter(m.impl)
				return m, nil
			} else {
				if m.impl.CurrentEditor == nil {
					m.impl.StartEdit(1)
				}
				return m, nil
			}
		}
		if msg.Type == tea.KeyEscape || msg.Type == tea.KeyF9 { // KeyF9 for testing in IDE
			if m.impl.CurrentEditor != nil {
				m.impl.CurrentEditor.OnEscape(m.impl)
				return m, nil
			}
		}
		if msg.Type == tea.KeyRunes {
			if m.impl.CurrentEditor == nil {
				m.impl.StartEdit(2)
			}
			if m.impl.CurrentEditor != nil {
				m.impl.CurrentEditor.OnRune(msg.String(), m.impl)
			}
			return m, nil
		}
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "right":
			if m.impl.CurrentEditor == nil {
				wasLen := len(m.impl.Path)
				m.impl.Move(1)
				if len(m.impl.Path) > wasLen { // do not enter deep into objects
					m.impl.Path = m.impl.Path[:wasLen]
				}
			}
			return m, nil
		case "left":
			if m.impl.CurrentEditor == nil {
				wasLen := len(m.impl.Path)
				m.impl.Move(-1)
				if len(m.impl.Path) > wasLen { // do not enter deep into objects
					m.impl.Path = m.impl.Path[:wasLen]
				}
			}
			return m, nil
		case "up":
			if m.impl.CurrentEditor == nil && len(m.impl.Path) > 1 {
				m.impl.Path = m.impl.Path[:len(m.impl.Path)-1]
			}
			return m, nil
		case "down":
			if m.impl.CurrentEditor == nil {
				wasLen := len(m.impl.Path)
				m.impl.Fun.UIFixPath(-1, 0, m.impl)
				if len(m.impl.Path) > wasLen+1 { // level +1 only
					m.impl.Path = m.impl.Path[:wasLen+1]
				}
			}
		}
	}
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
			var bb tl2pure.ByteBuilder
			val.WriteTL2(&bb, false, false, 0, nil)
			js := val.WriteJSON(nil, nil)
			log.Printf(".   TL2: %s", bb.Print())
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

	uiModel.Fun.UIFixPath(-1, 0, uiModel)

	p := tea.NewProgram(model{impl: uiModel})
	go func() {
		for {
			p.Send(tl2pure.UIBlinkMsg{On: true})
			time.Sleep(500 * time.Millisecond)
			p.Send(tl2pure.UIBlinkMsg{On: false})
			time.Sleep(500 * time.Millisecond)
		}
	}()
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
