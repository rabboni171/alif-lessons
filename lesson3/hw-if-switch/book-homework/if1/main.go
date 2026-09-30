package main

import "fmt"

func main() {
	var number int

	fmt.Println("Введите число:")
	fmt.Scan(&number)

	if number > 0 {
		number += 1
	}

	fmt.Println("После проверки ваше число равняется:", number)
}
