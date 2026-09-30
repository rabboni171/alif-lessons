package main

import "fmt"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("нельзя разделить %.2f на %.2f", a, b)
	}
	return a / b, nil
}

func main() {
	pairs := [][2]float64{
		{10, 2},
		{5, 0},
		{9, 3},
		{7, 0},
	}

	for _, p := range pairs {
		result, err := Divide(p[0], p[1])
		if err != nil {
			fmt.Println("ошибка:", err)
			continue
		}
		fmt.Printf("%.2f / %.2f = %.2f\n", p[0], p[1], result)
	}
}
