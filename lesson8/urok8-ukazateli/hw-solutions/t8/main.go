// Задача 8: поиск города по логину - указатель или nil, если не найдено.
package main

import "fmt"

func findCity(users map[string]string, login string) *string {
	if city, ok := users[login]; ok {
		return &city
	}
	return nil
}

func main() {
	users := map[string]string{
		"ali01":  "Ташкент",
		"vera02": "Самарканд",
	}

	if city := findCity(users, "ali01"); city != nil {
		fmt.Println("Город:", *city)
	} else {
		fmt.Println("Пользователь не найден")
	}

	if city := findCity(users, "unknown"); city != nil {
		fmt.Println("Город:", *city)
	} else {
		fmt.Println("Пользователь не найден")
	}
}
