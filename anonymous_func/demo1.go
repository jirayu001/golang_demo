package main

import "fmt"

func createF() func() int {
	var x = 0
	return func() int {
		x++
		return x
	}
}

func main() {
	f := createF()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3

	f1 := createF()
	fmt.Println(f1()) // 1
	fmt.Println(f1()) // 2
	fmt.Println(f1()) // 3
}
