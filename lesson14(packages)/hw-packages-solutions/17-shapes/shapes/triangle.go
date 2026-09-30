package shapes

import "math"

// Triangle - треугольник, заданный длинами трёх сторон.
type Triangle struct {
	A, B, C float64
}

// Area считает площадь по формуле Герона.
func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}
