package main

import (
	"fmt"
	"sync"
	"time"
)

func greet(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Привет, %s!\n", name)
}

func main() {
	var wg sync.WaitGroup

	names := []string{"Аня", "Боря", "Вика", "Гена", "Даша"}

	for _, name := range names {
		wg.Add(1)
		go greet(name, &wg)
	}

	wg.Wait()
}
