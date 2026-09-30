package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // без защиты - гонка данных
		}()
	}

	wg.Wait()
	fmt.Println("counter:", counter) // почти никогда не будет ровно 1000
}
