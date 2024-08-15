package main

import "fmt"

// Define Stack Struct type : 1. items slice for keeping elements
type Stack struct {
	items []int
}

// Define push  func
func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

// Define pop func
func (s *Stack) pop() int {

	if len(s.items) == 0 {
		return -1
	}

	lastIndex := len(s.items)

	itmLi := s.items[lastIndex-1]

	s.items = s.items[:lastIndex-1]

	return itmLi

}

func main() {
	stack := Stack{}

	stack.push(4)

	stack.push(5)
	stack.push(9)
	stack.push(-32)
	stack.push(42)
	stack.push(82)

	fmt.Println(stack)

	fmt.Println(stack.pop())

}
