package iteration // this is the package that is imported

import "fmt"
import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectString := func(t testing.TB, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	}

	t.Run("repeat single character", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertCorrectString(t, repeated, expected)
	})

	t.Run("repeat multiple character", func(t *testing.T) {
		repeated := Repeat("abc", 5)
		expected := "abcabcabcabcabc"

		assertCorrectString(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 6)
	}
}

func ExampleRepeat() {
	repeated := Repeat("Hi ", 3)
	fmt.Println(repeated)
	// Output: Hi Hi Hi
}
