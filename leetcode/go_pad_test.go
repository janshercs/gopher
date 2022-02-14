package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	testset := []struct {
		name  string
		input LCInput
		want  int
	}{
		{name: "base case",
			input: LCInput{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}},
			want:  11,
		},
		{name: "min case",
			input: LCInput{{2}},
			want:  2,
		},
		{name: "min case",
			input: LCInput{{2}, {-10, 1}, {-5, 7, 8}},
			want:  -13,
		},
	}

	checkAnswer := func(t testing.TB, input LCInput, want int) {
		t.Helper()
		got := driver(input)

		if got != want {
			t.Errorf("test failed, wanted %v got %v", want, got)

		}

		// if !reflect.DeepEqual(got, want) {
		// 	t.Errorf("test failed, wanted %v got %v", want, got)
		// }
	}

	for _, test := range testset {
		t.Run(test.name, func(t *testing.T) {
			checkAnswer(t, test.input, test.want)
		})
	}

}
