package main

import (
	"fmt"
	"runtime"
	"sync"
)

var balance = 0

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.NumCPU())
	wg := sync.WaitGroup{}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			balance++
		}()
	}
	wg.Wait()
	fmt.Println(balance)

}
