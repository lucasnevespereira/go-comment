package main

import "fmt"

func main() {
	result := sum(2, 2)
	fmt.Println(result)
}

func sum(a, b int) int {
	return a + b
}
