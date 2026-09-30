package main

import (
	"fmt"

	// Один импорт пакета shapes, хотя внутри него три файла и три типа -
	// компилятор сам собирает их все в единый пакет.
	"shapes-demo/shapes"
)

func main() {
	circle := shapes.Circle{Radius: 3}
	rectangle := shapes.Rectangle{Width: 4, Height: 5}
	triangle := shapes.Triangle{A: 3, B: 4, C: 5}

	fmt.Printf("круг: %.2f\n", circle.Area())
	fmt.Printf("прямоугольник: %.2f\n", rectangle.Area())
	fmt.Printf("треугольник: %.2f\n", triangle.Area())
}
