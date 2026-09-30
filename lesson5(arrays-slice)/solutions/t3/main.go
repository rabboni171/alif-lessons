// Задача 3: удалить из среза все нули.
package main

import "fmt"

func main() {
	nums := []int{1, 0, 2, 0, 3, 4, 0, 5}
	fmt.Println("до:", nums)

	// nums[:0] — срез с длиной 0, но тем же массивом под капотом,
	// поэтому новый массив не создаётся, просто перезаписываем его же память
	result := nums[:0]
	for _, n := range nums {
		if n != 0 {
			result = append(result, n)
		}
	}

	fmt.Println("после:", result)
}
