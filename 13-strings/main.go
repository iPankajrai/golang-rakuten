package main

import "fmt"

func main() {
	str1 := "Hey ハローワールド"

	rstr1 := ""
	for i := 0; i < len(str1); i++ {
		fmt.Print(string(str1[i]), " ")
		rstr1 = string(str1[i]) + rstr1
	}
	fmt.Println("\n", rstr1)
	rstr2 := ""
	for _, v := range str1 {
		fmt.Print(string(v), " ")
		rstr2 = string(v) + rstr2
	}
	fmt.Println("\n", rstr2)

}
