package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var num1 int8 = 100

	var num2 int64 = int64(num1) // type casting in go

	fmt.Println(num1, num2)

	float1 := 123.12312132132323123
	var float2 float32 = float32(float1)

	var num3 int64 = int64(float1)
	fmt.Println(num3, float2)

	num4 := 1231231212312312333
	var num5 uint16 = uint16(num4)
	fmt.Println(num5)

	//12495

	// ok := false
	//var num6 uint8 = uint8(ok)

	var num6 int = 12495
	fmt.Println(string(num6))

	// 1- Sprintf

	str2 := fmt.Sprint(num6) // formatter that formats any type of data in the form of string
	fmt.Println(str2, "-->", reflect.TypeOf(str2))

	str3 := strconv.Itoa(num6)
	fmt.Println(str3, "-->", reflect.TypeOf(str3))

	str4 := "95"
	var num7 int = int(str4[0])

	// num8, _ := strconv.Atoi(str4)
	num8, err := strconv.Atoi(str4)
	if err != nil {
		fmt.Println("--->>>>", err)
	} else {
		fmt.Println(num8)
	}
	fmt.Println(num7)

}
