// Package calc реализует простые арифметические операции: сложение,
// вычитание, умножение и деление с проверкой на ноль.
package calc

import "fmt"

// Add складывает два числа.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract вычитает b из a.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply умножает два числа.
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide делит a на b и возвращает ошибку, если b равно нулю.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("calc: деление на ноль")
	}
	return a / b, nil
}
