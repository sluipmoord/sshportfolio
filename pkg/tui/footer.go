package tui

import (
	"fmt"
	"log/slog"

	"github.com/charmbracelet/lipgloss"
	"github.com/sluipmoord/sshportfolio/pkg/assets"
)

// FooterView renders the footer of the application with navigation hints
func (m model) FooterView() string {
	bold := m.theme.TextAccent().Bold(true).Render
	base := m.theme.Base().Render

	esc := bold("q")
	back := bold("esc")
	navHints := base(fmt.Sprintf("%s back | %s quit | ", back, esc))

	githubLink := base("https://github.com/sluipmoord")
	githubIcon, err := assets.SVGToASCII("github", 1, 1)
	if err != nil {
		slog.Error("failed to load github icon", "error", err)
		githubIcon = "[GH]"
	}

	github := base(fmt.Sprintf("%s %s", bold(githubIcon), githubLink))

	footer := lipgloss.JoinHorizontal(lipgloss.Left, navHints, github)

	return m.theme.Base().PaddingTop(1).Render(footer)
}
