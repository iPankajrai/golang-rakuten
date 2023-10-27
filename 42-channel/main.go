package main

import (
	"demo/sq"
	"fmt"
)

var ch chan int
var done chan bool

// init func is automatically called/executed.
// there can be any anumber of init funcs in a package

func init() {
	fmt.Println("I am init-1")
}

func init() {
	fmt.Println("I am init-2")
}

func init() {
	ch = make(chan int)
	done = make(chan bool)
}

func main() {
	// go sendSq(50)
	// go receiveSq(50)

	go sq.SendSq(5000, ch)
	//go sq.SendSq(50, ch)
	//go sq.ReceiveSq(50, ch)
	//go sq.ReceiveSqR(ch)
	//go sq.ReceiveSqLoop(ch)
	go sq.ReceiveSqRDone(ch, done)

	//runtime.Goexit()
	<-done // The receiver is blocked untile the sender sends the value and it is received

}

func sendSq(num uint) {
	for i := 1; i <= int(num); i++ {
		ch <- i * i
	}
}

func receiveSq(num uint) {
	for i := 1; i <= int(num); i++ {
		fmt.Println(<-ch)
	}
}
