package main

import (
	"fmt"
	"strings"
)

// Hello returns greeting for passed username.
func Hello(name string) string {
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
	fmt.Printf(Hello("user"))
	fmt.Printf(Hello(""))
}
