package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World;Starts main")
	go fmt.Println("Hello Gorotine")
	go func() {
		c := add(10, 20)
		fmt.Println("Additon:", c)
	}()
	//defer func() {
	go func() {
		for i := 1; i <= 10000; i++ {
			println("Hello 2nd Goroutine-->", i)
			time.Sleep(time.Microsecond * 100)
		}
	}()

	go func() {
		for i := 1; i <= 10000; i++ {
			println("Hello 3nd Goroutine-->", i)
			time.Sleep(time.Microsecond * 100)
		}
	}()
	for i := 1; i <= 100; i++ {
		println("Hello Main Goroutine-->", i)
		time.Sleep(time.Microsecond * 100)
	}
	//}()
	fmt.Println("Hello World; Ends main")
	time.Sleep(time.Second * 1)
} // What is it? Exit of main is nothing but exit of your application

func add(a, b int) int {
	return a + b
}

// What is a Gorouine?: It is kind of a small/green thread. That is very fast , efficient and scalable
// When to create a Goroutine?When ever you want to execute asynchronously/concurrently then execute it as a Goroutine
// How to create a Goroutine? By using a keyword call go. Appnd the func or method call with go keyword

// 1. Main is also a Goroutine
// 2. By default or design no Gorouine wait to complete the execution of another Goroutine
// 3. The order of execution of Goroutins is not guaranteed
