package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Projects []string
	Cursor   int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(m.Projects)-1 {
				m.Cursor++
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	view := "My Portfolio\n\n"
	for i, project := range m.Projects {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		view += fmt.Sprintf("%s %s\n", cursor, project)
	}
	view += "\nPress q to quit."
	return view
}
