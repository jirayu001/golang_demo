// Data Race
// ตั้งแต่ 2 goroutine ขึ้นไป มีการใช้ variable ตัวเดียวกัน และใช้พร้อมๆกัน
// อย่างน้อย 1 goroutine มีการเขียนข้อมูลไปใน variabla นั้นๆ
package main

import (
	"fmt"
	"sync"
)

var chBalance = make(chan int)
var chDeposit = make(chan int, 10000)

func bankSystem() {
	balance := 0
	for {
		if len(chDeposit) == 0 {
			select {
			case x := <-chDeposit:
				balance += x
			case chBalance <- balance:
				//do nothing
			}
		} else {
			select {
			case x := <-chDeposit:
				balance += x

			}

		}
	}
}

func deposit(x int, w *sync.WaitGroup) {
	chDeposit <- x
	w.Done()
}

func finalBalance() int {
	return <-chBalance

}

func main() {
	go bankSystem()
	w := &sync.WaitGroup{}
	fmt.Println("Balance before deposit: ", finalBalance())
	for i := 0; i < 2000000; i++ {
		w.Add(1)
		go deposit(100, w)

	}
	fmt.Println("Balance before waitgroup: ", finalBalance())
	w.Wait()
	fmt.Println("Balance : ", finalBalance())
	fmt.Println("Balance : ", finalBalance())

}
