package sq

import "fmt"

func init() {
	fmt.Println("I am from sq package-1")
}

func init() {
	fmt.Println("I am from sq package-2")
}

func SendSq(num uint, ch chan<- int) {
	for i := 1; i <= int(num); i++ {
		ch <- i * i
	}
	close(ch)
}

func ReceiveSq(num uint, ch <-chan int) {
	for i := 1; i <= int(num); i++ {
		fmt.Println(<-ch)
	}
}

func ReceiveSqR(ch <-chan int) {
	for v := range ch { // range loop on a channel receive values until the channel is closed
		fmt.Println(v)
	}
}

func ReceiveSqLoop(ch <-chan int) {
	for {
		v, ok := <-ch // can check this way whether a channel is closed or not. If ok is false that means the channel is closed
		if !ok {
			return
		}
		fmt.Println(v)
	}
}

func ReceiveSqRDone(ch <-chan int, done chan<- bool) {
	for v := range ch { // range loop on a channel receive values until the channel is closed
		fmt.Println(v)
	}

	done <- true
}
