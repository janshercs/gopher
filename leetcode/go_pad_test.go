package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	testset := []struct {
		name  string
		input [][]byte
		want  int
	}{
		{name: "base case", input: [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}, want: 1},
		{name: "base case", input: [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}, want: 3},
		{name: "1 col", input: [][]byte{
			{'1'},
			{'1'},
			{'0'},
			{'0'},
		}, want: 1},
		{name: "1 row", input: [][]byte{
			{'1', '1', '0', '0', '1'},
		}, want: 2},
	}

	checkAnswer := func(t testing.TB, input [][]byte, want int) {
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
