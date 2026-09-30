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

type Triangle struct {
	A, B, C float64
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func main() {
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 5},
		Triangle{A: 3, B: 4, C: 5},
	}

	biggest := shapes[0]
	for _, s := range shapes {
		if s.Area() > biggest.Area() {
			biggest = s
		}
	}

	fmt.Printf("Фигура с максимальной площадью: %.2f\n", biggest.Area())

	switch shape := biggest.(type) {
	case Circle:
		fmt.Println("Это круг, радиус:", shape.Radius)
	case Rectangle:
		fmt.Println("Это прямоугольник, ширина:", shape.Width, "высота:", shape.Height)
	case Triangle:
		fmt.Println("Это треугольник, стороны:", shape.A, shape.B, shape.C)
	}
}
