package libtui

import (
	"errors"
	"fmt"
	"strings"
)

type Modal struct {
	Width    int
	Height   int
	Position Vector2
	Active   bool
	Elems    []LibtuiObject
}

// Function calls `RenderToArrRunes()` on all elements in the modals scope and returns the result and any error.
// Errors are returned as they are found.
func (modal *Modal) RenderElems() ([][]rune, error) {
	output := [][]rune{}

	for _, val := range modal.Elems {
		rendered, err := val.RenderToArrRunes()
		if err != nil {
			return [][]rune{}, err
		}
		output = append(output, rendered)
	}
	return output, nil
}

// Function renders just the outline of the modal. For rendering of elements call `Modal.RenderElems`.
// Function returns errors, and does not panic on any.
// Modal lines are rendered using ANSI codes to move cursor
func (modal *Modal) RenderToArrRunes() ([][]rune, error) {
	if modal.Width < 2 {
		return [][]rune{}, errors.New("Width must be at least 2 to draw the outline of the modal")
	}

	base := []rune(strings.Repeat(" ", modal.Width))

	top_box := []rune("┌" + strings.Repeat("─", modal.Width-2) + "┐")
	bottom_box := []rune("└" + strings.Repeat("─", modal.Width-2) + "┘")

	if modal.Height == 2 {
		output := append(top_box, bottom_box...)
		return [][]rune{output}, nil
	}

	for i := modal.Width; i < len(base)-modal.Width; i += modal.Width {
		base[i] = '│'
		base[i+modal.Width-1] = '│'
	}

	output := make([][]rune, 0)
	base_output := DuplicateRuneSlice(base, modal.Width)
	return base, nil
}
