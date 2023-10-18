package main

import "fmt"

func main() {

	func() { // params{
		println("Hello Rakuten")
	}() // executor

	func(a int, b int) {
		println("addition:", a+b)
	}(10, 20)

	c := func(a int, b int) int {
		return a + b
	}(10, 20)
	println("additon c:", c)

	f1 := func(a int, b int) int {
		return a + b
	}

	// func f1(a,b int)int{
	// 	return a + b
	// }

	c2 := f1(10, 20)
	println("additon c2:", c2)

	a1 := calc(10, 20, func(i1, i2 int) int {
		return i1 + i2
	})

	fmt.Println("Addition:a1:", a1)

	s1 := calc(20, 10, sub)
	fmt.Println("Substraction:s1:", s1)

	mul := func(a, b int) int {
		return a * b
	}
	m1 := calc(10, 20, mul)
	fmt.Println("Multiplication:m1:", m1)

	fm1 := multiRet("Hello Rakuten")
	fmt.Println(fm1(10, 20))

	fm2 := mulRet2(10, 20)
	fmt.Println(fm2())
}

func calc(a int, b int, f1 func(int, int) int) int {
	// a++
	// b++
	return f1(a, b)
}

func sub(a, b int) int {
	return a - b
}

func multiRet(str string) func(int, int) int {
	fmt.Println("it is a return function..:", str)
	return func(a int, b int) int {
		return a * b
	}
}

func mulRet2(a, b int) func() int {
	return func() int {
		return a * b
	}
}
