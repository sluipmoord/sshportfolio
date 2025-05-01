package tui

import (
	"context"
	"fmt"
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
	menu page = iota
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

	page  page
	pages []string

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
		renderer: renderer,
		context:  ctx,
		command:  command,
		cursor:   0,
		page:     menu,
		pages:    []string{"Page 1", "Page 2", "Page 3"},
		theme:    theme.BasicTheme(renderer, nil),
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
		m.heightContent = m.heightContainer - lipgloss.Height(m.HeaderView()) - lipgloss.Height(m.FooterView()) - 2
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			// Navigate back to the main page list
			m.page = menu
			m.cursor = 0
			return m, nil
		}
	}

	var cmd tea.Cmd
	switch m.page {
	case menu:
		m, cmd = m.MenuUpdate(msg)
	}

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	items := []string{}
	header := m.HeaderView()
	items = append(items, header)
	switch m.page {
	case menu:
		items = append(items, m.MenuView())
	case page1, page2, page3:
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

func (m model) HeaderView() string {
	bold := m.theme.TextAccent().Bold(true).Render
	base := m.theme.Base().Render

	header := base("My Portfolio Page:")
	header += bold(fmt.Sprintf(" %d", m.page))

	return m.theme.Base().PaddingBottom(1).Render(header)
}

func (m model) FooterView() string {
	bold := m.theme.TextAccent().Bold(true).Render
	base := m.theme.Base().Render

	esc := bold("q")
	back := bold("esc")
	footer := base(fmt.Sprintf("%s back | %s quit", back, esc))

	return m.theme.Base().PaddingTop(1).Render(footer)
}

func (m model) MenuView() string {
	bold := m.theme.TextAccent().Bold(true).Render
	render := m.theme.Base().Render

	var renderedMenu []string
	for i, item := range m.pages {
		cursor := " " // no cursor
		if i == m.cursor {
			cursor = ">" // cursor for the selected item
			item = bold(item)
		}
		renderedMenu = append(renderedMenu, render(fmt.Sprintf("%s %s", cursor, item)))
	}

	return lipgloss.JoinVertical(lipgloss.Left, renderedMenu...)
}

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
			switch m.cursor {
			case 0:
				m.page = page1
			case 1:
				m.page = page2
			case 2:
				m.page = page3
			}
		}
	}

	return m, nil
}

func (m model) PageView() string {

	page := ""
	switch m.page {
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
func (m model) Page1View() string {
	return m.theme.Base().Render("This is the content of Page 1")
}
func (m model) Page2View() string {
	return m.theme.Base().Render("This is the content of Page 2")
}
func (m model) Page3View() string {
	return m.theme.Base().Render("This is the content of Page 3")
}
