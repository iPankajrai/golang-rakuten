package main

import (
	"fmt"
	"reflect"
)

// global variables

var num int

func main() {
	var num1 int = 100
	var num2 uint8
	fmt.Println("Num1:", num1)
	//fmt.Println("Num2:", num2)
	fmt.Printf("Num2:%d", num2)

	var (
		age    uint8   = 38
		num3   int     = 100
		ok1    bool    = true
		float1 float32 = 123.1231
		str1   string  = "Hello Rakuten"
	)

	var age2, num4, ok2, float2, str2 = 38, 100, true, 123.1231, "Hello Rakuten"

	str3 := "Hello Rakuten"
	float3 := 123.1231
	fmt.Println(num3, num4, float1, float2, float3, str1, str2, str3, ok1, ok2, age, age2)

	fmt.Printf("\nValue of num3 %d, Type of num3:%T\n", num3, num3)
	fmt.Println("Value of ok1:", ok1, "Type of ok1:", reflect.TypeOf(ok1))

	// work with chars

	str4 := "ハローワールド"
	fmt.Println(str4)

	rune1 := 'ハ'
	var rune2 int32 = 'ハ'
	var rune3 rune = 'ハ'
	fmt.Println(rune1, rune2, rune3)

	const pi1 float32 = 3.14
	//var days1 uint8 = 8
	const (
		pi2      = 3.145
		days     = 0 + 7
		noofdays = (1*days)/days + 7
		//noofdays2= 1 * days1
		company = "Rakuten"
	)
	const pi3 = 3.14 // short hand declaration of constant
	fmt.Println("pi2:", pi2, "type of pi2:", reflect.TypeOf(pi2))

}

// numbers
// int, uint, uint8,uint16,uint32,uint64, int8,int16,int32,int64
// float32, float64
// byte - 1 byte alias uint8
// rune - 4 bytes alias int32

// strings
// string

// boolean
// bool - 1 byte

// complex types
// complex64, complex128

// format verbs
// %d
// %v - value of a variable
// %s
// %T - type of a variable
