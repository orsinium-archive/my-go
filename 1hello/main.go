package main

import (
	"fmt"
	"strings"
)

func hello(name string) string {
	if name == "" {
		// set default value
		name = "World"
	} else {
		// set first letter to uppercase
		name = strings.Title(name)
	}
	// make and return hello string
	return "Hello, " + name + "!\n"
}

func main() {
	fmt.Printf(hello("user"))
	fmt.Printf(hello(""))
}
