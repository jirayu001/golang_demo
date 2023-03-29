package main

import "fmt"

func sender(ch chan int, done chan string) {
	for i := 0; i <= 5; i++ {
		fmt.Println("Sending value : ", i)
		ch <- i
	}
	close(ch)
	fmt.Println("Sender is about to complete")
	done <- "Done from Sender"
	fmt.Println("Sender done")
}
func receiver(ch chan int, done chan string) {
	for v := range ch {
		fmt.Println("\tReceiver value : ", v)

	}

	/*for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("\tReceiver value : ", v, ok)

	}*/

	//v := <-ch
	/*fmt.Println("\tReceiver value : ", <-ch) //0
	fmt.Println("\tReceiver value : ", <-ch) //1
	fmt.Println("\tReceiver value : ", <-ch) //2
	fmt.Println("\tReceiver value : ", <-ch) //3
	fmt.Println("\tReceiver value : ", <-ch) //4
	fmt.Println("\tReceiver value : ", <-ch) //5
	fmt.Println("\tReceiver value : ", <-ch) //0
	fmt.Println("\tReceiver value : ", <-ch) //0
	fmt.Println("\tReceiver value : ", <-ch) //0
	fmt.Println("\tReceiver value : ", <-ch) //0*/
	done <- "Done from receiver"
	close(done)
}
func square(ch chan int, chSquare chan int, done chan string) {
	for v := range ch {
		chSquare <- v * v
	}
	close(chSquare)
	done <- "Done from square"
}
func main() {
	ch := make(chan int)
	chSquare := make(chan int)
	done := make(chan string) // done <- struct{}{}
	go sender(ch, done)
	go square(ch, chSquare, done)
	go receiver(chSquare, done)

	for v := range done {

		fmt.Println("Done status : ", v)
	}
	//doneStruct := <-done
	//fmt.Println("Done status : ", doneStruct)
	//doneStruct = <-done
	//fmt.Println("Done status : ", doneStruct)
	//doneStruct = <-done
	//fmt.Println("Done status : ", doneStruct)

	fmt.Println("Main exit")

}
