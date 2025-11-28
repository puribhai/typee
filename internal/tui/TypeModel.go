package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TypeModel struct {
	TargetString string
	OutputString string
	StartTime    time.Time
	Finished     bool
	Wpm          int
}

func NewTypeModel() TypeModel {
	return TypeModel{
		TargetString: "a quick brown fox jumps over the lazy dog",
		OutputString: "",
		StartTime:    time.Time{},
		Finished:     false,
	}
}

func (m TypeModel) Init() tea.Cmd {
	return nil
}

func (m TypeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyBackspace:
			if len(m.OutputString) > 0 {
				m.OutputString = m.OutputString[:len(m.OutputString)-1]
			}

		case tea.KeyRunes, tea.KeySpace:
			// Start timer on first keystroke
			if m.StartTime.IsZero() {
				m.StartTime = time.Now()
			}

			// Add the character typed
			m.OutputString += msg.String()

			// Check if finished
			if len(m.OutputString) >= len(m.TargetString) {
				m.Finished = true
				m.calculateWPM()
				// We don't return tea.Quit here anymore because we want
				// the Session to handle the transition (e.g. to Results screen)
				return m, nil
			}
		}
	}
	return m, nil
}

func (m *TypeModel) calculateWPM() {
	duration := time.Since(m.StartTime).Minutes()
	wordCount := float64(len(m.OutputString)) / 5.0
	m.Wpm = int(wordCount / duration)
}

func (m TypeModel) View() string {
	// If finished, the Session might switch views, but we can return a placeholder just in case
	if m.Finished {
		return fmt.Sprintf("\nFinished! WPM: %d\nPress Ctrl+C to exit.\n", m.Wpm)
	}

	var s string
	for i, char := range m.TargetString {
		if i >= len(m.OutputString) {
			if i == len(m.OutputString) {
				s += CursorStyle.Render(string(char))
			} else {
				s += TodoStyle.Render(string(char))
			}
		} else {
			typedChar := rune(m.OutputString[i])
			if typedChar == char {
				s += CorrectStyle.Render(string(char))
			} else {
				if typedChar == ' ' {
					s += WrongStyle.Render("_")
				} else {
					s += WrongStyle.Render(string(char))
				}
			}
		}
	}
	return "\n" + s + "\n"
}
