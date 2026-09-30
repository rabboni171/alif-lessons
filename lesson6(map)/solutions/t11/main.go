// Задача 11: найти сотрудника с максимальной зарплатой.
package main

import "fmt"

func main() {
	salaries := map[string]int{
		"Али":   3500,
		"Вера":  4200,
		"Тимур": 3900,
	}

	maxName, maxSalary := "", 0
	for name, s := range salaries {
		if s > maxSalary {
			maxName, maxSalary = name, s
		}
	}

	fmt.Printf("Больше всех получает %s: %d\n", maxName, maxSalary)
}
