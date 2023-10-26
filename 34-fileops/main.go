package main

import (
	"fmt"
	"os"
)

func main() {
	//f, err := os.OpenFile("abcd.txt", 0755, syscall.O_RDWR)
	f, err := os.Create("abcd.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	for a, b := 10, 9; a >= 0; a, b = a-1, b-1 {
		c := a / b
		_, err = f.WriteString(fmt.Sprintln("Value of a,b, a/b:", a, b, c))
		if err != nil {
			fmt.Println(err)
		}
	}

	//f.Close()
}
