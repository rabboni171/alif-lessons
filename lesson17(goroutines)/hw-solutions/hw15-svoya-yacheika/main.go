package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	results := make([]int, 1000)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i] = 1 // каждая горутина трогает только свою ячейку
		}(i)
	}

	wg.Wait()

	sum := 0
	for _, v := range results {
		sum += v
	}

	fmt.Println("сумма:", sum) // теперь всегда ровно 1000
}
