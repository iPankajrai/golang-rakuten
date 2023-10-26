package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func(*sync.WaitGroup) {
		//wg.Add(1)
		for i := 1; i <= 5; i++ {
			fmt.Println("I-->", i)
		}
		wg.Done()
	}(wg)

	go func(*sync.WaitGroup) {
		//wg.Add(1)
		for i := 1; i <= 5; i++ {
			fmt.Println("J-->", i)
		}
		wg.Done()
	}(wg)

	go func(*sync.WaitGroup) {
		//wg.Add(1)
		for i := 1; i <= 5; i++ {
			fmt.Println("K-->", i)
		}
		wg.Done()
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		Even(10, 20)
		wg.Done()
	}(wg)

	for i := 1; i <= 5; i++ {
		fmt.Println("L-->", i)
	}

	wg.Wait() // This waits untile the counter of waitgroup becomes zero

}

// WaitGroup
//

func EvenP(wg *sync.WaitGroup, i, e int) {
	wg.Add(1)
	if i <= e {
		for k := i; k <= e; k++ {
			if k%2 == 0 {
				fmt.Println("Even number:", k)
			}
		}
		wg.Done()
	}
}

func Even(i, e int) {
	if i <= e {
		for k := i; k <= e; k++ {
			if k%2 == 0 {
				fmt.Println("Even number:", k)
			}
		}
	}
}
