package tests

import (
	"os"
	"testing"

	"github.com/sluipmoord/sshportfolio/pkg/tui"

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

	// Instead of running the program with a TTY, test the model directly
	// Test initial view
	initialView := model.View()
	if initialView == "" {
		t.Errorf("Initial view should not be empty")
	}

	// Test model update
	cmd := model.Init()
	if cmd != nil {
		actions := cmd()
		if actions != nil {
			model, _ = model.Update(actions)
		}
	}

	// Verify the model's state after initialization
	updatedView := model.View()
	if updatedView == "" {
		t.Errorf("Updated view should not be empty")
	}

	// Note: We're not calling program.Run() since it requires a TTY
	// which is not available in test environments
}
