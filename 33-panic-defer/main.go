package main

import "fmt"

func main() {

	fmt.Println("Main has started")

	func() {
		defer fmt.Println("Big anonymous func has ended")

		for a, b := 10, 9; a >= 0; a, b = a-1, b-1 {
			c := func(a, b int) int {
				fmt.Println("sub Anonymous func has started")
				return b / a
			}(a, b)

			fmt.Println("c:", c)
		}

		func() {
			for i := 1; i <= 100; i++ {
				fmt.Print(i, " ")
			}
		}()

	}()

	fmt.Println("Finally Main has ended")

}

// by default panic panics the whole call stack, therefore it is a cascading effect
// 1. panic is created by the runtime
// 2. panic can also be created by the user
