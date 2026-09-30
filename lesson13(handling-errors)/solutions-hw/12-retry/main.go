package main

import "fmt"

func Retry(attempts int, f func() error) error {
	var err error
	for i := 0; i < attempts; i++ {
		err = f()
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("все попытки исчерпаны: %w", err)
}

func main() {
	count := 0
	unstable := func() error {
		count++
		if count < 3 {
			return fmt.Errorf("попытка %d не удалась", count)
		}
		return nil
	}

	err := Retry(5, unstable)
	if err != nil {
		fmt.Println("ошибка:", err)
	} else {
		fmt.Println("успех после", count, "попыток")
	}
}
