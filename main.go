package libtui

import (
	"errors"
	"fmt"
	"strings"
)

func SplitIntoChunks(s string, size int) []string {
	var result []string

	for i := 0; i < len(s); i += size {
		end := min(i+size, len(s))
		result = append(result, s[i:end])
	}

	return result
}

type RecoverableError struct {
	Field string
}

func (e RecoverableError) Error() string {
	return e.Field + " But this error is recoverable, a valid output was given."
}

type Alignment uint8

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
)

type Button struct {
	highlighted bool
	Width       uint8
	Height      uint8
	Align       Alignment
	Value       string
	callback    func()
}

// Renders out button as array of strings. Returns errors, some are recoverable and valid outputs are given.
func (btn *Button) RenderToArrRunes() ([]rune, error) {
	if btn.Height == 1 {
		len_val := len(btn.Value)

		var overflowError error

		if len_val > int(btn.Width) {
			btn.Value = btn.Value[:btn.Width]
			len_val = int(btn.Width)
			overflowError = RecoverableError{Field: "button.Value overflowed btn.width."}
		}

		base := []rune("[" + strings.Repeat(" ", int(btn.Width)) + "]")
		offset := 1

		switch btn.Align {
		case AlignCenter:
			offset = (int(btn.Width)+2)/len_val - (len_val / 2)
			if offset+len_val >= int(btn.Width) {
				overflowError = RecoverableError{Field: "button.Value overflowed btn.width with centre alignment."}
			}
		case AlignLeft:
			if len_val >= int(btn.Width) {
				overflowError = RecoverableError{Field: "button.Value overflowed btn.width with left alignment"}
			}
		case AlignRight:
			offset = int(btn.Width) - len_val + 1
			if offset+len_val >= int(btn.Width) {
				overflowError = RecoverableError{Field: "button.Value overflowed btn.width with right alignment."}
			}
		default:
			return []rune{}, errors.New("alignment not supported, please choose between AlignLeft, AlignCenter, AlignRight")
		}

		copy(base[offset:], []rune(btn.Value))
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
