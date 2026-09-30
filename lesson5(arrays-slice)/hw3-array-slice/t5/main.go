package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for i := range nums {
		nums[i] *= 2 // меняем сам срез, а не копию
	}

	fmt.Println("результат:", nums)
}
