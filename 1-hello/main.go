package main // mandatory to have main function

import (
	"fmt"
	"strconv"
	"time"
)

func main() { // func main must be there inside main package

	println("Hello Rakuten")

	fmt.Println("Hello Rakuten")

	var x int = 42
	var y float64 = float64(x)
	fmt.Println(y)

	go func() {
		fmt.Println("Hello from goroutine")
	}()
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from main")

	num, err := strconv.Atoi("123")

	fmt.Println(num, err)

}
