package main

import (
	"fmt"
	"sync"
	"time"
)

func countTo(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= n; i++ {
		fmt.Println(i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go countTo(5, &wg)

	wg.Add(1)
	go countTo(3, &wg)

	wg.Add(1)
	go countTo(7, &wg)

	wg.Wait()
}
