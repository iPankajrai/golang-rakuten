package main

import "fmt"

func main() {

	a := 10
	defer func(b *int) {
		*b = *b + 10
		fmt.Println("a during defer:", *b)
	}(&a)
	a = 0
	fmt.Println(a)
	//a = a - 10
	//a = a - 10
	//fmt.Println(a)
	fmt.Println()
	str := "Hello Rakuten"

	for _, v := range str {
		defer fmt.Print(string(v))
	}
}
