package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 [3]int // type inference
	fmt.Println(arr1)

	var arr2 [4]bool
	fmt.Println(arr2)

	arr1[0] = 100
	arr1[1] = 200
	arr1[2] = 300
	fmt.Println(arr1)

	arr3 := [3]int{10, 11, 12} // short hand declartion of array

	arr4 := [...]int{123, 12, 124, 43, 23, 34, 4} // the length of an array is evaluated at compile time

	fmt.Println(arr3, arr4)

	fmt.Println("Type of arr3:", reflect.TypeOf(arr3))
	fmt.Println("Type of arr4:", reflect.TypeOf(arr4))

	fmt.Println("Sum of elements of array1:", sumofElements(arr1))
	fmt.Println("Sum of elements of array4:", sumofElements1(arr4))

	arr5 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr5)
	arr6 := [2][2][2]int{{{1, 2}, {3, 4}}, {{1, 2}, {3, 4}}}
	fmt.Println(arr6)

	// for range loop

	for i, v := range arr4 {
		fmt.Println(i, "=>", v)
		//fmt.Println(v)
	}
	for i := range arr4 {
		//fmt.Println(i, "=>", v)
		fmt.Println(i)
	}

}

func sumofElements(arr [3]int) int {
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	return sum
}

func sumofElements1(arr any) int {
	sum := 0

	switch arr.(type) {
	case [3]int:
		for i := 0; i < len(arr.([3]int)); i++ {
			sum += arr.([3]int)[i]
		}
	case [7]int:
		for i := 0; i < len(arr.([7]int)); i++ {
			sum += arr.([7]int)[i]
		}
	}

	return sum
}
