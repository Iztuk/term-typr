package practice

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type PracticeModel struct {
	TargetText   []string
	Input        []string
	CurrentIndex int
}

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
			m.Input = append(m.Input, k)
			m.CurrentIndex++
		}
	}
	return m, nil
}

func (m PracticeModel) View() string {
	text := fmt.Sprintf("%v %s", strings.Join(m.TargetText, " "), m.Input)
	return text
}
