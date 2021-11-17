package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	testset := []struct {
		name  string
		input struct {
			arr    []int
			target int
		}
		want []int
	}{
		{name: "base case",
			input: LCInput{
				arr:    []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			want: []int{3, 4}},
		{name: "negative case",
			input: LCInput{
				arr:    []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			want: []int{-1, -1}},
		{name: "single 0 index case",
			input: LCInput{
				arr:    []int{5, 7, 7, 8, 8, 10},
				target: 5,
			},
			want: []int{0, 0}},
		{name: "single len-1 case",
			input: LCInput{
				arr:    []int{5, 7, 7, 8, 8, 10},
				target: 10,
			},
			want: []int{5, 5}},
		{name: "end at len-1 case",
			input: LCInput{
				arr:    []int{5, 7, 7, 8, 8, 8},
				target: 8,
			},
			want: []int{3, 5}},
		{name: "start at 0 case",
			input: LCInput{
				arr:    []int{5, 5, 7, 8, 8, 8},
				target: 5,
			},
			want: []int{0, 1}},
		{name: "null case",
			input: LCInput{
				arr:    []int{},
				target: 6,
			},
			want: []int{-1, -1}},
	}

	checkAnswer := func(t testing.TB, input LCInput, want []int) {
		t.Helper()
		got := driver(input)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("test failed, wanted %v got %v", want, got)
		}
	}

	for _, test := range testset {
		t.Run(test.name, func(t *testing.T) {
			checkAnswer(t, test.input, test.want)
		})
	}

}
