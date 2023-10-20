package main

import (
	"fmt"
	"math/rand"
)

type IArea interface {
	Area() float32
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

type Square float32

func (s Square) Area() float32 {
	return float32(s * s)
}

func main() {

	var iarea1 IArea
	r1 := &Rect{L: 100., B: 101.5}
	var s1 Square = 25.25

	num := rand.Intn(100)
	if num%2 == 0 {
		iarea1 = r1
	} else {
		iarea1 = s1
	}

	a1 := iarea1.Area()
	//a1 := r1.Area()
	fmt.Println("Area:", a1)

}

// compile time invocation
// runtime invocation
