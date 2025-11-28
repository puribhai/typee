package tui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up      key.Binding
	Down    key.Binding
	Enter   key.Binding
	Quit    key.Binding
	Restart key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("k", "up"),        // Actual keys
			key.WithHelp("↑/k", "move up"), // Text for help menu
		),
		Down: key.NewBinding(
			key.WithKeys("j", "down"),
			key.WithHelp("↓/j", "move down"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter", "space"),
			key.WithHelp("enter", "select"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		Restart: key.NewBinding(
			key.WithKeys("ctrl+r"),
			key.WithHelp("ctrl+r", "restart"),
		),
	}
}
