package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := GenSq(20)
	ch2 := GenSq(10)
	// done1 := Receive(ch1)
	// done2 := Receive(ch2)
	// <-done1
	// <-done2

	<-Receive(ch1)
	<-Receive(ch2)

	// var e Empty
	// e.Todo()

}

func GenSq(num int) chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= num; i++ {
			time.Sleep(time.Second * 1)
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

func Receive(ch <-chan int) chan struct{} {
	done := make(chan struct{})
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		done <- struct{}{} // Signal to the receiver but no data
	}()
	return done
}

// Empty struct does not store any value but it can be used to give a signal

// type Empty struct{}

// func (e Empty) Todo() {
// 	//todo
// }
