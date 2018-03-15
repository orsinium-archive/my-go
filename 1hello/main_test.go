package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("\ngot:\n'%s'\nwant\n'%s'", got, want)
		}
	}

	t.Run("saying hello to some user", func(t *testing.T) {
		got := Hello("user")
		want := "Hello, User!\n"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to world", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!\n"
		assertCorrectMessage(t, got, want)
	})

}
