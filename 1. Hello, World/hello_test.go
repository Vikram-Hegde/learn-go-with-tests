package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Vikram")
	want := "Hello, Vikram"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}