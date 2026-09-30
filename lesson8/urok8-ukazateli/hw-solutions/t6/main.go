// Задача 6: премия к зарплате через указатель.
package main

import "fmt"

func applyBonus(salary *float64, bonus float64) {
	*salary += bonus
}

func main() {
	salary := 5000.0
	applyBonus(&salary, 750)
	fmt.Printf("Зарплата с премией: %.2f\n", salary)
}
