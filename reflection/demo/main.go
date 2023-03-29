package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct{ a int }

func main() {
	look(3)
	look(struct{ a int }{a: 3})
	look(MyStruct{a: 3})
}

func look(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(reflect.TypeOf(v).Kind())
	fmt.Println(reflect.ValueOf(v))
}
