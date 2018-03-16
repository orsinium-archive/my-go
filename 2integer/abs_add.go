package main

import (
	"fmt"
)

// AbsAdd return absolute sum of 2 args
func AbsAdd(a, b int) int {
	var sum int
	sum = a + b
	if sum < 0 {
		return -sum
	}
	return sum
}

func main() {
	fmt.Printf("|2 + 4| = %d\n", AbsAdd(2, 4))
	fmt.Printf("|2 - 4| = %d\n", AbsAdd(2, -4))
}
