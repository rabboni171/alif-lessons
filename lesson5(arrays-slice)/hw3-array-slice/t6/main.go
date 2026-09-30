package main

import "fmt"

func main() {
	nums := []int{7, 2, 9, 4, 1, 8}

	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	fmt.Println("срез:", nums)
	fmt.Println("максимум:", max)
}
