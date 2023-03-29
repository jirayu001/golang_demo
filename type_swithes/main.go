package main

import (
	"fmt"
)

func testValue(x interface{}) {
	if v, ok := x.(string); ok {
		fmt.Println("string value: ", v)
		return
	}
	if v, ok := x.(int); ok {
		fmt.Println("int value: ", v)
		return
	}
	if v, ok := x.(bool); ok {
		fmt.Println("bool value: ", v)
		return
	}
	if v, ok := x.(float32); ok {
		fmt.Println("float32 value: ", v)
		return
	}
	if v, ok := x.(float64); ok {
		fmt.Println("float64 value: ", v)
		return
	}
	fmt.Println("No match any")
}
func testValueSwitch(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Println("string value: ", v)
	case int:
		fmt.Println("int value: ", v)
	case bool:
		fmt.Println("bool value: ", v)
	case float32:
		fmt.Println("float32 value: ", v)
	case float64:
		fmt.Println("float64 value: ", v)
	default:
		fmt.Println("No match any")
	}
}

func main() {
	testValue("sdf")
	testValue(1234)
	testValue(true)
	testValue(234.23)
	testValue(float32(234.23))

	fmt.Println("--------------")

	testValueSwitch("sdf")
	testValueSwitch(1234)
	testValueSwitch(true)
	testValueSwitch(234.23)
	testValueSwitch(float32(234.23))
}
