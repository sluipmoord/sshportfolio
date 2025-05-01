package tui

import (
	"context"
	"log/slog"
	"math"
	"sshportfolio/pkg/tui/theme"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type page = int
type cursor = int
type size = int

const (
	menuPage page = iota
	page1
	page2
	page3
)

const (
	undersized size = iota
	small
	medium
	large
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

	currentPage page
	pages       []string

	cursor cursor
}

func NewModel(
	renderer *lipgloss.Renderer,
	clientIP *string,
	command []string,
) (tea.Model, error) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "client_ip", clientIP)

	return model{
		renderer:    renderer,
		context:     ctx,
		command:     command,
		cursor:      0,
		currentPage: menuPage,
		pages:       []string{"Page 1", "Page 2", "Page 3"},
		theme:       theme.BasicTheme(renderer, nil),
	}, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	slog.Info("Update", "msg", msg)

	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
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
		m.heightContent = m.heightContainer - lipgloss.Height(HeaderView(m)) - lipgloss.Height(FooterView(m)) - 2
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			//  back to menu
			m.currentPage = menuPage
			m.cursor = 0
			return m, nil
		}
	}

	var cmd tea.Cmd
	switch m.currentPage {
	case menuPage:
		m, cmd = MenuUpdate(m, msg)
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	items := []string{}
	header := HeaderView(m)
	items = append(items, header)

	switch m.currentPage {
	case menuPage:
		items = append(items, MenuView(m))
	case page1, page2, page3:
		items = append(items, PageView(m))
	}

	footer := FooterView(m)

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
