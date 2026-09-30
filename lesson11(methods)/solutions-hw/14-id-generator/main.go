package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func NewIDGenerator() func() int {
	next := 1
	return func() int {
		id := next
		next++
		return id
	}
}

func main() {
	nextID := NewIDGenerator()

	names := []string{"Алишер", "Диана", "Марат", "Санжар"}

	users := make([]User, 0, len(names))
	for _, name := range names {
		users = append(users, User{ID: nextID(), Name: name})
	}

	for _, u := range users {
		fmt.Printf("#%d %s\n", u.ID, u.Name)
	}
}
