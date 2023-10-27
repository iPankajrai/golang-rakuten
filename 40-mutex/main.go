package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0 // This is a global variable

func main() {
	mu := new(sync.Mutex)
	//wg.Add(4) // it is okay to add two since there are two goroutines other than main
	go incr(mu)
	go decr(mu)
	go incr(mu)
	for i := 1; i <= 10; i++ {
		mu.Lock()
		count--
		mu.Unlock()
	}
	//go decr(mu)
	//wg.Wait()
	time.Sleep(time.Millisecond * 1)
	fmt.Println(count)

}

func incr(mu *sync.Mutex) {
	for i := 1; i <= 10; i++ {
		mu.Lock()
		count++
		mu.Unlock()
	}
}

func decr(mu *sync.Mutex) {
	for i := 1; i <= 10; i++ {
		mu.Lock()
		count--
		mu.Unlock()
	}
}
