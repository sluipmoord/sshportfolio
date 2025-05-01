package tui

import (
	"fmt"
)

// HeaderView renders the header of the application
func (m model) HeaderView() string {
	bold := m.theme.TextAccent().Bold(true).Render
	base := m.theme.Base().Render

	header := base("My Portfolio Page:")
	header += bold(fmt.Sprintf(" %d", m.currentPage))

	return m.theme.Base().PaddingBottom(1).Render(header)
}
