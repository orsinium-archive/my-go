package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("")
	want := "Hello, World!\n"
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

	got = Hello("user")
	want = "Hello, User!\n"
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}

}
