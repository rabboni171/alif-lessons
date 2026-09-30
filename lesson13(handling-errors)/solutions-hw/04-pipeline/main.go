package main

import "fmt"

func step1(n int) error {
	if n <= 0 {
		return fmt.Errorf("шаг 1: число должно быть положительным, получено %d", n)
	}
	return nil
}

func step2(n int) error {
	if n%3 == 0 {
		return fmt.Errorf("шаг 2: число %d делится на 3", n)
	}
	return nil
}

func step3(n int) error {
	if n > 100 {
		return fmt.Errorf("шаг 3: число %d слишком большое", n)
	}
	return nil
}

func Pipeline(n int) error {
	if err := step1(n); err != nil {
		return err
	}
	if err := step2(n); err != nil {
		return err
	}
	if err := step3(n); err != nil {
		return err
	}
	return nil
}

func main() {
	values := []int{5, -1, 9, 200, 7}

	for _, n := range values {
		err := Pipeline(n)
		if err != nil {
			fmt.Printf("Pipeline(%d): %v\n", n, err)
			continue
		}
		fmt.Printf("Pipeline(%d): успешно пройден\n", n)
	}
}
