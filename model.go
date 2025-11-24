package main

import (
	"term-typr/internal/pages/menu"
	"term-typr/internal/pages/practice"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AppModel struct {
	currentPage string

	menu     menu.MenuModel
	practice practice.PracticeModel

	target      []string
	input       []string
	commandMode bool
	command     string

	width  int
	height int
}

var (
	appStyle = lipgloss.NewStyle().
			Padding(0, 0)

	cmdLineStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			Background(lipgloss.Color("235")).
			Padding(0, 0)
)

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		k := msg.String()

		if k == "ctrl+c" {
			return m, tea.Quit
		}

		if m.commandMode && !m.practice.ActiveTest {
			switch k {
			case "enter":
				// Cases for page navigation commands
				switch m.command {
				case "q":
					return m, tea.Quit
				case "m", "menu":
					m.currentPage = "menu"
				case "p", "practice":
					m.currentPage = "practice"
				}

				// Cases for practice commands
				if m.currentPage == "practice" {
					switch m.command {
					// Restarts the same test
					case "r":
						m.practice = practice.RestartSessionText(m.practice)
					// Creates a new test
					case "t":
						m.practice = practice.InitialPracticeModel()
					}
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

		if k == ":" && m.currentPage == "practice" && m.practice.ActiveTest {
		} else {
			switch k {
			case ":":
				m.commandMode = true
				m.command = ""
			}
		}
	}

	var cmd tea.Cmd
	switch m.currentPage {
	case "menu":
		m.menu, cmd = m.menu.Update(msg)
	case "practice":
		m.practice, cmd = m.practice.Update(msg)
	default:
		m.currentPage = "menu"
	}

	return m, cmd
}

func (m AppModel) View() string {
	// content := fmt.Sprintf("Target: %s\nInput: %s\nCommand: :%s", strings.Join(m.target, "_"), strings.Join(m.input, ""), m.command)

	var content string

	switch m.currentPage {
	case "menu":
		content = m.menu.View()
	case "practice":
		content = m.practice.View()

		targetTextStyle := lipgloss.NewStyle().
			Width(m.width - 15).
			AlignHorizontal(lipgloss.Center)

		content = targetTextStyle.Render(content)
	default:
		content = m.menu.View()

	}

	var cmdLineText string
	if m.commandMode {
		cmdLineText = ":" + m.command
	}

	if m.width == 0 || m.height == 0 {
		cmdLine := cmdLineStyle.Render(cmdLineText)
		body := lipgloss.JoinVertical(lipgloss.Left, content, cmdLine)
		return appStyle.Render(body)
	}

	innerWidth := m.width
	innerHeight := m.height
	if innerWidth < 0 {
		innerWidth = 0
	}
	if innerHeight < 0 {
		innerHeight = 0
	}

	bodyHeight := innerHeight - 1
	if bodyHeight < 1 {
		bodyHeight = 1
	}

	bodyArea := lipgloss.Place(
		innerWidth,
		bodyHeight,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)

	cmdLine := cmdLineStyle.
		Width(innerWidth).
		Render(cmdLineText)

	screen := lipgloss.JoinVertical(
		lipgloss.Left,
		bodyArea,
		cmdLine,
	)

	centeredContent := lipgloss.Place(
		innerWidth,
		innerHeight,
		lipgloss.Center,
		lipgloss.Center,
		screen,
	)

	return appStyle.
		Width(m.width).
		Height(m.height).
		Render(centeredContent)
}
