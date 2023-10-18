package main

import "fmt"

func main() {

	var i int = 200

	var m1 myint = myint(i)

	var m2 myint1 = myint1(m1)

	var i2 int = int(m2)

	fmt.Println(i, m1, m2, i2)

	fmt.Println(myint(i).ToString())
	fmt.Println(m1.ToString())
	fmt.Println(myint(m2).ToString())

	fmt.Println(myint1(i).ToBytes())
	fmt.Println(myint1(m1).ToBytes())
	fmt.Println(m2.ToBytes())

	var f1 float64 = 10.23

	fmt.Println(myint(f1).ToString())
	fmt.Println(myint1(f1).ToBytes())

}

type myint int // creating a new type

type myint1 myint

func (mi myint1) ToBytes() []byte {
	return []byte(fmt.Sprint(mi))
}

func (mi myint) ToString() string {
	return fmt.Sprint(mi)
}
