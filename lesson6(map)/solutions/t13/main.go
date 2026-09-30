// Задача 13: посчитать, сколько раз встречается каждый город в срезе.
package main

import "fmt"

func main() {
	cities := []string{"Душанбе", "Худжанд", "Душанбе", "Куляб", "Душанбе", "Худжанд"}

	counts := make(map[string]int)
	for _, c := range cities {
		counts[c]++
	}

	fmt.Println(counts)
}
