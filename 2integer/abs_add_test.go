package main

import (
	"fmt"
	"testing"
)

func TestAbsAdd(t *testing.T) {
	sum := AbsAdd(2, 4)
	expected := 6
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}

	sum = AbsAdd(2, -4)
	expected = 2
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAbsAdd() {
	sum := AbsAdd(2, -8)
	fmt.Println(sum)

	sum = AbsAdd(3, 9)
	fmt.Println(sum)

	// Output:
	// 6
	// 12
}
