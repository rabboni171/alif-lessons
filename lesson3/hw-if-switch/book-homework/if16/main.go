package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Println("Введите 3 числа (через пробел)")
	fmt.Scan(&a, &b, &c)

	if (a >= b && b >= c) {
		a *= 2
		b *= 2
		c *= 2
	} else {
		a = -a
		b = -b
		c = -c
	}

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
}
