package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffered channel with the capacity of 2.
	ch <- 100
	ch <- 200
	v1 := <-ch
	ch <- 300
	fmt.Println(v1)
	v2 := <-ch
	fmt.Println(v2)
	fmt.Println(<-ch)
}

// Buffered channels
