package main

import (
	"errors"
	"fmt"
)

func ValidatePassword(password string) []error {
	var errs []error

	if len(password) < 8 {
		errs = append(errs, errors.New("пароль должен быть не короче 8 символов"))
	}

	hasDigit := false
	hasUpper := false
	for i := 0; i < len(password); i++ {
		c := password[i]
		if c >= '0' && c <= '9' {
			hasDigit = true
		}
		if c >= 'A' && c <= 'Z' {
			hasUpper = true
		}
	}

	if !hasDigit {
		errs = append(errs, errors.New("пароль должен содержать хотя бы одну цифру"))
	}
	if !hasUpper {
		errs = append(errs, errors.New("пароль должен содержать хотя бы одну заглавную букву"))
	}

	return errs
}

func main() {
	passwords := []string{"weak", "onlylower1", "Strong1Pass", "Short1"}

	for _, p := range passwords {
		errs := ValidatePassword(p)
		if len(errs) == 0 {
			fmt.Printf("%q: пароль надёжный\n", p)
			continue
		}

		fmt.Printf("%q:\n", p)
		for _, err := range errs {
			fmt.Println("  -", err)
		}
	}
}
