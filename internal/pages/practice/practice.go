package practice

import (
	"fmt"
	"strings"
	"time"

	"github.com/NimbleMarkets/ntcharts/linechart"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PracticeModel struct {
	TargetText   []Glyph
	CurrentInput []string
	CurrentIndex int
	TotalInput   []string
	TotalTime    time.Time

	ActiveTest bool
	StopWatch  StopWatch

	StatsChart linechart.Model
}

// Styles
var (
	// Input coloring styles
	correctStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("82"))
	wrongStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("203"))

	// Session stats styles
	defaultStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63")) // purple
	lineStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("4")) // blue
	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("6")) // cyan

)

var replacedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("5")) // pink

var axisStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("3")) // yellow

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
			m.StopWatch.Start()
		}
		if m.ActiveTest {
			switch k {
			case "backspace":
				if m.CurrentIndex != 0 {
					m.CurrentIndex--
					m.CurrentInput = m.CurrentInput[:len(m.CurrentInput)-1]

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

				m.CurrentInput = append(m.CurrentInput, k)
				m.TotalInput = append(m.TotalInput, k)

				m.TargetText[m.CurrentIndex] = evaluateInput(k, m.TargetText[m.CurrentIndex])

				m.CurrentIndex++

				if m.CurrentIndex >= len(m.TargetText) {
					m.ActiveTest = false
					m.StopWatch.Stop()
					m.StatsChart = generateSessionStatsChart()
				}
			}
		}
	}
	return m, nil
}

func (m PracticeModel) View() string {
	var b strings.Builder

	if m.CurrentIndex >= len(m.TargetText) {
		wpm := fmt.Sprintf("Raw WPM: %d WPM: %d Accuracy: %.2f\n", evaluateRawWPM(m.TotalInput, m.StopWatch.elapsed), evaluateWPM(m.TargetText, m.StopWatch.elapsed), evaluateAccuracy(m.TargetText, m.TotalInput))
		b.WriteString(wpm)
		b.WriteString("\n:r to restart and :t for a new test\n\n")

		s := "any key to draw randomized line, `r` to reset, `q/ctrl+c` to quit\n"
		s += lipgloss.JoinHorizontal(lipgloss.Top,
			defaultStyle.Render("DrawBrailleLine()\n"+m.StatsChart.View()),
		) + "\n"
		b.WriteString(s)
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
