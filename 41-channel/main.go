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
}

// Which goroutine is the sender and which goroutine is the receiver
