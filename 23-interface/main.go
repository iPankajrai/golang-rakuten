package main

import (
	"fmt"
	"os"
)

func main() {
	s := Stdout{}

	fo := FileOut{Filename: "log.out"}

	fmt.Fprintln(s, "Hello Rakuten!")
	fmt.Fprintln(fo, "Hello Rakuten!")
	//fmt.Println("Hello Rakuten!") // it gonna write on your stdout
	// stdin, stdout, stderr
}

// writer could be any type that implements the writer interface

type Stdout struct{}

func (s Stdout) Write(p []byte) (n int, err error) {
	return fmt.Print(string(p))
}

type FileOut struct {
	Filename string
}

func (fo FileOut) Write(p []byte) (n int, err error) {
	f, err := os.Create(fo.Filename)
	if err != nil {
		return -1, err
	}
	return f.Write(p)
}
