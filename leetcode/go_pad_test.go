package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	testset := []struct {
		name  string
		input [][]int
		want  int
	}{
		{name: "base case",
			input: [][]int{
				{1, 5, 3},
				{1, 5, 1},
				{6, 6, 5},
			},
			want: 8},
		{name: "base case",
			input: [][]int{
				{1, 3, 2},
				{1, 5, 2},
				{2, 4, 5},
			},
			want: 5},
	}

	checkAnswer := func(t testing.TB, input [][]int, want int) {
		t.Helper()
		got := driver(input)
		if got != want {
			t.Errorf("test failed, wanted %d got %d", want, got)
		}
	}

	for _, test := range testset {
		t.Run(test.name, func(t *testing.T) {
			checkAnswer(t, test.input, test.want)
		})
	}

}
