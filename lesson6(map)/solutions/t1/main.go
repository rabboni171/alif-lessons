// Задача 1: через литерал создать map с оценками учеников.
package main

import "fmt"

func main() {
	grades := map[string]int{
		"Али":   90,
		"Вера":  85,
		"Тимур": 78,
		"Вова":  95,
	}

	fmt.Println(grades)
	fmt.Println("Оценка Тимура:", grades["Тимур"])
}
