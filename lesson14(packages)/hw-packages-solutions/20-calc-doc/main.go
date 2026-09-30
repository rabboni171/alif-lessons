package main

import (
	"fmt"

	"calc-doc/calc"
)

// Документацию пакета calc смотри командой `go doc ./calc` из корня
// этого модуля (или `go doc calc-doc/calc` из любого места) - doc-комментарии
// из calc/calc.go превращаются в читаемое описание пакета и его функций.

func main() {
	fmt.Println(calc.Add(2, 3))
	fmt.Println(calc.Subtract(5, 2))
	fmt.Println(calc.Multiply(4, 3))

	result, err := calc.Divide(10, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(result)
	}
}
