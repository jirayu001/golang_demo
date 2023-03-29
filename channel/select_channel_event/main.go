package main

import (
	"fmt"
	"math/rand"
	"time"
)

func running(track chan<- struct{}) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(3 * time.Duration(rand.Intn(100)) * time.Microsecond)
	track <- struct{}{}

}
func main() {
	track1 := make(chan struct{})
	track2 := make(chan struct{})
	track3 := make(chan struct{})

	go running(track1)
	go running(track2)
	go running(track3)

	select {
	case <-track1:
		fmt.Println("Winner is Horse 1")
	case <-track2:
		fmt.Println("Winner is Horse 2")
	case <-track3:
		fmt.Println("Winner is Horse 3")
	}
	fmt.Println("Done")
}
