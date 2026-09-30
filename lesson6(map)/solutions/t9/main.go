// Задача 9: уникальные посетители через map[string]bool.
package main

import "fmt"

func main() {
	visitors := []string{"Али", "Вера", "Али", "Тимур", "Вера", "Али"}

	unique := make(map[string]bool)
	for _, v := range visitors {
		unique[v] = true
	}

	fmt.Println("Уникальных посетителей:", len(unique))
}
