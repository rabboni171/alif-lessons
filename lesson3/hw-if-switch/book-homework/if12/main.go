package main

import "fmt"

func main() {
	var a, b, c int

	var min int

	fmt.Println("Введите 3 числа (через пробел)")
	fmt.Scan(&a, &b, &c)

	if a <= b && a <= c {
		min = a
	} else if b <= a && b <= c {
		min = b
	} else {
		min = c
	}

	fmt.Println("Минимальное число", min)
}
