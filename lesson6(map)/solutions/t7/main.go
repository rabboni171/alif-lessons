// Задача 7: range по map — пары ключ-значение, потом только ключи.
package main

import "fmt"

func main() {
	grades := map[string]int{
		"Али":   90,
		"Вера":  85,
		"Тимур": 78,
	}

	for name, grade := range grades {
		fmt.Printf("%s: %d\n", name, grade)
	}

	fmt.Println("Ученики:")
	for name := range grades {
		fmt.Print(name, " ")
	}
	fmt.Println()
}
