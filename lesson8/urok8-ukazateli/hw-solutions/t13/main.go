// Задача 13: изменение температуры через указатель с ограничением по диапазону.
package main

import "fmt"

func adjustTemperature(current *float64, delta, min, max float64) bool {
	newValue := *current + delta

	if newValue > max {
		*current = max
		return true
	}
	if newValue < min {
		*current = min
		return true
	}

	*current = newValue
	return false
}

func main() {
	temp := 20.0
	min, max := 15.0, 30.0

	clamped := adjustTemperature(&temp, 5, min, max)
	fmt.Println("Температура:", temp, "| ограничили:", clamped) // 25, false

	clamped = adjustTemperature(&temp, 10, min, max)
	fmt.Println("Температура:", temp, "| ограничили:", clamped) // 30, true

	clamped = adjustTemperature(&temp, -20, min, max)
	fmt.Println("Температура:", temp, "| ограничили:", clamped) // 15, true
}
