package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Введите 3 числа (через пробел)")
	fmt.Scan(&a, &b, &c)

	if a > b {
		a, b = b, a
	}

	if b > c {
		b, c = c, b
	}

	if a > b {
		a, b = b, a
	}

	fmt.Println("Сумма наибольших чисел:", b + c)
}
