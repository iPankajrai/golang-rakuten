package main

import "fmt"

type integer = int

func main() {

	var num1 integer = 100

	var any1 any = num1

	any1 = false

	//var ok1 bool = bool(any1) // not okay

	var ok1 bool = any1.(bool) // okay with type assertion

	fmt.Println(ok1)

	any1 = "Hello Rakuten"

	any1 = 3.14

	any1 = 'A'

	num2 := 100
	num3 := 12.45
	var num4 uint8 = 100
	var num5 float32 = 12.45
	var any2 any = 12312.12

	var float2 float32 = 12312.123
	var any3 any = float2

	var sum float64 = float64(num2) + num3 + float64(num4) + float64(num5) + any2.(float64) + float64(any3.(float32))
	fmt.Println(sum)
	fmt.Printf("%.2f", sum)
	// math package contains round, ceil ,abc
}
