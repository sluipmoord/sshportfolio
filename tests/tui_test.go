package tests

import (
	"sshportfolio/pkg/tui"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestTUIModel(t *testing.T) {
	projects := []string{"Project 1", "Project 2", "Project 3"}
	model := tui.Model{Projects: projects, Cursor: 0}

	// Test initial state
	if model.Cursor != 0 {
		t.Errorf("Expected cursor to be 0, got %d", model.Cursor)
	}

	// Test moving down
	updatedModel, _ := model.Update(tea.KeyMsg{Type: tea.KeyDown})
	model = updatedModel.(tui.Model)
	if model.Cursor != 1 {
		t.Errorf("Expected cursor to be 1, got %d", model.Cursor)
	}

	updatedModel, _ = model.Update(tea.KeyMsg{Type: tea.KeyUp})
	model = updatedModel.(tui.Model)
	model = updatedModel.(tui.Model)
	if model.Cursor != 0 {
		t.Errorf("Expected cursor to be 0, got %d", model.Cursor)
	}
}
