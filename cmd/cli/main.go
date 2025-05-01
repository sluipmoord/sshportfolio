package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/sluipmoord/sshportfolio/pkg/config"
	"github.com/sluipmoord/sshportfolio/pkg/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Setup logger with configured log level
	if err := config.SetupLogger(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set up logger: %v\n", err)
		os.Exit(1)
	}

	slog.Info("starting cli application", "logLevel", cfg.LogLevel)

	model, err := tui.NewModel(lipgloss.DefaultRenderer(), nil, []string{})
	if err != nil {
		slog.Error("Failed to create TUI model", "error", err)
		os.Exit(1)
	}

	if _, err := tea.NewProgram(model, tea.WithAltScreen()).Run(); err != nil {
		slog.Error("Error running program", "error", err)
		os.Exit(1)
	}
}
