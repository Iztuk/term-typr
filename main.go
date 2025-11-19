package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	counter int
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	return model{
		counter: 0,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "ctrl+c" {
			return m, tea.Quit
		}

		switch k {
		case "up":
			m.counter++
		case "down":
			m.counter--
		}
	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("Counter: %d", m.counter)
}
