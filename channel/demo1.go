package main

import (
	"fmt"
	"math/rand"
	"time"
)

func trailing(ch chan bool) {
	rand.Seed(int64(time.Now().Nanosecond()))
	x := rand.Intn(1000)
	fmt.Println("trining for : ", x, "milisecond")
	time.Sleep(time.Duration(x) * time.Millisecond)
	ch <- false
}
func main() {
	done := make(chan bool)
	go trailing(done)
	trainingStatus := <-done
	fmt.Println("Done status : ", trainingStatus)
}
