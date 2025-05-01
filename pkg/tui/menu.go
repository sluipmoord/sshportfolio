package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// MenuView renders the menu with selectable options
func (m model) MenuView() string {
	bold := m.theme.TextAccent().Bold(true).Render
	render := m.theme.Base().Render

	var renderedMenu []string
	for i, item := range m.pages {
		cursor := " " // no cursor
		if i == m.cursor {
			cursor = ">" // cursor for the selected item
			item = bold(item)
		}
		renderedMenu = append(renderedMenu, render(fmt.Sprintf("%s %s", cursor, item)))
	}

	return lipgloss.JoinVertical(lipgloss.Left, renderedMenu...)
}

// MenuUpdate handles menu-specific key events
func MenuUpdate(m model, msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.pages)-1 {
				m.cursor++
			}
		case "enter":
			switch m.cursor {
			case 0:
				m.currentPage = page1
			case 1:
				m.currentPage = page2
			case 2:
				m.currentPage = page3
			}
		}
	}

	return m, nil
}
