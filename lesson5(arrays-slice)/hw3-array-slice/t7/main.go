package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println("обратный порядок:")
	for i := len(nums) - 1; i >= 0; i-- {
		fmt.Println(nums[i]) // просто печатаем с конца, сам срез не трогаем
	}
}
