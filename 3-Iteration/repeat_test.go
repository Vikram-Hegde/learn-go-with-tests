package iteration

import (
	"fmt"
	"testing"
)

func assert(t testing.TB, repeated, expected string) {
	t.Helper()

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 10)
	fmt.Println(repeated)
	// Output: aaaaaaaaaa
}

func TestRepeat(t *testing.T) {
	t.Run("repeating character with amount", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assert(t, repeated, expected)
	})

	t.Run("repeating character with 0 amount", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""

		assert(t, repeated, expected)
	})
}
