package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = i + 1
	}

	var wg sync.WaitGroup
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			SquareNumber(num)
		}(num)
	}
	wg.Wait()
}

func SquareNumber(a int) {
	fmt.Printf("Square of %d is %d\n", a, a*a)
}
