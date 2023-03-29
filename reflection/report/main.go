package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `report:"ชื่อ,uppercase"`
	Age  int    `report:"อายุ"`
}
type Employee struct {
	name string `report:"ชื่อ"`
	age  int    `report:"อายุ"`
}

func main() {
	person := Person{Name: "David", Age: 12}
	fmt.Println("Before : ", person)

	v := reflect.ValueOf(&person)
	fmt.Println(v)

	v.Elem().Field(0).SetString("DAVID")
	v.Elem().Field(1).SetInt(v.Elem().Field(1).Int() * 2)
	fmt.Println("After :", person)
	/*fmt.Println(report.Text(Person{name: "David", age: 12}))

	fmt.Println(report.Text(Employee{name: "Emily", age: 13}))

	fmt.Println(report.Text(struct {
		name string `report:"FirtName"`
		age  int
	}{name: "Emily", age: 13}))*/
}
