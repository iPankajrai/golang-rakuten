package main

import "fmt"

func main() {
	var num1 int = 100
	ptr1 := &num1
	println(num1, *ptr1)
	//fmt.Println(num1, false, "Hello")
	//arr1 := [...]int{10, 12, 123, 43, 43, 34}
	slice1 := make([]int, 3)
	//println(arr1, slice1)
	slice1 = squareSlice(slice1)

	a, b, c := 10, 20, 30
	// t := b
	// b = a
	// a = t

	a, b, c = b, c, a

	fmt.Println(a, b, c)
}

// this is not a wrong code but not efficient code
func square(num int) *int {
	var ptr *int = &num
	*ptr = num * num
	return ptr
}

// this is not a wrong code but not efficient code
func squareParam(num int, ptr *int) {
	if ptr != nil {
		*ptr = num * num
	}
}

func getPointer() *int {
	ptr := new(int)
	*ptr = 100 * 100
	return ptr
}

func squareSlice(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		slice[i] = slice[i] * slice[i]
	}
	slice = append(slice, 100)
	return slice
}
