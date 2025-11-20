package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	target      []string
	input       []string
	commandMode bool
	command     string
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
		target: []string{"hello", "world", "it's", "me", "mario"},
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "ctrl+c" {
			return m, tea.Quit
		}

		if m.commandMode {
			switch k {
			case "enter":
				if m.command == "q" {
					return m, tea.Quit
				}
				m.commandMode = false
				m.command = ""
				return m, nil
			case "esc":
				m.commandMode = false
				m.command = ""
				return m, nil
			case "backspace":
				if len(m.command) > 0 {
					m.command = m.command[:len(m.command)-1]
				}
				return m, nil
			default:
				if len(k) == 1 {
					m.command += k
				}
				return m, nil
			}
		}

		switch k {
		case ":":
			m.commandMode = true
			m.command = ""
		case "backspace":
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		case " ":
			m.input = append(m.input, "_")
		default:
			m.input = append(m.input, k)
		}
	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf("Target: %s\nInput: %s\nCommand: :%s", strings.Join(m.target, "_"), strings.Join(m.input, ""), m.command)
}
