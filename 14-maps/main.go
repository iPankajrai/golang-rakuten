package main

import "fmt"

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
