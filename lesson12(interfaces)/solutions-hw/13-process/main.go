package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Product struct {
	Name  string
	Price float64
}

func Process(v any) {
	switch x := v.(type) {
	case Product:
		fmt.Printf("Товар: %s, цена: %.2f\n", x.Name, x.Price)
	case Shape:
		fmt.Printf("Площадь фигуры: %.2f\n", x.Area())
	default:
		fmt.Println("не знаю, что с этим делать")
	}
}

func main() {
	items := []any{
		Product{Name: "Хлеб", Price: 5},
		Circle{Radius: 2},
		42,
	}

	for _, item := range items {
		Process(item)
	}
}
