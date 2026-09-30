package main

import "fmt"

func main() {
	var (
		number1 int
		number2 int
	)

	fmt.Println("Введите первое число:")
	fmt.Scan(&number1)

	fmt.Println("Введите второе число:")
	fmt.Scan(&number2)

	if number1 < number2 {
		fmt.Println("Порядковый номер наименьшего числа: 1")
	} else {
		fmt.Println("Порядковый номер наименьшего числа: 2")
	}
}
