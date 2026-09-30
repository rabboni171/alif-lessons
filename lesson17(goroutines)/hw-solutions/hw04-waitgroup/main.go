package main

import (
	"fmt"
	"sync"
	"time"
)

func printSlowly(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Начали:", name)
	time.Sleep(1 * time.Second)
	fmt.Println("Готово:", name)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printSlowly("Аня", &wg)

	wg.Add(1)
	go printSlowly("Боря", &wg)

	wg.Add(1)
	go printSlowly("Вика", &wg)

	wg.Wait()
}
