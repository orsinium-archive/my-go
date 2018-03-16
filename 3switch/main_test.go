package main

import "testing"

func TestFizzBazz(t *testing.T) {
	got := FizzBazz(1)
	want := "1"
	if got != want {
		t.Errorf("\ngot:\n'%s'\nwant\n'%s'", got, want)
	}

	got = FizzBazz(3)
	want = "Fizz"
	if got != want {
		t.Errorf("\ngot:\n'%s'\nwant\n'%s'", got, want)
	}

	got = FizzBazz(5)
	want = "Bazz"
	if got != want {
		t.Errorf("\ngot:\n'%s'\nwant\n'%s'", got, want)
	}

	got = FizzBazz(15)
	want = "FizzBazz"
	if got != want {
		t.Errorf("\ngot:\n'%s'\nwant\n'%s'", got, want)
	}
}
