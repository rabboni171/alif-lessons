package main

import "fmt"

func main() {
	nums := []int{4, 7, 10, 15, 22, 33, 40}

	even, odd := 0, 0
	for _, n := range nums {
		if n%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	fmt.Println("срез:", nums)
	fmt.Println("чётных:", even)
	fmt.Println("нечётных:", odd)
}
