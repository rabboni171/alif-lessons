package main

import (
	"fmt"

	"github.com/google/uuid"
)

// User - пользователь с уникальным ID.
type User struct {
	ID   string
	Name string
}

// NewUser создаёт пользователя, ID генерируется через внешний пакет uuid.
func NewUser(name string) User {
	return User{ID: uuid.NewString(), Name: name}
}

func main() {
	users := []User{
		NewUser("Али"),
		NewUser("Фарход"),
		NewUser("Азиза"),
	}

	for _, u := range users {
		fmt.Printf("%s - %s\n", u.ID, u.Name)
	}
}
