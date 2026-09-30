package main

import (
	"fmt"

	// Анонимный (blank) импорт: пакет stats нигде в коде по имени не
	// используется, но его init() всё равно выполняется - импорт нужен
	// только ради побочного эффекта.
	_ "blank-import-demo/stats"
)

func main() {
	fmt.Println("main: программа работает")
}
