package main

import "fmt"

func main() {
	fmt.Println("Hello", "Rakuten", "!")
	arr := [...]int{13, 43, 434}
	slice := []int{1, 20, 45, 2, 3, 45, 87, 98, 678}

	fmt.Println(sum())

	fmt.Println(sum(100))
	fmt.Println(sum(100, 10))
	fmt.Println(sum(100, 4, 5, 6, 7, 8, 9))
	fmt.Println(sum(arr[:]...)) // an array cannot be passsed by a slice can be
	fmt.Println(sum(slice...))
	fmt.Println("Hello", false, 100, 100.12, 'A', slice, arr, nil)
}

func sum(nums ...int) int {
	sum := 0
	for _, v := range nums {

		sum += v
	}
	return sum
}

// func calc(nums ...int, method string) int {
func calc(method string, nums ...int) int {

	cac := 0
	switch method {
	case "add", "ADD":
		for _, v := range nums {

			cac += v
		}
		return cac
	case "sub", "SUB":
		for _, v := range nums {

			cac -= v
		}
		return cac

	case "mul", "MUL":
		for _, v := range nums {

			cac *= v
		}
		return cac

	default:
		return 0
	}

}

// if you see an ... symbol in a function or a method it is a variadic parameter
// variadic parameter : you can pass any number of arguments of the same type
// the variadic parameter must be the last parameter for a func or method
// variadic arguments can be fetched using for range loop
