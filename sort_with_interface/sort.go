package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type byName []*Person

type customSort struct {
	Persons []*Person
	less    func(i, j int) bool
}

func (p byName) Len() int {
	return len(p)
}
func (p byName) Less(i, j int) bool {
	return p[i].name < p[j].name
}
func (p byName) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func main() {
	p := []*Person{
		{"A", 22},
		{"B", 21},
		{"B", 24},
		{"C", 25},
		{"C", 20},
		{"A", 21},
		{"A", 22},
		{"B", 22},
	}

	printPerson(p)

	sort.Sort(byName(p))
	printPerson(p)
	//sort.Sort(customSort{Persons: p,less: func(i, j int) bool {}})
}
func printPerson(p []*Person) {
	fmt.Println("-------")
	for _, v := range p {
		fmt.Println(*v)
	}
}
