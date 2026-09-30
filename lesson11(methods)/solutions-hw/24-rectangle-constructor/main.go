package main

import (
	"errors"
	"fmt"
)

type Rectangle struct {
	width, height float64
}

func NewRectangle(width, height float64) (*Rectangle, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("ширина и высота должны быть положительными")
	}
	return &Rectangle{width: width, height: height}, nil
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	_, err := NewRectangle(-5, 10)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	rect, err := NewRectangle(4, 6)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Площадь:", rect.Area())
	fmt.Println("Периметр:", rect.Perimeter())
}
