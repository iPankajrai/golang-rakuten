package main

func main() {

	day := 1

	switch day {
	case 1:
		println("sunday")
	case 2:
		println("monday")
	default:
	}

	num := 32

	switch { // if empty swtch is used then case can have conditons
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	// case num%3 == 0:
	// 	println(num, "is divisible by 3")
	// 	fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

}
