package main

import (
	"fmt"
	"runtime"
)

var ch chan int

func init() {
	ch = make(chan int, 3)
}

// What should be the size of the buffer?
func main() {
	go sendSq(10)
	//go sendSq(10)

	go receiveSq("receiver1")
	go receiveSq("receiver2")
	go receiveSq("receiver3")
	// go func() {
	// 	for {
	// 		if ch == nil {
	// 			fmt.Println("chan is nil")
	// 			runtime.Goexit()
	// 		}
	// 	}
	// }()

	runtime.Goexit()
}

func sendSq(num uint) {
	for i := 1; i <= int(num); i++ {
		//v, ok := <-ch // This is a receiver
		//if ok {
		if ch != nil {
			ch <- i * i
		}
		//}
	}
	close(ch)
}

func receiveSq(name string) {
	for v := range ch {
		fmt.Println("Received from:", name, v)
	}
}
