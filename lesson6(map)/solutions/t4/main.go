// Задача 4: увеличить возраст друга на 1 прямо в map.
package main

import "fmt"

func main() {
	ages := map[string]int{
		"Али":   25,
		"Вера":  30,
		"Тимур": 19,
	}

	ages["Тимур"]++ // у Тимура сегодня день рождения

	fmt.Println(ages)
}
