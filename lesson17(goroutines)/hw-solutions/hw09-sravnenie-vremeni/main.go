package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Обычный цикл, без горутин
	startObychno := time.Now()
	for id := 1; id <= 8; id++ {
		time.Sleep(time.Duration(id) * 50 * time.Millisecond)
	}
	obychnoeVremya := time.Since(startObychno)

	// Тот же цикл, но с горутинами
	var wg sync.WaitGroup
	startGorutiny := time.Now()

	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 50 * time.Millisecond)
			fmt.Printf("Горутина %d завершена\n", id)
		}(i)
	}

	wg.Wait()
	gorutinyVremya := time.Since(startGorutiny)

	fmt.Println("Обычным циклом:", obychnoeVremya)
	fmt.Println("С горутинами:  ", gorutinyVremya)
}
