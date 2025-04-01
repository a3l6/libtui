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
	Elems    []libtuiObject
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
func (modal *Modal) RenderToArrRunes() ([]rune, error) {
	if modal.Width < 2 {
		return []rune{}, errors.New("Width must be at least 2 to draw the outline of the modal")
	}

	base := []rune(strings.Repeat(" ", modal.Width*modal.Height))

	top_box := []rune("┌" + strings.Repeat("─", modal.Width-2) + "┐")
	bottom_box := []rune("└" + strings.Repeat("─", modal.Width-2) + "┘")

	copy(base[:], top_box[:])
	copy(base[modal.Width*(modal.Height-1):], bottom_box[:])

	if modal.Height == 2 {
		base, err := insertEveryN(base, []rune(fmt.Sprintf("\033[1B\033[%dD", modal.Width)), modal.Width)
		if err != nil {
			return []rune{}, errors.Join(errors.New("Could not insert ANSI codes into rende3. Encountered error: "), err)
		}
		return base, nil
	}

	for i := modal.Width; i < len(base)-modal.Width; i += modal.Width {
		base[i] = '│'
		base[i+modal.Width-1] = '│'
	}

	base, err := insertEveryN(base, []rune(fmt.Sprintf("\033[1B\033[%dD", modal.Width)), modal.Width)
	if err != nil {
		return []rune{}, errors.Join(errors.New("Could not insert ANSI codes into render. Encountered error: "), err)
	}

	return base, nil
}
