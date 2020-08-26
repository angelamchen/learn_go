package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper() // tells the test suite that this method is a helper (when it fails, the line # reported is function call rather than test helper)
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Angela", "")
		want := "Hello, Angela"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	
}