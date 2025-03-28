package libtui

import (
	"errors"
	"fmt"
	"strings"
)

type Button struct {
	Width       uint8
	Height      uint8
	Position    Vector2
	Align       Alignment
	Value       string
	Callback    func()
	Key         rune // No Internal Use
	highlighted bool
}

func (btn *Button) FocusOn() {
	btn.highlighted = true
}

func (btn *Button) FocusOff() {
	btn.highlighted = false
}

func (btn *Button) GetFocus() bool {
	return btn.highlighted
}

// Renders out button as array of strings. Returns errors, some are recoverable and valid outputs are given.
func (btn *Button) RenderToArrRunes() ([]rune, error) {
	if btn.Height == 1 {
		len_val := len(btn.Value)

		var overflowError error

		if len_val-1 > int(btn.Width)-2 {
			btn.Value = btn.Value[:btn.Width-2]
			len_val = int(btn.Width) - 2
			overflowError = RecoverableError{Field: "button.Value overflowed btn.width."}
		}

		var base []rune
		if btn.highlighted {
			base = []rune("[\033[7m" + strings.Repeat(" ", int(btn.Width-2)) + "\033[0m]")
		} else {
			base = []rune("[" + strings.Repeat(" ", int(btn.Width-2)) + "]")
		}

		var offset int
		switch btn.Align {
		case AlignCenter:
			offset = (int(btn.Width) - len_val) / 2
			if offset+len_val >= int(btn.Width) {
				overflowError = RecoverableError{Field: "button.Value overflowed btn.width with centre alignment."}
			}
		case AlignLeft:

			offset = 1
			if len_val >= int(btn.Width) {
				overflowError = RecoverableError{Field: "button.Value overflowed btn.width with left alignment"}
			}
		case AlignRight:
			offset = int(btn.Width) - len_val - 1
			if offset+len_val >= int(btn.Width) {
				overflowError = RecoverableError{Field: "button.Value overflowed btn.width with right alignment."}
			}
		default:
			return []rune{}, fmt.Errorf("alignment not supported, please choose between AlignLeft, AlignCenter, AlignRight: %w", errors.ErrUnsupported)
		}

		if btn.highlighted {
			// 4 for the initial \033[7m
			copy(base[offset+4:], []rune(btn.Value))
		} else {
			copy(base[offset:], []rune(btn.Value))
		}
		return base, overflowError

	} else {
		return []rune{}, fmt.Errorf("Currently do not support multiline buttons. %w", errors.ErrUnsupported)
		/*var errors_encountered error

		base := make([][]rune, btn.height)

		chunkedString := SplitIntoChunks(btn.Value, int(btn.width)-2)
		if int(btn.height) < len(chunkedString) {
			errors_encountered = RecoverableError{Field: "button.Value overflowed the allotted buffer."}
		}
		switch btn.align {
		case AlignCenter:
			for i := 0; i >= min(len(chunkedString), int(btn.height)); i++ {
				offset := (int(btn.width)+2)/len(chunkedString[i]) - (len(chunkedString[i]) / 2)
				base[i] = []rune("[" + strings.Repeat(" ", int(btn.width)) + "]")
				copy(base[i][offset:], []rune(chunkedString[i]))
			}
		case AlignLeft:
		case AlignRight:
		default:

		}

		return [][]rune{}, errors_encountered*/
	}
}
