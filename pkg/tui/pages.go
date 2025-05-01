package tui

// PageView determines which page to render based on the current page state
func (m model) PageView() string {
	page := ""
	switch m.currentPage {
	case page1:
		page = m.Page1View()
	case page2:
		page = m.Page2View()
	case page3:
		page = m.Page3View()
	default:
		page = m.theme.Base().Render("Unknown page")
	}

	return page
}

// Page1View renders the content for Page 1
func (m model) Page1View() string {
	return m.theme.Base().Render("This is the content of Page 1")
}

// Page2View renders the content for Page 2
func (m model) Page2View() string {
	return m.theme.Base().Render("This is the content of Page 2")
}

// Page3View renders the content for Page 3
func (m model) Page3View() string {
	return m.theme.Base().Render("This is the content of Page 3")
}
