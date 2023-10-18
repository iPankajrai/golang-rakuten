package main

import "fmt"

func main() {
	r1 := Rect{L: 10., B: 12.}
	a1 := Area(r1)
	fmt.Println("Area of Rect:", a1)

	a2 := r1.Area()
	fmt.Println("Area of Rect:", a2)
	fmt.Println("Area of Rect in rect:", r1.A)
	r1.Incr()
	fmt.Println("I in r1:", r1.I)
	//(&r1).Area()
}

type Rect struct {
	L, B, A float32
	I       int
}

func (r Rect) Incr() { // normal receiver
	r.I++
	fmt.Println("Inside the rect:", r.I)
}
func (r *Rect) Area() float32 {
	r.A = r.L * r.B
	return r.A
}

func Area(r Rect) float32 {
	return r.L * r.B
}
