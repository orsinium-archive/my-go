package main

import "fmt"

func types() {
	var a string
	a = "lol"

	var b int32 = 2

	c := 3.0 // float

	d := [3]int{1, 2, 3} // array

	e := []int{1, 2, 3} // array

	fmt.Printf("%s %d %f %d %d", a, b, c, d[0], e[0])
}
