package main

import "fmt"

func main() {
    fmt.Println("Hello")
    foo("a")
}

func foo(a string) bool {
	if a == "a" {
		return true
	}

	fmt.Println("not a")

	return false
}
