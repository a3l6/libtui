package libtui

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

// This function duplicates the given 1D input slice into a 2D output with contains n number of times.
// NOTE: the output of this function should not be mutated. If one thing is mutated the mutation propogates across entire array.
// NOTE: This occurs due to the function copying the slice header not the underlying array.
// I will not change this behaviour. This works for my use case.
func DuplicateRuneSlice(input []rune, n int) [][]rune {
	result := make([][]rune, n)
	for i := range n {
		result[i] = input
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
	ArrRunes  RenderTarget = iota
	ArrString              // Currently only runes is supported
)

type Overflow uint8

const (
	Visible Overflow = iota
	Hidden
	Scroll
)

type LibtuiObject interface {
	RenderToArrRunes() ([]rune, error)
}
