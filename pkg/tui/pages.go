package tui

// PageView determines which page to render based on the current page state
func PageView(m model) string {
	page := ""
	switch m.currentPage {
	case page1:
		page = Page1View(m)
	case page2:
		page = Page2View(m)
	case page3:
		page = Page3View(m)
	default:
		page = m.theme.Base().Render("Unknown page")
	}

	return page
}

// Page1View renders the content for Page 1
func Page1View(m model) string {
	return m.theme.Base().Render("This is the content of Page 1")
}

// Page2View renders the content for Page 2
func Page2View(m model) string {
	return m.theme.Base().Render("This is the content of Page 2")
}

// Page3View renders the content for Page 3
func Page3View(m model) string {
	return m.theme.Base().Render("This is the content of Page 3")
}
