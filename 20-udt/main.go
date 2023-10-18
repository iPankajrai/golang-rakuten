package main

import "fmt"

func main() {

	var i int = 100
	var j interger = i
	var k myint = 200

	str1 := k.ToString()

	str2 := myint(i).ToString()

	str3 := myint(j).ToString()

	fmt.Println(str1, str2, str3)

	var mm1 mymap
	mm1 = make(map[string]any)
	mm1["Bangalore-1"] = 560086
	mm1["Bangalore-2"] = 560096
	mm1["Bangalore-3"] = 560034
	mm1["Bangalore-4"] = 560074
	keys := mm1.GetKeys()
	fmt.Println(keys)

	var mm2 map[string]any
	mm2 = make(map[string]any)
	mm2["Bangalore-1"] = 560086
	mm2["Bangalore-2"] = 560096
	mm2["Bangalore-3"] = 560034
	mm2["Bangalore-4"] = 560074
	keys2 := mymap(mm2).GetKeys()
	fmt.Println(keys2)

	e1 := empty{}
	e1.Greet()

}

type myint int // creating a new type

func (mi myint) ToString() string {
	return fmt.Sprint(mi)
}

type interger = int // aliasing the existing type

type mymap map[string]any

func (mm mymap) GetKeys() []string {
	keys := make([]string, len(mm))
	count := 0
	for k, _ := range mm {
		keys[count] = k
		count++
	}
	return keys
}

type empty struct{} // empty struct

func (e *empty) Greet() {
	fmt.Println("Hello Rakuten!")
}
