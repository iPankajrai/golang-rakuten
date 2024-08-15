package main

import (
	"fmt"
	"strings"
)

func main() {

	var map1 map[string]int // what type can be a key
	// comparable .. any type that satisfies the comparable interface that can be used as a key
	// == opeartor can be used that can be a key

	if map1 == nil {
		fmt.Println("nil map")
	}

	map1 = make(map[string]int, 10)

	map1["india"] = 91
	map1["pakisthan"] = 92
	map1["banglasesh"] = 93

	for key, value := range map1 {
		fmt.Println("Key:", key, " Value:", value)
	}

	delete(map1, "pakisthan")

	for key, value := range map1 {
		fmt.Println("Key:", key, " Value:", value)
	}

	v, okay := map1["india"]
	if !okay {
		fmt.Println("no key")
	} else {
		fmt.Println(v)
	}

}

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	//fmt.Printf("Type of fields is %T\n", fields)

	wordMap := make(map[string]int)
	for _, wordStr := range fields {
		//fmt.Printf("value at index %d is %v\n", wordIndex, wordStr)
		_, ok := wordMap[wordStr]
		if ok {
			wordMap[wordStr] += 1
		} else {
			wordMap[wordStr] = 1
		}
	}
	return wordMap
}

//Mutating maps
/*
Mutating Maps
Insert or update an element in map m:

m[key] = elem
Retrieve an element:

elem = m[key]
Delete an element:

delete(m, key)
Test that a key is present with a two-value assignment:

elem, ok = m[key]
If key is in m, ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type.

Note: If elem or ok have not yet been declared you could use a short declaration form:

elem, ok := m[key]

*/
