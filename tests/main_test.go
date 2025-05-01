package tests

import (
	"os"
	"sshportfolio/pkg/tui"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func TestMainProgram(t *testing.T) {
	logFile, err := os.CreateTemp("", "output.log")
	if err != nil {
		t.Fatalf("Failed to create temp log file: %v", err)
	}
	defer os.Remove(logFile.Name())

	model, err := tui.NewModel(lipgloss.DefaultRenderer(), nil, []string{})
	if err != nil {
		t.Fatalf("Failed to create TUI model: %v", err)
	}

	program := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		t.Errorf("Program run failed: %v", err)
	}
}
