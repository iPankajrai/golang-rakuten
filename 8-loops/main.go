package main

import "fmt"

func main() {

	// unconditional loop
	count := 0
	for {
		if count >= 10 {
			break
		}
		fmt.Println("Hello Rakuten->", count)
		count++
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			print(i, " ")
		}
	}

	println()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		print(i, " ")
	}
	println()
	for i, j := 1, 10; i <= 10 && j <= 10; i, j = j+1, j+1 {
		println(i, " ", j)
	}

	println()
	println()
a:
	//out := false
	for i := 1; i <= 7; i++ {
		// if out {
		// 	break
		// }
		for j := 1; j <= 10; j++ {
			if i == 5 && j == 5 {
				//out = true
				break a
			}
			println(i, " ", j)
		}
		println()
	}

	i := 1

	for ; i <= 10; i++ {
		// todo some thing

	}

	for i := 1; i <= 10; {
		// todo some thing

		i++

	}

	for i := 1; ; i++ {
		if i >= 10 {
			break
		}
		// todo some thing

	}

	j := 1
	for j <= 10 {

		// todo some thing
		j++
	}
}
