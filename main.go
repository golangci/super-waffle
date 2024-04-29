package main

import "fmt"

func main() {
    fmt.Println("Hello")
}

func foo(a string) error {
	if a == "a" {
		return nil
	}

	return fmt.Errorf("OOPS")
}
