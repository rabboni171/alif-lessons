package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Введите координаты A B C (через пробел)")
	fmt.Scan(&a, &b, &c)

	distanceFromB := math.Abs(a - b)
	distanceFromC := math.Abs(a - c)

	if distanceFromB < distanceFromC {
		fmt.Println("Точка B ближе к точке A и находится на расстоянии:", distanceFromB)
	} else {
		fmt.Println("Точка C ближе к точке A и находится на расстоянии:", distanceFromC)
	}
}
