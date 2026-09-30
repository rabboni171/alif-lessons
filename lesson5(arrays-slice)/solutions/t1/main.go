// Задача 1: создать срез из 10 чисел и вывести только те, что больше среднего.
package main

import "fmt"

func main() {
	nums := []int{4, 8, 15, 16, 23, 42, 3, 9, 12, 30}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	avg := float64(sum) / float64(len(nums))

	fmt.Println("срез:", nums)
	fmt.Println("среднее:", avg)

	fmt.Println("больше среднего:")
	for _, n := range nums {
		if float64(n) > avg {
			fmt.Println(n)
		}
	}
}
