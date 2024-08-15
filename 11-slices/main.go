package main

import "fmt"

func main() {

	var slice1 []int        // this is only the declaration of slice
	slice1 = make([]int, 3) // len(slice1)=3
	fmt.Println(slice1)
	fmt.Println("Length:", len(slice1), "Capacity:", cap(slice1), "Address:", &slice1[0])
	slice1[0] = 100
	slice1[1] = 200
	slice1[2] = 300
	slice1 = append(slice1, 400)
	fmt.Println(slice1)
	fmt.Println("Length:", len(slice1), "Capacity:", cap(slice1), "Address:", &slice1[0])
	slice1 = append(slice1, 500)
	fmt.Println("Length:", len(slice1), "Capacity:", cap(slice1), "Address:", &slice1[0])
	slice1 = append(slice1, 600)
	fmt.Println("Length:", len(slice1), "Capacity:", cap(slice1), "Address:", &slice1[0])
	slice1 = append(slice1, 700)
	fmt.Println("Length:", len(slice1), "Capacity:", cap(slice1), "Address:", &slice1[0])
}
