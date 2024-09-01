package main

import (
	"testing"
)

func TestPrintHello(t *testing.T) {
	t.Run("say hello to people", func (t *testing.T) {
		got := printHello("", "Vikram")
		want := "Hello, Vikram!"
		assert(t, got, want)
	})
	t.Run("handle empty string as an argument", func (t *testing.T) {
		got := printHello("", "")
		want := "Hello, World!"
		assert(t, got, want)
	})
	t.Run("say hello in Spanish", func (t *testing.T) {
		got := printHello("Spanish", "Vikram")
		want := "Hola, Vikram!"
		assert(t, got, want)
	})
	t.Run("say hello in Spanish with empty string as an argument", func (t *testing.T) {
		got := printHello("Spanish", "")
		want := "Hola, Amigo!"
		assert(t, got, want)
	})
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}