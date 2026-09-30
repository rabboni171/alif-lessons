package main

import (
	"fmt"
	"strconv"
)

func ParseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("не удалось разобрать возраст: %w", err)
	}
	return age, nil
}

func main() {
	values := []string{"25", "abc", "-3", "999"}

	for _, v := range values {
		age, err := ParseAge(v)
		if err != nil {
			fmt.Printf("%q: %v\n", v, err)
			continue
		}

		if age < 0 {
			fmt.Printf("%q: возраст не может быть отрицательным\n", v)
			continue
		}
		if age > 120 {
			fmt.Printf("%q: возраст слишком большой\n", v)
			continue
		}

		fmt.Printf("%q: возраст %d принят\n", v, age)
	}
}
