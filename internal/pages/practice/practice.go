package practice

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type PracticeModel struct {
	input string
}

func InitialPracticeModel() PracticeModel {
	return PracticeModel{}
}

func (m PracticeModel) Update(msg tea.Msg) (PracticeModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.input += msg.String()
	}
	return m, nil
}

func (m PracticeModel) View() string {
	text := fmt.Sprintf("%s", m.input)
	return text
}
