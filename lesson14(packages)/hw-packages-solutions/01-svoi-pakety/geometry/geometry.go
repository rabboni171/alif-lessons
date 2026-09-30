package geometry

// RectangleArea считает площадь прямоугольника.
func RectangleArea(width, height float64) float64 {
	return width * height
}

// RectanglePerimeter считает периметр прямоугольника.
func RectanglePerimeter(width, height float64) float64 {
	return 2 * (width + height)
}
