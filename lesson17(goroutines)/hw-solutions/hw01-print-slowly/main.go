package main

import (
	"fmt"
	"time"
)

func printSlowly(name string) {
	fmt.Println("Начали:", name)
	time.Sleep(1 * time.Second)
	fmt.Println("Готово:", name)
}

func main() {
	start := time.Now()

	printSlowly("Аня")
	printSlowly("Боря")
	printSlowly("Вика")

	fmt.Println("Всего заняло:", time.Since(start)) // ~3 секунды
}
