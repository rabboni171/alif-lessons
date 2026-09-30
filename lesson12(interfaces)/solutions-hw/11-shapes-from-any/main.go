package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	items := []any{
		Circle{Radius: 2},
		42,
		Rectangle{Width: 3, Height: 4},
		"шум",
		Circle{Radius: 1},
		100,
	}

	total := 0.0
	for _, item := range items {
		shape, ok := item.(Shape)
		if !ok {
			continue
		}
		total += shape.Area()
	}

	fmt.Printf("Суммарная площадь фигур: %.2f\n", total)
}
