package main

import "fmt"

func declaration() {
	var a int
	a = 1

	var b int = 2.0

	c := 3

	fmt.Printf("%d", a+b+c)
}
