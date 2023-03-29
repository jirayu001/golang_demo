// Data Race
// ตั้งแต่ 2 goroutine ขึ้นไป มีการใช้ variable ตัวเดียวกัน และใช้พร้อมๆกัน
// อย่างน้อย 1 goroutine มีการเขียนข้อมูลไปใน variabla นั้นๆ
package main

import (
	"fmt"
	"sync"
)

var mu = sync.Mutex{}
var balance = 0

func deposit(x int, w *sync.WaitGroup) {
	//start
	mu.Lock()
	defer func() {
		mu.Unlock()
		w.Done()
	}()
	balance += x
	addBonus(1)

}

func AddBonus(x int) {
	mu.Lock()
	defer func() {
		mu.Unlock()
	}()
	addBonus(x)
}

func addBonus(x int) {

	balance += x
}

func finalBalance() int {
	/*mu.Lock()
	defer func() {
		mu.Unlock()
	}()*/
	return balance

}

func main() {
	w := &sync.WaitGroup{}
	for i := 0; i < 200000; i++ {
		w.Add(1)
		go deposit(100, w)

	}
	w.Wait()
	fmt.Println("Balance : ", finalBalance())

}
