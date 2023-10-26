package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Hello Main")

	go func() {
		count := 1
	loop:
		fmt.Println("I am another goroutine-->", count)
		if count <= 100 {
			count++
			goto loop
		}
	}()

	go func() {
		go func() {
			for i := 1; i <= 10; i++ {
				fmt.Println("Hello I ama special case go rountine inside another one-->", i)
				time.Sleep(time.Millisecond * 100)
			}
		}()
		for i := 1; ; {
			fmt.Println("Another gorouinte--i", i)
			i++
			if i == 100 {
				runtime.Goexit()
			}
		}
	}()

	for i := 1; i <= 100; i++ {
		fmt.Println("I am called in the main-->", i)
	}
	go func() {
		for {
		}
	}()
	//time.Sleep(time.Second * 2)
	runtime.Goexit()
} // What is it? Exit of main is nothing but exit of your application
// Goexit() -> 1. When you call in Main : It waits for all other gorouintes to complete their execition and then exit with error status
// 			   2. When you call in another Goroutin
