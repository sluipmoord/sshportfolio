package theme

import (
	"github.com/charmbracelet/huh"
)

// copy returns a copy of a TextInputStyles with all children styles copied.
func copyTextStyles(t huh.TextInputStyles) huh.TextInputStyles {
	return huh.TextInputStyles{
		Cursor:      t.Cursor,
		Placeholder: t.Placeholder,
		Prompt:      t.Prompt,
		Text:        t.Text,
	}
}

// copy returns a copy of a FieldStyles with all children styles copied.
func copyFieldStyles(f huh.FieldStyles) huh.FieldStyles {
	return huh.FieldStyles{
		Base:           f.Base,
		Title:          f.Title,
		Description:    f.Description,
		ErrorIndicator: f.ErrorIndicator,
		ErrorMessage:   f.ErrorMessage,
		SelectSelector: f.SelectSelector,
		// NextIndicator:       f.NextIndicator,
		// PrevIndicator:       f.PrevIndicator,
		Option: f.Option,
		// Directory:           f.Directory,
		// File:                f.File,
		MultiSelectSelector: f.MultiSelectSelector,
		SelectedOption:      f.SelectedOption,
		SelectedPrefix:      f.SelectedPrefix,
		UnselectedOption:    f.UnselectedOption,
		UnselectedPrefix:    f.UnselectedPrefix,
		FocusedButton:       f.FocusedButton,
		BlurredButton:       f.BlurredButton,
		TextInput:           copyTextStyles(f.TextInput),
		Card:                f.Card,
		NoteTitle:           f.NoteTitle,
		Next:                f.Next,
	}
}
