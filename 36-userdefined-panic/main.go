package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Start of main")
	fn := "Ramesh"
	ln := ""
	if fn != "" && ln != "" {
		FullName(fn, ln)
	} else {
		fmt.Println("ln is empty")
	}
	func() {
		defer RevoverMe()
		FullName("Jiten", "")
	}()
	func() {
		defer RevoverMe()
		FullName("", "Tiwari")
	}()

	FullName("Lakshmi", "Kumari")
	Fatal("seems no error but let me Fatal it")
	fmt.Println("End of main")
}

func Fatal(msg string) {
	fmt.Println(msg)
	os.Exit(1) // 0 exit code is always success. Non zero is failure
}

// FullName function takes input of fn and ln.
// if fn or ln is empty it panics
func FullName(fn, ln string) {
	//defer RevoverMe()
	if fn == "" {
		panic("firstname is empty")
	}

	if ln == "" {
		panic("lastname is empty")
	}

	fmt.Println(fn, ln)
}

func RevoverMe() {
	if r := recover(); r != nil { // there is some panic happned
		fmt.Println("---->", r) // What do you or how do you want to handle when there is a panic
	}
}
