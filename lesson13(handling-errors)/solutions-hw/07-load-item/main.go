package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("элемент не найден")

func findItem(id int) error {
	return ErrNotFound
}

func loadItem(id int) error {
	if err := findItem(id); err != nil {
		return fmt.Errorf("не удалось загрузить item %d: %w", id, err)
	}
	return nil
}

func main() {
	err := loadItem(42)
	fmt.Println(err)

	if errors.Is(err, ErrNotFound) {
		fmt.Println("первопричина: элемент действительно не найден")
	}
}
