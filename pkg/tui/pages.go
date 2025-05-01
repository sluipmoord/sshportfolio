package tui

import (
	"log/slog"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

// PageView determines which page to render based on the current page state
func (m model) PageView() string {
	// Find the current page in the pages slice
	for _, p := range m.pages {
		if p.id == m.currentPage {
			return p.content(m)
		}
	}

	return m.theme.Base().Render("Unknown page")
}

// ReadmeView renders the content for the README page
func (m model) ReadmeView() string {
	// Return the viewport content directly if it's already initialized
	if m.readmeViewport.Height > 0 {
		return m.readmeViewport.View()
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
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.readmeViewport.ScrollDown(1)
			return m, nil
		case "k", "up":
			m.readmeViewport.ScrollUp(1)
			return m, nil
		case "g", "home":
			m.readmeViewport.GotoTop()
			return m, nil
		case "G", "end":
			m.readmeViewport.GotoBottom()
			return m, nil
		case "d", "pgdown":
			m.readmeViewport.PageDown()
			return m, nil
		case "u", "pgup":
			m.readmeViewport.PageUp()
			return m, nil
		case "f":
			m.readmeViewport.PageDown()
			return m, nil
		case "b":
			m.readmeViewport.PageUp()
			return m, nil
		}
	}

	return m, nil
}
