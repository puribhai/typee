package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	StateHome sessionState = iota
	StateTyping
	StateResults
)

type Model struct {
	width  int
	height int

	home   HomeScreenModel
	typing TypeModel

	state sessionState
}

func NewModel() Model {
	return Model{
		home:   NewHomeScreenModel(),
		typing: NewTypeModel(),
		state:  StateHome,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		m.home.Init(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		newHome, cmdHome := m.home.Update(msg)
		m.home = newHome.(HomeScreenModel)
		cmds = append(cmds, cmdHome)

		newTyping, cmdTyping := m.typing.Update(msg)
		m.typing = newTyping.(TypeModel)
		cmds = append(cmds, cmdTyping)
		return m, tea.Batch(cmds...)

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	}
	switch m.state {
	case StateHome:
		newHome, newCmd := m.home.Update(msg)
		m.home = newHome.(HomeScreenModel)
		cmd = newCmd

		if m.home.Selected {
			if m.home.Choice == MenuStart {
				m.state = StateTyping
				m.typing = NewTypeModel()
				m.typing.StartTime = time.Now()
				m.typing.Update(tea.WindowSizeMsg{Width: m.width, Height: m.height})

				return m, m.typing.Init()
			}
		} else if m.home.Choice == MenuQuit {
			return m, tea.Quit
		}
	case StateTyping:
		newTyping, newCmd := m.typing.Update(msg)
		m.typing = newTyping.(TypeModel)
		cmd = newCmd
	}
	return m, cmd
}

func (m Model) View() string {
	switch m.state {
	case StateHome:
		return m.home.View()
	case StateTyping:
		return m.typing.View()
	default:
		return "unknown state"
	}
}
