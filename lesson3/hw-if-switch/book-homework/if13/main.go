package main

import "fmt"

func main() {
	var a, b, c int

	var middle int

	fmt.Println("Введите 3 числа (через пробел)")
	fmt.Scan(&a, &b, &c)

	if (a >= b && a <= c) || (a <= b && a >=c ) {
		middle = a
	} else if (b >= a && b <= c) || (b <= a && b >= c ) {
		middle = b
	} else {
		middle = c
	}

	fmt.Println("Среднее число", middle)
}
