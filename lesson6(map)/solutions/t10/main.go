// Задача 10: добавить студента в группу через append.
package main

import "fmt"

func main() {
	groups := map[string][]string{
		"Группа A": {"Али", "Вера"},
		"Группа B": {"Тимур", "Вова"},
	}

	groups["Группа A"] = append(groups["Группа A"], "Наргис")

	for g, students := range groups {
		fmt.Printf("%s (%d): %v\n", g, len(students), students)
	}
}
