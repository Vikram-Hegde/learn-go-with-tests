package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assert(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d wanted %d", got, want)
	}
}

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4
	assert(t, got, want)
}