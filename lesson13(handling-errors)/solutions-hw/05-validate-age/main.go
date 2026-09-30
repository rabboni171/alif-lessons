package main

import (
	"errors"
	"fmt"
)

var ErrNegativeAge = errors.New("возраст не может быть отрицательным")
var ErrTooOld = errors.New("возраст слишком большой")

func ValidateAge(age int) error {
	if age < 0 {
		return ErrNegativeAge
	}
	if age > 120 {
		return ErrTooOld
	}
	return nil
}

func main() {
	ages := []int{25, -5, 130, 60}

	for _, age := range ages {
		err := ValidateAge(age)
		switch {
		case err == nil:
			fmt.Printf("возраст %d корректен\n", age)
		case errors.Is(err, ErrNegativeAge):
			fmt.Printf("возраст %d: нельзя указать отрицательный возраст\n", age)
		case errors.Is(err, ErrTooOld):
			fmt.Printf("возраст %d: слишком много лет для этой системы\n", age)
		}
	}
}
