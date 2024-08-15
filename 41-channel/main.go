package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // shorthand instantiation of Unbuffered channel

	go func() {
		time.Sleep(time.Second * 5)
		ch <- 100 // ch <- 100 send a value to the channel
	}()

	v := <-ch // <-ch receive a value from the channel
	fmt.Println("Will I be printed?")
	fmt.Println(v)

	//Example 2 with buffered channel
	ch2 := make(chan int, 5) //buffered channel
	go func() {
		time.Sleep(time.Second * 10)
		ch2 <- 9988888
	}()
	r := <-ch2
	fmt.Println("I will be printed after 10 seconds!!")
	fmt.Println(r)

}

// Which goroutine is the sender and which goroutine is the receiver
