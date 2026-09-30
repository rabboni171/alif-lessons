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

	go printSlowly("Аня")
	go printSlowly("Боря")
	go printSlowly("Вика")

	fmt.Println("Всего заняло:", time.Since(start))

	// Что видно в выводе: почти всегда ничего не печатается (или печатается
	// только часть). main() не ждёт горутины - как только main() дошёл до
	// конца, программа завершается, даже если горутины ещё не успели
	// отработать time.Sleep. При разных запусках вывод может отличаться.
}
