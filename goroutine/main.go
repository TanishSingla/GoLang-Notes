package main

import (
	"fmt"
	"sync"
	// "time"
)

// func sayHello() {
// 	fmt.Println("hello Function")
// 	time.Sleep(3000 * time.Millisecond)
// 	fmt.Println("hello function ended")
// }

// func sayHi() {
// 	fmt.Println("hi function")
// 	time.Sleep(1000 * time.Millisecond)
// 	fmt.Println("hi function ended")
// }

func temp(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Function temp Started", i)

	// some task
	fmt.Println("Function temp Ended", i)

}
func main() {
	// fmt.Println("-----------GOROUTINE---------------")
	// go sayHello()
	// go sayHi()
	// time.Sleep(3100 * time.Millisecond)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go temp(i, &wg)
	}
	fmt.Println("Main Function ended")
	wg.Wait()
}
