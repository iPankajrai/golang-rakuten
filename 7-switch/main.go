package main

import (
	"fmt"
	"reflect"
)

func main() {
	var any1 any // if the type is numeric then write the square of the value
	var any2 any = false
	any1 = any2
	switch v := any1.(type) {
	case int:
		fmt.Println(reflect.TypeOf(v))
		println(v * v)
	case uint:
		fmt.Println(reflect.TypeOf(v))
		println(v * v)
	case uint8:
		fmt.Println(reflect.TypeOf(v))
		println(v * v)
	case uint16:
		fmt.Println(reflect.TypeOf(v))
		println(v * v)
	case uint32:
		fmt.Println(reflect.TypeOf(v))
		println(v * v)
	case uint64:
		fmt.Println(reflect.TypeOf(v))
		println(v * v)
	case int8:
		fmt.Println(reflect.TypeOf(v))
		println(v * v)
	case int16:
		fmt.Println(reflect.TypeOf(v))
		println(v * v)
	case int32:
		fmt.Println(reflect.TypeOf(v))
		println(v * v)
	case int64:
		fmt.Println(reflect.TypeOf(v))
		println(v * v)
	case float32:
		fmt.Println(reflect.TypeOf(v))
		fmt.Println(v * v)
	case float64:
		fmt.Println(reflect.TypeOf(v))
		fmt.Println(v * v)
	default:
		fmt.Println(reflect.TypeOf(v))
		println("cannot perform square on this type")
	}
}
