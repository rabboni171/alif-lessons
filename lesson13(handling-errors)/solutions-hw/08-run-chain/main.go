package main

import (
	"errors"
	"fmt"
)

func readConfig() error {
	return errors.New("файл конфигурации не найден")
}

func startServer() error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("не удалось запустить сервер: %w", err)
	}
	return nil
}

func run() error {
	if err := startServer(); err != nil {
		return fmt.Errorf("ошибка при запуске приложения: %w", err)
	}
	return nil
}

func main() {
	err := run()

	fmt.Println("Полный текст ошибки:")
	fmt.Println(err)

	fmt.Println("\nЦепочка ошибок:")
	for e := err; e != nil; e = errors.Unwrap(e) {
		fmt.Println(e)
	}
}
