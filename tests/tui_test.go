package tests

import (
	"sshportfolio/pkg/tui"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestTUIModel(t *testing.T) {
	projects := []string{"Project 1", "Project 2", "Project 3"}
	model := tui.Portfolio{Pages: projects, Cursor: 0}

	// Test initial state
	if model.Cursor != 0 {
		t.Errorf("Expected cursor to be 0, got %d", model.Cursor)
	}

	// Test moving down
	updatedModel, _ := model.Update(tea.KeyMsg{Type: tea.KeyDown})
	model = updatedModel.(tui.Portfolio)
	if model.Cursor != 1 {
		t.Errorf("Expected cursor to be 1, got %d", model.Cursor)
	}

	updatedModel, _ = model.Update(tea.KeyMsg{Type: tea.KeyUp})
	model = updatedModel.(tui.Portfolio)
	model = updatedModel.(tui.Portfolio)
	if model.Cursor != 0 {
		t.Errorf("Expected cursor to be 0, got %d", model.Cursor)
	}
}
