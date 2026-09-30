package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doubled := make([]int, 10)

	for i, n := range numbers {
		wg.Add(1)
		go func(i, n int) {
			defer wg.Done()
			doubled[i] = n * 2
		}(i, n)
	}

	wg.Wait()

	fmt.Println("исходный:", numbers)
	fmt.Println("удвоенный:", doubled)
}
