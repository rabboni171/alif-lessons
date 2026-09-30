// Задача 2: развернуть срез задом наперёд (без создания нового).
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("до:", nums)

	// два указателя идут навстречу друг другу и меняют элементы местами
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

	fmt.Println("после:", nums)
}
