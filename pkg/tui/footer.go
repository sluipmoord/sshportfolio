package tui

import (
	"fmt"
)

// FooterView renders the footer of the application with navigation hints
func FooterView(m model) string {
	bold := m.theme.TextAccent().Bold(true).Render
	base := m.theme.Base().Render

	esc := bold("q")
	back := bold("esc")
	footer := base(fmt.Sprintf("%s back | %s quit", back, esc))

	return m.theme.Base().PaddingTop(1).Render(footer)
}
