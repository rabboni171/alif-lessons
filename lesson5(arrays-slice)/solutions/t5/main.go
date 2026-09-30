// Задача 5: объединить два среза и отсортировать «пузырьком».
package main

import "fmt"

func main() {
	a := []int{5, 3, 8, 1}
	b := []int{9, 2, 7, 4}

	merged := append(a, b...)
	fmt.Println("объединили:", merged)

	// пузырьковая сортировка: вложенные циклы + обмен соседей
	for i := 0; i < len(merged)-1; i++ {
		for j := 0; j < len(merged)-1-i; j++ {
			if merged[j] > merged[j+1] {
				merged[j], merged[j+1] = merged[j+1], merged[j]
			}
		}
	}

	fmt.Println("отсортировали:", merged)
}
