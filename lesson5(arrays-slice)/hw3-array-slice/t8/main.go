package main

import "fmt"

func main() {
	nums := make([]int, 0, 2) // len=0, cap=2

	for i := 1; i <= 10; i++ {
		nums = append(nums, i)
		fmt.Printf("len=%d cap=%d %v\n", len(nums), cap(nums), nums)
		// cap меняется скачком, когда места не хватает — Go выделяет новый массив побольше
	}
}
