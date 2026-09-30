// Задача 3: проверить через comma-ok, есть ли товар "сыр".
package main

import "fmt"

func main() {
	prices := map[string]float64{
		"хлеб":   5.5,
		"молоко": 12,
		"яйца":   20,
	}

	if price, ok := prices["сыр"]; ok {
		fmt.Println("Цена сыра:", price)
	} else {
		fmt.Println("такого товара нет")
	}
}
