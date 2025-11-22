package practice

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type PracticeModel struct {
	TargetText []string
	Input      string
}

func InitialPracticeModel() PracticeModel {
	return PracticeModel{
		TargetText: createSessionText(),
	}
}

func (m PracticeModel) Update(msg tea.Msg) (PracticeModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.Input += msg.String()
	}
	return m, nil
}

func (m PracticeModel) View() string {
	text := fmt.Sprintf("%v %s", strings.Join(m.TargetText, " "), m.Input)
	return text
}
