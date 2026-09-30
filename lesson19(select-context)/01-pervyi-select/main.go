package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch1 <- "сообщение из ch1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "сообщение из ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Получено:", msg)
		case msg := <-ch2:
			fmt.Println("Получено:", msg)
		}
	}

	// Обрати внимание: ch2 пришёл первым, потому что он быстрее.
	// select не ждёт по порядку - берёт готовое.
}
