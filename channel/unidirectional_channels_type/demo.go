package main

import "fmt"

func sender(ch chan<- int) {
	ch <- 88
}
func f(ch chan int) {
	ch <- 99
	close(ch)
}
func receiver(ch <-chan int, done chan bool) {
	fmt.Println("Receiver ", <-ch)
	done <- true
}
func main() {
	ch := make(chan int)
	done := make(chan bool)
	go sender(ch)
	go receiver(ch, done)
	<-done
	fmt.Println("Done")
}
