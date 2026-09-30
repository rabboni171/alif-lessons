package main

import "fmt"

type Point struct {
	X, Y int
}

func Describe(v any) {
	fmt.Printf("значение: %v, тип: %T\n", v, v)
}

func main() {
	Describe(42)
	Describe("привет")
	Describe(3.14)
	Describe(true)
	Describe([]int{1, 2, 3})
	Describe(Point{X: 1, Y: 2})
}
