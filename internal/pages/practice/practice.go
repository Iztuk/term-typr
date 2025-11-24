package practice

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PracticeModel struct {
	TargetText   []Glyph
	Input        []string
	CurrentIndex int

	ActiveTest bool
}

var (
	correctStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("82"))
	wrongStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("203"))
)

func InitialPracticeModel() PracticeModel {
	return PracticeModel{
		TargetText: createSessionText(),
	}
}

func (m PracticeModel) Update(msg tea.Msg) (PracticeModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		k := msg.String()

		// Conditional to start test with any other key other than ":"
		if len(k) == 1 && !m.ActiveTest && k != ":" {
			m.ActiveTest = true
		}
		if m.ActiveTest {
			switch k {
			case "backspace":
				if m.CurrentIndex != 0 {
					m.CurrentIndex--
					m.Input = m.Input[:len(m.Input)-1]

					m.TargetText[m.CurrentIndex].State = 0
				}
			default:
				if len(k) != 1 {
					return m, nil
				}

				if m.CurrentIndex >= len(m.TargetText) {
					m.ActiveTest = false
					return m, nil
				}

				m.Input = append(m.Input, k)

				m.TargetText[m.CurrentIndex] = evaluateInput(k, m.TargetText[m.CurrentIndex])

				m.CurrentIndex++

				if m.CurrentIndex >= len(m.TargetText) {
					m.ActiveTest = false
				}
			}
		}
	}
	return m, nil
}

func (m PracticeModel) View() string {
	var b strings.Builder

	if m.CurrentIndex == len(m.TargetText) {
		b.WriteString("Done! ")
		at := fmt.Sprintf("\nActive Test: %t", m.ActiveTest)
		b.WriteString(at)
		return b.String()
	}

	for i, g := range m.TargetText {
		ch := string(g.Char)
		if i != len(m.TargetText)-1 {
			ch += "  "
		}

		switch g.State {
		case 1:
			b.WriteString(correctStyle.Render(ch))
		case 2:
			b.WriteString(wrongStyle.Render(ch))
		default:
			b.WriteString(ch)
		}
	}

	return b.String()
}
