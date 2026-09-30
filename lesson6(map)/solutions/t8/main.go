// Задача 8: вывести map с оценками в отсортированном по ключу порядке.
package main

import (
	"fmt"
	"sort"
)

func main() {
	grades := map[string]int{
		"Али":   90,
		"Вера":  85,
		"Тимур": 78,
		"Вова":  95,
	}

	keys := make([]string, 0, len(grades))
	for k := range grades {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, grades[k])
	}
}
