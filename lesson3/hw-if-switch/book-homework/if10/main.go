package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Введите значение a:")
	fmt.Scan(&a)

	fmt.Println("Введите значение b:")
	fmt.Scan(&b)



	if a != b {
		sum := a + b
		a, b = sum, sum
	} else {
		a, b = 0, 0
	}

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
