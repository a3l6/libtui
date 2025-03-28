package libtui

import (
	"errors"
	"fmt"
	"strings"
)

type Text struct {
	Width       uint16
	Height      uint16
	Position    Vector2
	Align       Alignment
	Value       string
	YOverflow   Overflow
	highlighted bool
}

func (text *Text) FocusOn() {
	text.highlighted = true
}

func (text *Text) FocusOff() {
	text.highlighted = false
}

func (text *Text) GetFocus() bool {
	return text.highlighted
}

func (text *Text) RenderToArrRunes() ([]rune, error) {
	lines := SplitArrRunesIntoChunks([]rune(text.Value), int(text.Width))
	var lines2 [][]rune

	// Instantiate blank lines
	for range lines {
		blank := []rune(strings.Repeat(" ", int(text.Width)))
		lines2 = append(lines2, blank)
	}

	switch text.Align {
	case AlignCenter:
		for idx, val := range lines {
			offset := (int(text.Width) - len(val)) / 2
			copy(lines2[idx][offset:], val[:])
		}
	case AlignLeft:
		for idx, val := range lines {
			copy(lines2[idx][:], val)
		}
	case AlignRight:
		for idx, val := range lines {
			offset := int(text.Width) - len(val)
			copy(lines2[idx][offset:], val)
		}
	default:
		return []rune{}, fmt.Errorf("Invalid alignment: %w", errors.ErrUnsupported)
	}

	var result []rune
	for _, val := range lines2 {
		sep := []rune(fmt.Sprintf("\033[%dD\033[1B", text.Width))
		result = append(result, val...)
		result = append(result, sep...)
	}
	result = append(result, []rune("\n")...) // do this to avoid a weird % in bash

	return result, nil
}
