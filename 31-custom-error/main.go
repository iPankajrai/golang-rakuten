package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("abcd.txt")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		buf := make([]byte, 100)
		fmt.Println(f.Read(buf))
	}

	str := err.Error()
	fmt.Println("Str is a string contains:", str)

	//....

	r1 := &Rect{L: -12.34, B: 34.56}
	if a1, err := r1.Area(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Area of Rect:", a1)
	}

}

type Rect struct {
	L, B float32
}

func (r *Rect) Area() (float32, error) {

	if r.L <= 0 {
		return 0., &MyError{Code: "ERR-Length", Message: fmt.Sprint("Invalid Length field.Given Length is:", r.L)}
	}
	if r.B <= 0 {
		return 0., &MyError{Code: "ERR-Bredth", Message: fmt.Sprint("Invalid Bredth field.Given Bredth is:", r.B)}
	}

	return r.L * r.B, nil
}

type MyError struct {
	Code, Message string
}

func (me *MyError) Error() string {
	return fmt.Sprintf("Custom Error Code:%s,Message:%s", me.Code, me.Message)
}

// 1. Error
// 2- Panic
// 3. Fatal
