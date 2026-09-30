package main

import "fmt"

func main() {
	var a, b, c, d int

	fmt.Println("Введите 4 числа, 3 из которых будут равны (через пробел)")
	fmt.Scan(&a, &b, &c, &d)

	switch {
	case a == b && b == c && c == d:
		fmt.Println("Все числа равны")
	case a == b && a == c:
		fmt.Println("Индекс отличающегося числа равен 4")
	case a == b && a == d :
		fmt.Println("Индекс отличающегося числа равен 3")
	case a == c && a == d:
		fmt.Println("Индекс отличающегося числа равен 2")
	case b == c && b == d:
		fmt.Println("Индекс отличающегося числа равен 1")
	default:
		fmt.Println("Три числа должно быть равны по значению :)")
	}
}
