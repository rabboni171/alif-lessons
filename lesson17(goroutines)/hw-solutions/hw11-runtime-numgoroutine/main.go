package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	fmt.Println("До запуска горутин:", runtime.NumGoroutine())

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
		}()
	}

	fmt.Println("Сразу после запуска:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("После wg.Wait():", runtime.NumGoroutine())
}
