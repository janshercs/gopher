package main

import "testing"

func TestMain(t *testing.T) {
	testset := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "base case", input: "aba", want: true},
		{name: "base case", input: "notapalindrome", want: false},
		{name: "base case", input: "aaa", want: true},
		{name: "base case", input: "aaba", want: false},
		{name: "base case", input: "aaaa", want: true},
		{name: "string conversion", input: "Aaaa", want: true},
		{name: "strip space", input: "Aaa a", want: true},
		{name: "strip symbols", input: "a$#@Aaaa", want: true},
		{name: "leaves numbers", input: "1$#@Aaa1", want: true},
		{name: "leaves numbers", input: "0p", want: false},
	}

	checkAnswer := func(t testing.TB, input string, want bool) {
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
