package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Введите значение a:")
	fmt.Scan(&a)

	fmt.Println("Введите значение b:")
	fmt.Scan(&b)



	if a > b {
		b = a	
	} else if a < b {
		a = b
	} else {
		a, b = 0, 0
	}

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
