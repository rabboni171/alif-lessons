package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Println("Введите целочисленные координаты X и Y (через пробел)")
	fmt.Scan(&x, &y)

	switch {
	case x == 0 && y == 0:
		fmt.Println("0")
	case x == 0:
		fmt.Println("1")
	case y == 0:
		fmt.Println("2")
	default:
		fmt.Println("3")
	}
}
