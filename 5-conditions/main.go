package main

import "strings"

func main() {
	num := 101

	if num%2 == 0 {
		println(num, "even number")
		return
	} else {
		println(num, "is odd number")
	}

	if num1 := 1001; num1%2 == 0 {
		println(num1, "even number")
	} else {
		println(num1, "is odd number")
	}

	if (true && true) || (true || false) {
		println("if is executed")
	} else {
		println("else is executed")
	}

	// State Age Height   	Gender  Ticket
	// Ka         		  	F       No ticket
	// Ka    <6   <3Feet  	M       No ticket
	// AP    <5   <3.5Feet  F|M     No ticket
	// TN                    F      No ticket

	var age uint8 = 5
	var height float32 = 3.
	state, gender := "AP", 'F'
	println(gender)
	if age <= 5 && height <= 3.5 && strings.ToUpper(state) == "AP" {
		println("no ticket for AP")
	} else if (strings.ToUpper(state) == "KA" || strings.ToUpper(state) == "TN") && gender == 'F' {
		println("no ticket for TN")
	} else {

	}

}

//true || true
