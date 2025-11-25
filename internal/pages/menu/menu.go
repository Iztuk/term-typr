package menu

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuModel struct {
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("12")).
			Padding(0, 1)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("249"))

	hintKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("150")).
			Bold(true)

	hintTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245"))
)

func InitialMenuModel() MenuModel {
	return MenuModel{}
}

func (m MenuModel) Update(msg tea.Msg) (MenuModel, tea.Cmd) {
	return m, nil
}

func (m MenuModel) View() string {
	// Left side: "logo" + subtitle
	left := lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render("TERM-TYPR"),
		subtitleStyle.Render("A minimal typing practice tool for your terminal."),
	)

	// Right side: hints like Neovim's start screen
	right := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			hintKeyStyle.Render(":h"),
			lipgloss.NewStyle().Render("  "),
			hintTextStyle.Render("Open help"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			hintKeyStyle.Render(":p"),
			lipgloss.NewStyle().Render("  "),
			hintTextStyle.Render("Start typing practice"),
		),
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			hintKeyStyle.Render(":q"),
			lipgloss.NewStyle().Render("  "),
			hintTextStyle.Render("Quit term-typr"),
		),
	)

	// Put them side-by-side, like Neovim's logo + hints
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		lipgloss.NewStyle().Padding(0, 4).Render(""), // spacer between columns
		right,
	)
}
