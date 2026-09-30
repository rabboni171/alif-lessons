package main

import "fmt"

func main() {
	nums := []int{4, 8, 15, 16, 23, 42}

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	fmt.Println("срез:", nums)
	fmt.Println("сумма:", sum)
}
