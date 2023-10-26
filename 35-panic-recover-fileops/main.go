package main

import (
	"fmt"
	"os"
)

func main() {
	//f, err := os.OpenFile("abcd.txt", 0755, syscall.O_RDWR)

	func() {
		f, err := os.Create("abcd.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		func(f *os.File) {
			defer RevoverMe()
			for a, b := 10, 9; a >= 0; a, b = a-1, b-1 {
				c := a / b
				_, err = f.WriteString(fmt.Sprintln("Value of a,b, a/b:", a, b, c))
				if err != nil {
					fmt.Println(err)
				}
			}
		}(f)
		fmt.Println("Done with this writing, whether panic or not")
	}()
	func(fn string) {
		buf, err := os.ReadFile(fn)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(buf))
		}
		//f.Close()

	}("abcd.txt")
}

func RevoverMe() {
	if r := recover(); r != nil { // there is some panic happned
		fmt.Println("---->", r) // What do you or how do you want to handle when there is a panic
	}
}
