// Задача 8: разбить срез на куски по N элементов.
package main

import "fmt"

func chunk(nums []int, size int) [][]int {
	var result [][]int
	for size < len(nums) {
		result = append(result, nums[:size:size])
		nums = nums[size:]
	}
	return append(result, nums)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums, "по 2 ->", chunk(nums, 2))
}
