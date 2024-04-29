package main

import "fmt"

func main() {
    fmt.Println("Hello")
    foo("a")
}

func foo(a string) error {
	if a == "a" {
		return nil
	}

	return fmt.Errorf("OOPS")
}
