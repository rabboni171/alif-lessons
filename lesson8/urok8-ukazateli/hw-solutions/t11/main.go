// Задача 11: удвоение элементов массива через указатель на массив.
package main

import "fmt"

func doubleArray(arr *[5]int) {
	for i := range arr {
		arr[i] *= 2
	}
}

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("До:   ", numbers)

	doubleArray(&numbers)
	fmt.Println("После:", numbers)
}
