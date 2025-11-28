package tui

import "github.com/charmbracelet/lipgloss"

// General App Styles
var (
	// Typing Game Styles
	CorrectStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#A3BE8C")) // Nord Green
	WrongStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#BF616A")) // Nord Red
	TodoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#4C566A")) // Nord Grey
	CursorStyle  = lipgloss.NewStyle().Background(lipgloss.Color("#D8DEE9")).Foreground(lipgloss.Color("#2E3440"))

	// Welcome / Dashboard Styles
	LogoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#5E81AC")). // Nord Blue
			Bold(true).
			MarginBottom(1)

	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#88C0D0")). // Cyan
			Bold(true).
			MarginBottom(2)

	MenuNormalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4C566A")) // Dimmed Grey

	MenuSelectedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ECEFF4")).                      // White
				Border(lipgloss.NormalBorder(), false, false, false, true). // Left Border
				BorderForeground(lipgloss.Color("#B48EAD")).                // Purple Border
				PaddingLeft(1)

	FooterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#434C5E")).
			Italic(true)
)

