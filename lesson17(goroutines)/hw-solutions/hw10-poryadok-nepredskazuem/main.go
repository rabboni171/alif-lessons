package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Print(id, " ")
		}(i)
	}

	wg.Wait()
	fmt.Println()

	// Запусти программу несколько раз подряд - порядок чисел
	// будет каждый раз разным, потому что планировщик Go сам решает,
	// в каком порядке дать горутинам поработать.
}
