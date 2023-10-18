package main

import "fmt"

func main() {

	var num1 int = 100
	var ptr1 *int = &num1

	var ptr11 **int = &ptr1

	fmt.Println("Value of ptr11:", **ptr11)
	var ptr111 ***int = &ptr11
	fmt.Println("Value of ptr111:", ***ptr111)
	fmt.Println("Value of num1:", num1, "Address of num1:", &num1)
	fmt.Println("Value of num1 thru ptr1:", *ptr1, "Address of num1 thru ptr1:", ptr1)

	incr(num1)

	println(num1)

	incrPtr(&num1)
	println(num1)

	ptr2 := square(100)
	fmt.Println(*ptr2)

	var num2 int = 25

	var ptr3 *int = &num2

	squareParam(num2, ptr3)
	fmt.Println(*ptr3)

	var ptr4 *int
	squareParam(20, ptr4)
	if ptr4 != nil {
		fmt.Println(*ptr4)
	}

	ptr5 := new(int)
	squareParam(20, ptr5)
	fmt.Println(*ptr5)
}

func incr(num int) {
	num++
}

func incrPtr(num *int) {
	*num++
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
