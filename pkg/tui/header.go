package tui

// HeaderView renders the header of the application
func (m model) HeaderView() string {
	bold := m.theme.TextAccent().Bold(true).Render
	base := m.theme.Base().Render

	// If we're in the menu, display a welcome message
	if m.currentPage == menuPage {
		return m.theme.Base().PaddingBottom(1).Render(bold("Welcome to My Portfolio"))
	}

	// Find the current page title
	var pageTitle string
	for _, p := range m.pages {
		if p.id == m.currentPage {
			pageTitle = p.title
			break
		}
	}

	// If page title was found, display it
	if pageTitle != "" {
		return m.theme.Base().PaddingBottom(1).Render(bold(pageTitle))
	}

	// Fallback to generic header
	return m.theme.Base().PaddingBottom(1).Render(base("My Portfolio"))
}
