package main

import "fmt"

func ExampleIfStatement() {
	capitals := map[string]string{"NSW": "Sydney", "WA": "Perth"}
	for _, state := range [2]string{"NSW", "AL"} {
		if capital, known := capitals[state]; known {
			fmt.Printf("The capital of %s is %s\n", state, capital)
		} else {
			fmt.Printf("I don't know the capital of %s\n", state)
		}
	}
	// Output:
	// The capital of NSW is Sydney
	// I don't know the capital of AL
}
