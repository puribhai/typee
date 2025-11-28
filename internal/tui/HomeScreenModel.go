package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	MenuStart = iota
	MenuSettings
	MenuQuit
)

type HomeScreenModel struct {
	keys KeyMap

	width   int
	height  int
	cursor  int
	options []string

	Choice   int
	Selected bool
}

func NewHomeScreenModel() HomeScreenModel {
	return HomeScreenModel{
		keys:     DefaultKeyMap(),
		cursor:   0,
		options:  []string{"Start Test", "Settings", "Quit"},
		Choice:   MenuStart,
		Selected: false,
	}
}

func (m HomeScreenModel) Init() tea.Cmd {
	return nil
}

func (m HomeScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, m.keys.Down):
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case key.Matches(msg, m.keys.Quit):
			m.Choice = MenuQuit
			m.Selected = true
			return m, tea.Quit
		case key.Matches(msg, m.keys.Enter):
			m.Choice = m.cursor
			m.Selected = true
		}
	}
	return m, nil
}

func (m HomeScreenModel) View() string {
	// 1. Render the Logo
	logo := LogoStyle.Render(logoASCII)
	title := TitleStyle.Render("Type Like a Pro.")

	// 2. Render the Menu
	var menuItems []string
	for i, option := range m.options {
		cursor := "  " // Empty space for unselected
		style := MenuNormalStyle

		if m.cursor == i {
			cursor = "> " // Cursor for selected
			style = MenuSelectedStyle
		}

		// Calculate "hotkey" hint (e.g., [s]tart) - optional visual
		// For now, simple list
		line := style.Render(cursor + option)
		menuItems = append(menuItems, line)
	}

	menuBlock := lipgloss.JoinVertical(lipgloss.Left, menuItems...)

	// 3. Render Footer/Help
	helpText := FooterStyle.Render("j/k to navigate • enter to select")

	// 4. Join everything vertically
	content := lipgloss.JoinVertical(lipgloss.Center,
		logo,
		title,
		menuBlock,
		"\n", // Spacer
		helpText,
	)

	// 5. Center the entire block in the terminal window
	// lipgloss.Place takes (width, height, horizontalAlign, verticalAlign, content)
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		content,
	)
}

