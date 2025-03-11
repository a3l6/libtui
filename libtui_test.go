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
