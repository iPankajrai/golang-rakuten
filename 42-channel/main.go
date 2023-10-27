package main

import (
	"demo/sq"
	"fmt"
	"runtime"
)

var ch chan int

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
}

func main() {
	// go sendSq(50)
	// go receiveSq(50)

	go sq.SendSq(50, ch)
	//go sq.SendSq(50, ch)
	//go sq.ReceiveSq(50, ch)
	//go sq.ReceiveSqR(ch)
	go sq.ReceiveSqLoop(ch)

	runtime.Goexit()
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
