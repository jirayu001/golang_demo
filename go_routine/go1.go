package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println(1)
	go fmt.Println(2)
	go fmt.Println(3)
	go fmt.Println(4)
	time.Sleep(100 * time.Microsecond)
}
