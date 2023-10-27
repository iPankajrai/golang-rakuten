package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ping, pong := make(chan string), make(chan string, 1)

	//go func() {
	pong <- "Ping"
	//}()

	go Ping(ping, pong)
	go Pong(pong, ping)

	runtime.Goexit()

}

func Ping(ping chan<- string, pong <-chan string) {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println(<-pong, "--> ", time.Now().Unix())
		ping <- "Ping"
	}
}

func Pong(pong chan<- string, ping <-chan string) {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println(<-ping, "--> ", time.Now().Unix())
		pong <- "Pong"
	}
}

// Ping
// Pong
