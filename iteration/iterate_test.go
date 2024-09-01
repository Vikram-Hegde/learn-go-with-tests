package iteration

import "testing"

func TestIterate(t *testing.T) {
	expected := "vvvvv"
	got := Repeat("v", 5)
	assert(t, got, expected)
}

func BenchmarkIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}