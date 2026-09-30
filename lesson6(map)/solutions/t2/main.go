// Задача 2: пустая map через make, добавление цен товаров через присваивание.
package main

import "fmt"

func main() {
	prices := make(map[string]float64)

	prices["хлеб"] = 5.5
	prices["молоко"] = 12
	prices["сыр"] = 45
	prices["яйца"] = 20

	fmt.Println(prices)
	fmt.Println("Товаров в map:", len(prices))
}
