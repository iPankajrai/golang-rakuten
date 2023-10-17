package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//slice1 := make([]int, 3, 6)
	slice1 := make([]int, 10)
	//slice2 := []int{12, 13, 14, 15}

	for i := range slice1 {
		slice1[i] = rand.Intn(100)
	}
	fmt.Println(slice1)
	slice2 := slice1

	slice2[9] = 9999
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

	slice1 = append(slice1, 1111)
	slice2[0] = 2222
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

	slice3 := make([]int, 4, 5)
	slice3[0], slice3[1], slice3[2], slice3[3] = 1, 2, 3, 4
	slice4 := slice3 // only headers are copied
	slice3 = MultiplyByElement(slice3, 5)
	slice4 = slice3

	fmt.Println(slice3, "\n", slice4)

	slice5 := make([]int, 5) // in order to call copy the dest has to be instantiated
	copy(slice5, slice3)     // deep copy or clone
	fmt.Println(slice5)

	// convert array to a slice

	arr1 := [8]int{10, 11, 12, 13, 14, 15, 16, 17}

	//slice6 := arr1[:] // converting an array to a slice

	fmt.Println("All elements of arr1:", SumOf(arr1[:]))
	fmt.Println("All elements of arr1:", SumOf(arr1[5:]))
	fmt.Println("All elements of arr1:", SumOf(arr1[:5]))
	fmt.Println("All elements of arr1:", SumOf(arr1[4:8]))

	// can check nil with a slice

	var slice7 []int // make or []int{12,2133,23}
	if slice7 == nil {
		fmt.Println("nil slice7")
	}
}

// whenever append func is used then return that slice as a return type

func MultiplyByElement(slice []int, mul int) []int {
	for i, v := range slice {
		slice[i] = v * mul
	}
	slice = append(slice, mul) // Dangerous
	return slice
}

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
