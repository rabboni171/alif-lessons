// Задача 4: найти второе по величине число в срезе.
package main

import "fmt"

func main() {
	nums := []int{7, 2, 9, 4, 9, 1, 8}
	fmt.Println("срез:", nums)

	max, second := nums[0], nums[0]
	for _, n := range nums[1:] {
		switch {
		case n > max:
			second = max
			max = n
		case n > second && n != max:
			second = n
		}
	}

	fmt.Println("наибольшее:", max)
	fmt.Println("второе по величине:", second)
}
