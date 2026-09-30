// Задача 7: посчитать, сколько раз каждое число встречается (без map).
package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 1, 4, 2, 3}

	var values []int // уникальные числа
	var counts []int // счётчик для каждого из values, по тому же индексу

	for _, n := range nums {
		found := false
		for i, v := range values {
			if v == n {
				counts[i]++
				found = true
				break
			}
		}
		if !found {
			values = append(values, n)
			counts = append(counts, 1)
		}
	}

	for i, v := range values {
		fmt.Printf("%d встречается %d раз(а)\n", v, counts[i])
	}
}
