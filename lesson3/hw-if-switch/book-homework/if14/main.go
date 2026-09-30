package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Введите 3 числа (через пробел)")
	fmt.Scan(&a, &b, &c)

	min, max := a, a

	if b < min {
		min = b
	}
	if b > max {
		max = b
	}

	if c < min {
		min = c
	}
	if c > max {
		max = c
	}

	fmt.Println("Наименьшее число:", min)
	fmt.Println("Наибольшее число:", max)
}
