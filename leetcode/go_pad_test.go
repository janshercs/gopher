package main

import "testing"

func TestMain(t *testing.T) {
	testset := []struct {
		name  string
		input []int
		want  bool
	}{
		{name: "base case", input: []int{1, 2, 3, 1}, want: true},
		{name: "base case", input: []int{1, 2, 3}, want: false},
		{name: "one element case", input: []int{1}, want: false},
		{name: "negative element false case", input: []int{-1, -2}, want: false},
		{name: "negative element true case", input: []int{-1, -2, -1}, want: true},
		{name: "negative & positive element opposite case", input: []int{-1, -2, 1}, want: false},
	}

	checkAnswer := func(t testing.TB, input []int, want bool) {
		t.Helper()
		got := driver(input)
		if want != got {
			t.Errorf("test failed, wanted %v got %v", want, got)
		}
	}

	for _, test := range testset {
		t.Run(test.name, func(t *testing.T) {
			checkAnswer(t, test.input, test.want)
		})
	}

}
