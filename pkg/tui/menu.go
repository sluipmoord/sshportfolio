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
	for i, page := range m.pages {
		cursor := " " // no cursor
		title := page.title

		if i == m.cursor {
			cursor = ">" // cursor for the selected item
			title = bold(title)
		}
		renderedMenu = append(renderedMenu, render(fmt.Sprintf("%s %s", cursor, title)))
	}

	return lipgloss.JoinVertical(lipgloss.Left, renderedMenu...)
}

// MenuUpdate handles menu-specific key events
func (m model) MenuUpdate(msg tea.Msg) (model, tea.Cmd) {
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
			if m.cursor >= 0 && m.cursor < len(m.pages) {
				m.currentPage = m.pages[m.cursor].id
			}
		}
	}

	return m, nil
}
