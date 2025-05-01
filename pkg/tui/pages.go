package tui

import (
	"log/slog"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/sluipmoord/sshportfolio/pkg/assets"
)

// PageView determines which page to render based on the current page state
func (m model) PageView() string {
	// Find the current page in the pages slice
	for _, p := range m.pages {
		if p.id == m.currentPage {
			// If there's a viewport available for this page, use it
			if viewport, ok := m.viewports[p.id]; ok && viewport.Height > 0 {
				return viewport.View()
			}
			// Otherwise, generate the content (should trigger content loading)
			return p.content(m)
		}
	}

	return m.theme.Base().Render("Unknown page")
}

// AboutMeView renders the content from the about.md embedded asset
func (m model) AboutMeView() string {
	// Get the about.md content from the embedded assets
	content, err := assets.GetAsset("about.md")
	if err != nil {
		slog.Error("Error reading about.md from embedded assets", "error", err)
		return m.theme.TextError().Render(
			"About Me content not found. Please make sure about.md exists in the embedded assets.",
		)
	}

	// Render the markdown content
	out, err := glamour.Render(string(content), "dark")
	if err != nil {
		slog.Error("Error rendering About Me content", "error", err)
		return m.theme.TextError().Render("Error rendering About Me content")
	}

	return m.theme.Base().Render(out)

}

// ReadmeView renders the content for the README page
func (m model) ReadmeView() string {
	// Return the viewport content directly if it's already initialized
	if vp, ok := m.viewports[readmePage]; ok && vp.Height > 0 {
		return vp.View()
	}

	// Otherwise, this is a fallback (should not happen with proper initialization)
	// Try multiple possible locations for README.md
	possiblePaths := []string{
		"README.md",          // Current directory
		"./README.md",        // Explicit current directory
		"../README.md",       // Parent directory
		"../../README.md",    // Two levels up
		"../../../README.md", // Three levels up
	}

	var content []byte
	var err error
	var foundPath string

	for _, path := range possiblePaths {
		content, err = os.ReadFile(path)
		if err == nil {
			foundPath = path
			slog.Info("README found", "path", path)
			break
		}
	}

	if err != nil {
		slog.Error("Error reading README from all possible paths", "error", err)
		return m.theme.TextError().Render("README file not found. Please make sure the README.md exists in the project.")
	}

	out, err := glamour.Render(string(content), "dark")
	if err != nil {
		slog.Error("Error rendering README", "error", err, "path", foundPath)
		return m.theme.TextError().Render("Error rendering README")
	}

	return m.theme.Base().Render(out)
}

func (m model) ReadmeUpdate(msg tea.Msg) (model, tea.Cmd) {
	return m.PageUpdate(readmePage, msg)
}

// AboutMeUpdate handles scrolling for the About Me page
func (m model) AboutMeUpdate(msg tea.Msg) (model, tea.Cmd) {
	return m.PageUpdate(aboutMe, msg)
}

// PageUpdate handles scrolling for any page
func (m model) PageUpdate(pageID pageID, msg tea.Msg) (model, tea.Cmd) {
	if viewport, ok := m.viewports[pageID]; ok {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "j", "down":
				viewport.ScrollDown(1)
				m.viewports[pageID] = viewport
				return m, nil
			case "k", "up":
				viewport.ScrollUp(1)
				m.viewports[pageID] = viewport
				return m, nil
			case "g", "home":
				viewport.GotoTop()
				m.viewports[pageID] = viewport
				return m, nil
			case "G", "end":
				viewport.GotoBottom()
				m.viewports[pageID] = viewport
				return m, nil
			case "d", "pgdown":
				viewport.PageDown()
				m.viewports[pageID] = viewport
				return m, nil
			case "u", "pgup":
				viewport.PageUp()
				m.viewports[pageID] = viewport
				return m, nil
			case "f":
				viewport.PageDown()
				m.viewports[pageID] = viewport
				return m, nil
			case "b":
				viewport.PageUp()
				m.viewports[pageID] = viewport
				return m, nil
			}
		}
	}
	return m, nil
}

// InitializeViewport creates a viewport for a page and populates it with content
func (m *model) InitializeViewport(pageID pageID) tea.Cmd {
	return func() tea.Msg {
		// Find the page in the pages slice
		var page *page
		for i := range m.pages {
			if m.pages[i].id == pageID {
				page = &m.pages[i]
				break
			}
		}

		if page == nil {
			return nil // Page not found
		}

		// Generate the content
		content := page.content(*m)

		// Create a new viewport
		vp := viewport.New(m.widthContent, m.heightContent)
		vp.SetContent(content)

		// Store in the map
		m.viewports[pageID] = vp

		return nil
	}
}
