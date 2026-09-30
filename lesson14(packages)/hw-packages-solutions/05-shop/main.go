package main

import (
	"fmt"

	"shop/catalog"
)

func main() {
	products := []catalog.Product{
		{Name: "Хлеб", Price: 4.5},
		{Name: "Молоко", Price: 8},
		{Name: "Сыр", Price: 35.9},
	}

	fmt.Println("Общая стоимость:", catalog.TotalPrice(products))
}
