package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p Point) DistanceTo(other Point) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

type Polygon []Point

func (poly Polygon) Perimeter() float64 {
	total := 0.0
	n := len(poly)
	for i := 0; i < n; i++ {
		next := (i + 1) % n
		total += poly[i].DistanceTo(poly[next])
	}
	return total
}

func main() {
	square := Polygon{
		{X: 0, Y: 0},
		{X: 4, Y: 0},
		{X: 4, Y: 4},
		{X: 0, Y: 4},
	}

	fmt.Printf("Периметр квадрата: %.2f\n", square.Perimeter())
}
