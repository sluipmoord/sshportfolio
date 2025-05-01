package tui

import (
	"context"
	"log/slog"
	"math"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/sluipmoord/sshportfolio/pkg/tui/theme"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// page represents a content page in the application
type page struct {
	id      pageID
	title   string
	content func(m model) string
}

type pageID = int
type cursor = int
type size = int

const (
	menuPage pageID = iota
	aboutMe
	page2
	page3
	readmePage
)

const (
	undersized size = iota
	small
	medium
	large
)

// contextKey is a custom type for context keys to avoid collisions
type contextKey string

// Context keys
const (
	clientIPKey contextKey = "client_ip"
)

type model struct {
	renderer *lipgloss.Renderer
	command  []string
	context  context.Context
	theme    theme.Theme

	viewportWidth   int
	viewportHeight  int
	widthContainer  int
	heightContainer int
	widthContent    int
	heightContent   int
	size            size

	currentPage pageID
	pages       []page

	cursor cursor

	// Viewport for scrollable content
	viewports map[pageID]viewport.Model
}

func NewModel(
	renderer *lipgloss.Renderer,
	clientIP *string,
	command []string,
) (tea.Model, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, clientIPKey, clientIP)

	// Create pages with titles and content functions
	pages := []page{
		{id: aboutMe, title: "About Me", content: func(m model) string {
			return m.AboutMeView()
		}},
		{id: page2, title: "Projects", content: func(m model) string {
			return m.theme.Base().Render("This is the content of Page 2")
		}},
		{id: page3, title: "Skills", content: func(m model) string {
			return m.theme.Base().Render("This is the content of Page 3")
		}},
		{id: readmePage, title: "README", content: func(m model) string {
			return m.ReadmeView()
		}},
	}

	return model{
		renderer:    renderer,
		context:     ctx,
		command:     command,
		cursor:      0,
		currentPage: menuPage,
		pages:       pages,
		theme:       theme.BasicTheme(renderer, nil),
		viewports:   make(map[pageID]viewport.Model),
	}, nil
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		loadReadmeContent(), // Command to load README content
	)
}

// Message type for README content
type scrollContentMsg struct {
	content string
	err     error
}

// Command to load README content
func loadReadmeContent() tea.Cmd {
	return func() tea.Msg {
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
				slog.Debug("README found", "path", path)
				break
			}
		}

		if err != nil {
			return scrollContentMsg{
				err: err,
			}
		}

		out, err := glamour.Render(string(content), "dark")
		if err != nil {
			slog.Error("Error rendering README", "error", err, "path", foundPath)
			return scrollContentMsg{
				err: err,
			}
		}

		return scrollContentMsg{
			content: out,
			err:     nil,
		}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	slog.Debug("Update", "msg", msg)

	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case scrollContentMsg:
		// Initialize viewport with README content when it's loaded
		slog.Debug("README content loaded", "error", msg.err)
		if msg.err == nil {
			vp := viewport.New(m.widthContent, m.heightContent)
			vp.SetContent(msg.content)
			m.viewports[readmePage] = vp
		}

	case tea.WindowSizeMsg:
		m.viewportWidth = msg.Width
		m.viewportHeight = msg.Height

		switch {
		case m.viewportWidth < 20 || m.viewportHeight < 10:
			m.size = undersized
			m.widthContainer = m.viewportWidth
			m.heightContainer = m.viewportHeight
		case m.viewportWidth < 50:
			m.size = small
			m.widthContainer = m.viewportWidth
			m.heightContainer = m.viewportHeight
		case m.viewportWidth < 80:
			m.size = medium
			m.widthContainer = 50
			m.heightContainer = int(math.Min(float64(msg.Height), 30))
		default:
			m.size = large
			m.widthContainer = 80
			m.heightContainer = int(math.Min(float64(msg.Height), 30))
		}

		m.widthContent = m.widthContainer - 2
		m.heightContent = m.heightContainer - lipgloss.Height(m.HeaderView()) - lipgloss.Height(m.FooterView()) - 2

		// Update all viewport sizes
		for pageID, vp := range m.viewports {
			vp.Width = m.widthContent
			vp.Height = m.heightContent
			m.viewports[pageID] = vp
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			// Back to menu
			m.currentPage = menuPage
			m.cursor = 0
			return m, nil
		}
	}

	var cmd tea.Cmd
	switch m.currentPage {
	case menuPage:
		m, cmd = m.MenuUpdate(msg)
	case aboutMe:
		m, cmd = m.AboutMeUpdate(msg)
	case page2:
		m, cmd = m.PageUpdate(page2, msg)
	case page3:
		m, cmd = m.PageUpdate(page3, msg)
	case readmePage:
		m, cmd = m.ReadmeUpdate(msg)
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	items := []string{}
	header := m.HeaderView()
	items = append(items, header)

	switch m.currentPage {
	case menuPage:
		items = append(items, m.MenuView())
	case aboutMe, page2, page3, readmePage:
		items = append(items, m.PageView())
	}

	footer := m.FooterView()

	items = append(items, footer)

	child := lipgloss.JoinVertical(
		lipgloss.Left,
		items...,
	)
	return m.renderer.Place(
		m.viewportWidth,
		m.viewportHeight,
		lipgloss.Center,
		lipgloss.Center,
		m.theme.Base().
			MaxWidth(m.widthContainer).
			MaxHeight(m.heightContainer).
			Render(child),
	)
}
