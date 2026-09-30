package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Println("Введите 3 числа, 2 из которых будут равны (через пробел)")
	fmt.Scan(&a, &b, &c)

	switch {
	case a == b && b == c:
		fmt.Println("Все числа равны")
	case a == b:
		fmt.Println("Индекс отличающегося числа равен 3")
	case a == c:
		fmt.Println("Индекс отличающегося числа равен 2")
	case b == c:
		fmt.Println("Индекс отличающегося числа равен 1")
	default:
		fmt.Println("Два числа должно быть равны по значению :)")
	}
}
