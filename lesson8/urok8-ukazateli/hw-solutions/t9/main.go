// Задача 9: безопасное деление - указатель на результат или nil при делении на 0.
package main

import "fmt"

func safeDivide(a, b int) *int {
	if b == 0 {
		return nil
	}
	result := a / b
	return &result
}

func main() {
	if r := safeDivide(10, 2); r != nil {
		fmt.Println("10 / 2 =", *r)
	}

	if r := safeDivide(5, 0); r != nil {
		fmt.Println("5 / 0 =", *r)
	} else {
		fmt.Println("Деление на 0 невозможно")
	}
}
