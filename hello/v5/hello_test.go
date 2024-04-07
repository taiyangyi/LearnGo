package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Bamaw")
		want := "Hello,Bamaw"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string default to 'Go'", func(t *testing.T) {
		got := Hello("")
		want := "Hello,Go"
		assertCorrectMessage(t, got, want)
	})

}
