package libtui

import (
	"errors"
	"fmt"
	"strings"
)

type FocusableObject interface {
	GetFocus()
	focusOn()
	focusOff()
}

type FocusNode struct {
	object FocusableObject
	Next   *FocusNode
	Prev   *FocusNode
}

// The focus list holds the pointers to all focusable elements.
// Only one element may be focused at a time.
// Focus is implemented with a linked list.
type Focus struct {
	size int
	Head *FocusNode
	Tail *FocusNode
	Curr *FocusNode
}

// Fuction takes target and node of type *FocusNode.
// Target is the node being used as the reference point.
// node is the node being inserted.
// Function returns errors if any required pointers are nil.
// Required pointer is target.
// Passing node as nil is valid. Will cut off the rest of the linked list.
func (foc *Focus) InsertAfterNode(target *FocusNode, node *FocusNode) error {
	if target == nil {
		return errors.New("target cannot be nil")
	}
	if node == nil {
		foc.Tail = target
		target.Next = nil
		return nil
	}

	if target == foc.Tail {
		foc.Tail = node
	}
	node.Next = target.Next
	target.Next = node
	node.Prev = target

	foc.size++

	return nil
}

func (foc *Focus) GetSize() int {
	// Not sure if the size will be useful
	// But I want it to only be mutated by the list
	return foc.size
}

// Move focus to the next element in the array.
// If the focus is at the tail then nothing happens.
func (foc *Focus) Next() {
	if foc.Curr.Next != nil {
		foc.Curr.object.focusOff()
		foc.Curr = foc.Curr.Next
		foc.Curr.object.focusOn()
	}
}

// Moves focus to previous element.
// If the focus is at the head then nothing happens.
func (foc *Focus) Prev() {
	if foc.Curr.Prev != nil {
		foc.Curr.object.focusOff()
		foc.Curr = foc.Curr.Prev
		foc.Curr.object.focusOn()
	}
}

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

type RenderTarget uint8

const (
	ArrRunes RenderTarget = iota
	ArrString
)

type Button struct {
	highlighted  bool
	Width        uint8
	Height       uint8
	Align        Alignment
	Value        string
	Callback     func()
	RenderTarget RenderTarget // Not implemented yet
	focus        bool
}

func (btn *Button) render() {

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

		base := []rune("[" + strings.Repeat(" ", int(btn.Width-2)) + "]")
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
