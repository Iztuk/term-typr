package main

import (
	"fmt"
	"os"
	"term-typr/internal/pages/menu"
	"term-typr/internal/pages/practice"

	tea "github.com/charmbracelet/bubbletea"
)

func (m AppModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func initialModel() AppModel {
	return AppModel{
		menu:     menu.InitialMenuModel(),
		practice: practice.InitialPracticeModel(),
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
