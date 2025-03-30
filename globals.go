package libtui

import "errors"

type Vector2 struct {
	X int
	Y int
}

func SplitIntoChunks(s string, size int) []string {
	var result []string
	for i := 0; i < len(s); i += size {
		end := min(i+size, len(s))
		result = append(result, s[i:end])
	}
	return result
}

func SplitArrRunesIntoChunks(s []rune, size int) [][]rune {
	var result [][]rune
	for i := 0; i < len(s); i += size {
		end := min(i+size, len(s))
		result = append(result, s[i:end])
	}
	return result
}

func joinArrayArrayRunes(elems [][]rune, sep []rune) []rune {
	var result []rune
	for _, val := range elems {
		result = append(result, val...)
		result = append(result, sep...)
	}
	return result
}

func insertEveryN(source, insert []rune, n int) ([]rune, error) {
	if n <= 0 {
		errors.New("N must be greater than zero")
	}

	insertions := (len(source) - 1) / n
	if len(source)%n == 0 {
		insertions++
	}

	resultLen := len(source) + insertions*len(insert)
	result := make([]rune, resultLen)

	srcIdx, resIdx := 0, 0
	for i := 0; srcIdx < len(source); i++ {
		chunkEnd := srcIdx + n
		if chunkEnd > len(source) {
			chunkEnd = len(source)
		}
		copy(result[resIdx:], source[srcIdx:chunkEnd])
		resIdx += chunkEnd - srcIdx
		srcIdx = chunkEnd

		if srcIdx < len(source) {
			copy(result[resIdx:], insert)
			resIdx += len(insert)
		}
	}

	return result, nil
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
	ArrRunes  RenderTarget = iota
	ArrString              // Currently only runes is supported
)

type Overflow uint8

const (
	Visible Overflow = iota
	Hidden
	Scroll
)

type libtuiObject interface {
	RenderToArrRunes() ([]rune, error)
}
