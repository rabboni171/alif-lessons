package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("нельзя извлечь корень из отрицательного числа")
	}
	return math.Sqrt(x), nil
}

func main() {
	values := []float64{4, 9, -1, 16, -25}

	for _, v := range values {
		result, err := Sqrt(v)
		if err != nil {
			fmt.Printf("Sqrt(%.0f): ошибка: %v\n", v, err)
			continue
		}
		fmt.Printf("Sqrt(%.0f) = %.2f\n", v, result)
	}
}
