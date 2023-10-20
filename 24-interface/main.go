package main

import (
	"fmt"
)

type IArea interface {
	Area() float32
}

type IPerimeter interface {
	Perimeter() float32
}

type IAreaPerimeter interface {
	IArea
	IPerimeter
}

// class Rect extends IArea in Java
// class Rect: IArea in C#
// impl IArea for Rect  in Rust

type Rect struct {
	L, B float32
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}

func (r *Rect) Perimeter() float32 {
	return 2 * (r.L + r.B)
}

func (r *Rect) Display() {
	fmt.Println("Length:", r.L, "Bredth:", r.B)
}

type Square float32

func (s Square) Area() float32 {
	return float32(s * s)
}

func (s Square) Perimeter() float32 {
	return float32(4 * s)
}

type Cuboid struct {
	L, B, H float32
}

func (r *Cuboid) Area() float32 {
	return r.L * r.B * r.H
}

func (r *Cuboid) Perimeter() float32 {
	return 4 * (r.L + r.B + r.H)
}

func main() {

	shapes := make([]IAreaPerimeter, 5) // all the shapes that are passed here must implement IAreaPerimeter interface

	r1 := &Rect{L: 10.13, B: 13.45}

	r2 := &Rect{L: 12.14, B: 13.75}

	s1 := Square(20.25)
	s2 := Square(25.25)

	c1 := Cuboid{L: 12.45, B: 13.45, H: 10.78}

	shapes[0] = r1
	shapes[1] = r2
	shapes[2] = &s1
	shapes[3] = &s2
	shapes[4] = &c1
	//shapes[5] = &Circle{R: 25.} // since circle is not implementing the interface , it cannot be given to the slice

	for _, v := range shapes {
		if v == nil {
			fmt.Println("nil v")
		} else {
			AreaPetimer(v)
		}

	}

}

type Circle struct {
	R float32
}

func AreaPetimer(iap IAreaPerimeter) {
	fmt.Println("Area:", iap.Area())
	fmt.Println("Perimeter:", iap.Perimeter())
}
