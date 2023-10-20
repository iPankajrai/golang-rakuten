package main

import "fmt"

func main() {

	// i1 := new(Integer) // pointer

	// *i1 = 10 // deref

	// var iface ICalc = i1

	// iface.Add(11, 12, 13, 14, 15)
	// iface.Add(25)
	// fmt.Println(iface.Get())

	//iface.Add(10, 20).Add(30, 40).Sub(20, 3).Mul(12).Get()

	var i1 Integer = 100

	var iface ICalc = i1

	result := iface.Add(10).Add(20).Add(20, 30, 40).Add(1, 2, 3).Sub(10).Sub(6).Mul(2).Get()
	fmt.Println(result)

}

//	type ICalc interface {
//		Add(num ...int) int
//		//Sub(num ...int) int
//		//Mul(num ...int) int
//		Get() int
//	}
type Integer int

// func (i *Integer) Add(nums ...int) int {
// 	sum := int(*i)
// 	for _, v := range nums {
// 		sum += v
// 	}
// 	*i = Integer(sum)
// 	return sum
// }

// func (i *Integer) Get() int {
// 	return int(*i)
// }

type ICalc interface {
	Add(num ...int) ICalc
	Sub(num ...int) ICalc
	Mul(num ...int) ICalc
	Get() int
}

func (i Integer) Add(nums ...int) ICalc {
	sum := int(i)
	for _, v := range nums {
		sum += v
	}
	i = Integer(sum)
	return i
}

func (i Integer) Sub(nums ...int) ICalc {
	sub := int(i)
	for _, v := range nums {
		sub -= v
	}
	i = Integer(sub)
	return i
}

func (i Integer) Mul(nums ...int) ICalc {
	mul := int(i)
	for _, v := range nums {
		mul *= v
	}
	i = Integer(mul)
	return i
}

func (i Integer) Get() int {
	return int(i)
}
