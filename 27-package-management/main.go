package main

import (
	"fmt"
	"myshapes/rect"
	"myshapes/shape"
	"myshapes/square"

	rectangle "myshapes/shape/rect" // can alias or give another name to the package
)

func main() {

	r1, err := rect.New(10.12, 12.34)
	if err != nil {
		fmt.Println(err)
	}

	// s1 := rect.Square{S: 25.25}
	fmt.Println("Area of rect:", r1.Area())
	fmt.Println("Perimeter of rect:", r1.Perimeter())
	//fmt.Println("Area of rect:", s1.S)

	rect.PackageName = "rect pacakge"
	//rect.moduleName = "myshapes" // cannot  because m in moduleName is lowercase

	s1, _ := square.New(25.35)
	fmt.Println("Area of square:", s1.Area())
	fmt.Println("Perimeter of square:", s1.Perimeter())

	shape.Display()

	r2, err := rectangle.New(12.34, 14.56)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Area of rect:%.2f\n", r2.Area())
	fmt.Printf("Perimeter of rect:%.2f\n", r2.Perimeter())

}
