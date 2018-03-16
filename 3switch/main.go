package main

import "fmt"

// FizzBazz is a realisation of FizzBazz algorithm.
func FizzBazz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBazz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Bazz"
	default:
		return fmt.Sprintf("%d", n)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%d: %s\n", i, FizzBazz(i))
	}
}
