package main

import "fmt"

func printEachLine(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}
func main() {
	x := []interface{}{"abc", "def", 234, "ghi"}
	printEachLine(x...)
	printEachLine("abc", "def", 234, "ghi")
}
