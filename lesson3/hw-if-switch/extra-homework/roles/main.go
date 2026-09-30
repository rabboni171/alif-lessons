package main

import (
	"fmt"
	"strings"
)

func main() {
	type Role string

	const (
		ADMIN   Role = "АДМИНИСТРАТОР"
		MANAGER Role = "МЕНЕДЖЕР"
		USER    Role = "ПОЛЬЗОВАТЕЛЬ"
		GUEST   Role = "ГОСТЬ"
	)

	var role string

	fmt.Print("Введите вашу роль в системе и мы укажем ваши права (администратор, менеджер, пользователь, гость): ")
	fmt.Scan(&role)

	roleInUpperCase := strings.ToUpper(role)

	if roleInUpperCase == "АДМИН" {
		roleInUpperCase = "АДМИНИСТРАТОР"
	}

	userRole := Role(roleInUpperCase)

	switch userRole {
	case ADMIN:
		fmt.Println("Все права доступа")
	case MANAGER:
		fmt.Println("Управление пользователями")
	case USER:
		fmt.Println("Обычный доступ")
	case GUEST:
		fmt.Println("Только просмотр")
	default:
		fmt.Println("Такой роли нет в системе")
	}
}
