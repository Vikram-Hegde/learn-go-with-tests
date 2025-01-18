package main

import (
	"testing"
)

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Vikram", "")
		want := "Hello, Vikram"

		assert(t, got, want)
	})

	t.Run("saying hello world world by default in English", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assert(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Vikram", "Spanish")
		want := "Hola, Vikram"

		assert(t, got, want)
	})

	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Vikram", "French")
		want := "Bonjour, Vikram"

		assert(t, got, want)
	})

	t.Run("saying hello in Kannada", func(t *testing.T) {
		got := Hello("Vikram", "Kannada")
		want := "Namaskara, Vikram"

		assert(t, got, want)
	})

}
