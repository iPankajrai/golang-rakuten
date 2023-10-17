package main

import (
	"errors"
	"fmt"
	"reflect"
	_ "time"
)

func main() {
	a1 := Area(10.12, 13.45)
	p1 := Perimeter(10.12, 13.45)
	fmt.Println("Area:", a1)
	fmt.Println("Petimeter:", p1)

	a2, p2, err := AreaAndPerimeter(10.12, 11.)
	if err != nil {
		fmt.Println(err)
		return
	}
	println("Area:", a2, "Perimeter:", p2)

	a3, _, err := AreaAndPerimeter(10.12, 11.)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Area:", a3)

	if sum1, err := add(10, 12); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum1)
	}

	if sum1, err := add(10., 12.); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum1)
	}

	var a, b float32 = 10., 11.2

	if sum1, err := add(a, b); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum1)
	}

	if sum1, err := add(true, b); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum1)
	}

	if sum1, err := add("hello", "Rakuten"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum1)
	}
}

func Area(l, b float32) float32 {
	return l * b
}

func Perimeter(l, b float32) (p float32) {
	p = 2 * (l + b)
	return p
}

func AreaAndPerimeter(l, b float32) (float32, float32, error) {
	if l == 0 || b == 0 {
		//return 0, 0, errors.New("length or bredth is zero;invalid length or bredth")
		return 0, 0, fmt.Errorf("length or bredth is zero;invalid length or bredth")
	}

	return Area(l, b), Perimeter(l, b), nil
}

// a , b must of same type

func add(a, b any) (float64, error) {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return 0, errors.New("a and b must be of same type")
	}
	var sum float64
	switch a.(type) {
	case int:
		sum = float64(a.(int) + b.(int))
	case float32:
		sum = float64(a.(float32) + b.(float32))
	case float64:
		sum = a.(float64) + b.(float64)
	default:
		return 0, errors.New("invalid type")
	}

	return sum, nil

}
