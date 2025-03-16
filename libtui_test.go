package libtui

import (
	"errors"
	"reflect"
	"testing"
)

func Test_ButtonRenderArrRunes(t *testing.T) {

	var tests = []struct {
		name  string
		input Button
		want  []rune
	}{
		{"Base case of button with center alignment", Button{Width: 10, Height: 1, Value: "Hi", Align: AlignCenter}, []rune("[   Hi   ]")},
		{"Base case of button with left alignment", Button{Width: 10, Height: 1, Value: "Hi", Align: AlignLeft}, []rune("[Hi      ]")},
		{"Base case of button with right alignment", Button{Width: 10, Height: 1, Value: "Hi", Align: AlignRight}, []rune("[      Hi]")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.input.RenderToArrRunes()
			if !reflect.DeepEqual(tt.want, got) || err != nil {
				t.Errorf("Button.RenderToArrRunes() = %v, %v, want match for %v", string(got), err, string(tt.want))
			}
		})
	}
}

func Test_ButtonRenderArrRunesErrors(t *testing.T) {

	var tests = []struct {
		name  string
		input Button
		want  []rune
		err   error
	}{
		{
			"Test general overflow error",
			Button{Width: 10, Height: 1, Value: "Hello World", Align: AlignLeft},
			[]rune("[Hello Wo]"),
			RecoverableError{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.input.RenderToArrRunes()
			if !reflect.DeepEqual(tt.want, got) || errors.Is(err, tt.err) {
				t.Errorf("Button.RenderToArrRunes() = %v, %v, want match for %v, %v", string(got), err, string(tt.want), tt.err)
			}
		})
	}
}

func Test_SplitIntoChunks(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		size  int
		want  []string
	}{
		{
			"Testing General Split",
			"Lorem ipsum dolor sit amet. Ad beatae quibusdam in voluptas sint aut harum voluptas ut fugiat voluptatem sed consequuntur quis. Sed explicabo esse et iure debitis et eveniet architecto aut voluptas dolores qui eaque assumenda non incidunt assumenda id voluptate ullam.",
			50,
			[]string{"Lorem ipsum dolor sit amet. Ad beatae quibusdam in", " voluptas sint aut harum voluptas ut fugiat volupt", "atem sed consequuntur quis. Sed explicabo esse et ", "iure debitis et eveniet architecto aut voluptas do", "lores qui eaque assumenda non incidunt assumenda i", "d voluptate ullam."},
		},
		{
			"Base Case, no splitting needed",
			"Hello World",
			11,
			[]string{"Hello World"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := SplitIntoChunks(tt.input, tt.size)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("SplitIntoChunks() = \r\n%#v, want match for \r\n%#v", got, tt.want)
			}
		})
	}
}
