package main

import "fmt"

// main is the entry point of the program.
// It calls the sum function with arguments 2 and 2,
// and prints the result.
func main() {
	result := sum(2, 2)
	fmt.Println(result)
}

// sum takes two integer arguments, a and b,
// and returns their sum.
func sum(a, b int) int {
	return a + b
}
