package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := GenSq(time.Second * 2)
	ch2 := GenSq(time.Second * 1)
	justCloseCh := JustClose(time.Second * 5)
	//timeOutCh := time.After(time.Second * 10)

outer:
	for {
		select {
		case v1 := <-ch1:
			fmt.Println("Received from chan-1", v1)
		case v2 := <-ch2:
			fmt.Println("Received from chan-2", v2)
		case v3 := <-justCloseCh: // when a channel is closed it sends a default value
			fmt.Println("Just close channel:", v3)
			fmt.Println("Timedout---")
			break outer

		// case <-timeOutCh:
		// 	fmt.Println("Timedout---")
		// 	break outer
		default:
			//fmt.Println("Default no value")
		}
	}
}

func GenSq(duration time.Duration) chan int {
	ch := make(chan int)
	go func() {
		for i := 1; ; i++ {
			time.Sleep(duration)
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

func JustClose(duration time.Duration) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(duration)
		close(ch)
	}()
	return ch
}
