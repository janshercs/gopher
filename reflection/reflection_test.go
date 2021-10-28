package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name:          "Struct with 1 string field",
			Input:         struct{ Name string }{"Chris"},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Struct with 2 string field",
			Input: struct {
				Name   string
				Gender string
			}{"Chris", "Male"},
			ExpectedCalls: []string{"Chris", "Male"},
		},
		{
			Name: "Struct with string and int field",
			Input: struct {
				Name string
				Age  int
			}{"Chris", 32},
			ExpectedCalls: []string{"Chris"},
		},
		{
			Name: "Nested structs/fields",
			Input: Person{
				"Chris",
				Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "pass in pointers",
			Input: &Person{
				"Chris",
				Profile{33, "London"},
			},
			ExpectedCalls: []string{"Chris", "London"},
		},
		{
			Name: "pass in slices",
			Input: []Profile{
				{32, "Texas"},
				{33, "London"},
			},
			ExpectedCalls: []string{"Texas", "London"},
		},
		{
			Name: "pass in arrays",
			Input: [2]Profile{
				{32, "Texas"},
				{33, "London"},
			},
			ExpectedCalls: []string{"Texas", "London"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if len(got) != len(test.ExpectedCalls) {
				t.Errorf("wrong number of function calls, got %d want %d", len(got), len(test.ExpectedCalls))
			}

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("wrong contents, got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

	assertContains := func(t testing.TB, haystack []string, needle string) {
		t.Helper()
		contains := false
		for _, x := range haystack {
			if x == needle {
				contains = true
			}
		}
		if !contains {
			t.Errorf("expected %+v to contain %q but it did not", haystack, needle)
		}
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"foo":  "Texas",
			"bish": "London",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Texas")
		assertContains(t, got, "London")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{33, "Texas"}
			aChannel <- Profile{33, "London"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Texas", "London"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wrong contents, got %v want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Texas"}, Profile{33, "London"}
		}

		var got []string
		want := []string{"Texas", "London"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wrong contents, got %v want %v", got, want)
		}
	})

}
