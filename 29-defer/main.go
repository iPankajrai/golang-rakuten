package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	defer func() {
		defer fmt.Println("Hello Rakuten-2")
		fmt.Println("Hello Rakuten-1")
	}()

	defer fmt.Println("I am end of main function")

	fmt.Println("I am start of main functin")

	f, err := os.Open("go.mod")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}

	buf := make([]byte, 10)
	for {
		if n, err := f.Read(buf); (err != nil && err == io.EOF) || n <= 0 {
			println()
			break
		}
		fmt.Print(string(buf))
	}

}

// defer defers the execution towards the end of the  callstack
