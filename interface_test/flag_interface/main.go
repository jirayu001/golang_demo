package main

import (
	"flag"
	"fmt"
	"strconv"
)

var romanMap = map[string]int{
	"i":     1,
	"ii":    2,
	"iii":   3,
	"iiii":  4,
	"iiiii": 5,
}

type RomanAge struct {
	age string // i = 1 , ii = 2 , v = 5
}

func (r *RomanAge) String() string {
	return strconv.Itoa(romanMap[r.age])
}
func (r *RomanAge) Set(value string) error {
	r.age = value
	return nil
}
func flagRomanAge() *RomanAge {
	romanAge := RomanAge{}
	flag.Var(&romanAge, "age", "help message for flagname")
	return &romanAge
}

var name = flag.String("name", "anonymous", "Your name")

var age = flagRomanAge()

func main() {
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(age)
	// go run oop/interface/flag -name Filicity -age iii
	// go run main.go -name Filicity -age iii
}
