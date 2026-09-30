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
	go printSlowly("Аня")
	go printSlowly("Боря")
	go printSlowly("Вика")

	time.Sleep(2 * time.Second)

	// Минус: мы гадаем, сколько времени нужно горутинам, и либо ждём
	// слишком долго (тратим время впустую), либо слишком мало (часть
	// горутин не успеет доработать).
}
