package main

import (
	"fmt"
	"math"

	mygeometry "alias-demo/geometry"
)

func main() {
	radius := 5.0

	// Стандартный math.Pi и свой mygeometry.Pi - имена пакетов не совпадают
	// (math и mygeometry), поэтому оба используются в одном файле без конфликта.
	fmt.Println("math.Pi:", math.Pi)
	fmt.Println("mygeometry.Pi:", mygeometry.Pi)

	fmt.Println("площадь через math.Pi:", math.Pi*radius*radius)
	fmt.Println("площадь через mygeometry.CircleArea:", mygeometry.CircleArea(radius))
}
