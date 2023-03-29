package main

import "fmt"

func main() {

	x := add(1, 3)

	fmt.Println(x)
}
func add(a, b int) int {
	return a + b
}
