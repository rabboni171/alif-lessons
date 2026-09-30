package main

import "fmt"

func main() {
	var a, b float64

	fmt.Println("Введите значение a:")
	fmt.Scan(&a)

	fmt.Println("Введите значение b:")
	fmt.Scan(&b)



	if a != b {
		a, b = b, a
	}

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
