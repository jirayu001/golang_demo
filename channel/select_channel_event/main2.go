package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func running(name string, track chan<- struct{}, finishRacing <-chan struct{}) {
	defer func() {

		fmt.Printf("##### %s stop running #####\n", name)
	}()

	rand.Seed(time.Now().UnixNano())
	//<- chan Time
	chFinish := time.After(time.Duration(1*time.Second + time.Duration(rand.Intn(100))*time.Millisecond))

	for recing := true; !recing; {
		select {
		case <-chFinish:
			track <- struct{}{}
		case <-finishRacing:
			recing = false

		}
	}

}
func main() {
	defer func() {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Done")

	}()

	chInt := make(chan os.Signal, 1)
	//signal.Notify(chInt, os.Interrupt)
	signal.Notify(chInt, syscall.SIGINT)

	track1 := make(chan struct{})
	track2 := make(chan struct{})
	track3 := make(chan struct{})

	abort := make(chan struct{})
	finishRacing := make(chan struct{})

	go running("h1", track1, finishRacing)
	go running("h2", track2, finishRacing)
	go running("h3", track3, finishRacing)

	lapseChan := make(chan time.Time)
	go func() {

		lapseTricker := time.NewTicker(1 * time.Second)

		for {
			lapseChan <- (<-lapseTricker.C)
		}

	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("About to abort.")
		abort <- struct{}{}
	}()
	for done := false; !done; {
		select {
		case <-track1:
			fmt.Println("Winner is Horse 1")
			done = true
		case <-track2:
			fmt.Println("Winner is Horse 2")
			done = true
		case <-track3:
			fmt.Println("Winner is Horse 3")
			done = true
		case v := <-lapseChan:
			fmt.Println("Horse are running.", v.Second())
		case <-abort:
			//finishRacing <- struct{}{}
			close(finishRacing)
			done = true
			fmt.Println("User pressing Enter")
		case v := <-chInt:
			//finishRacing <- struct{}{}
			close(finishRacing)
			done = true
			fmt.Printf("User interrupt : %v \n", v)
			//default :
			//fmt.Println("Horse are running.")
		}
		time.Sleep(1 * time.Second)
	}
}
