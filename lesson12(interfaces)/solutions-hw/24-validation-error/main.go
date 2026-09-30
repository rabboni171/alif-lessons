package main

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return "поле " + e.Field + ": " + e.Message
}

func ValidateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "возраст не может быть отрицательным"}
	}
	if age > 130 {
		return &ValidationError{Field: "age", Message: "возраст слишком большой"}
	}
	return nil
}

func main() {
	ages := []int{25, -5, 200, 40}

	for _, age := range ages {
		err := ValidateAge(age)
		if err == nil {
			fmt.Printf("возраст %d прошёл проверку\n", age)
			continue
		}

		validationErr, ok := err.(*ValidationError)
		if ok {
			fmt.Printf("ошибка в поле %q: %v\n", validationErr.Field, err)
		} else {
			fmt.Println("ошибка:", err)
		}
	}
}
