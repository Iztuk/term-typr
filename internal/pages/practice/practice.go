package practice

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PracticeModel struct {
	TargetText   []Glyph
	Input        []string
	CurrentIndex int
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
		switch k {
		case "backspace":
			if m.CurrentIndex != 0 {
				m.Input = m.Input[:len(m.Input)-1]
				m.CurrentIndex--
			}
		default:
			if len(k) != 1 {
				return m, nil
			}

			m.Input = append(m.Input, k)

			m.TargetText[m.CurrentIndex] = evaluateInput(k, m.TargetText[m.CurrentIndex])

			m.CurrentIndex++
		}
	}
	return m, nil
}

func (m PracticeModel) View() string {
	var b strings.Builder

	for _, g := range m.TargetText {
		ch := string(g.Char)

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
