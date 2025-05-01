package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Portfolio struct {
	Pages  []string
	Cursor int
}

func (m Portfolio) Init() tea.Cmd {
	return nil
}

func (m Portfolio) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.Cursor < len(m.Pages)-1 {
				m.Cursor++
			}
		case "enter":
			return m, func() tea.Msg {
				return fmt.Sprintf("You selected: %s", m.Pages[m.Cursor])
			}
		}
	}

	return m, nil
}

func (m Portfolio) View() string {
	view := "My Portfolio\n\n"
	for i, page := range m.Pages {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		view += fmt.Sprintf("%s %s\n", cursor, page)
	}
	view += "\nUse up/down to navigate, enter to select, and q to quit."
	return view
}
