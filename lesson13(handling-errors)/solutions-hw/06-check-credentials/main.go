package main

import (
	"errors"
	"fmt"
)

var ErrEmptyLogin = errors.New("логин не может быть пустым")
var ErrShortPassword = errors.New("пароль должен быть не короче 6 символов")

func CheckCredentials(login, password string) error {
	if login == "" {
		return ErrEmptyLogin
	}
	if len(password) < 6 {
		return ErrShortPassword
	}
	return nil
}

func main() {
	cases := []struct {
		login    string
		password string
	}{
		{"aziz", "123456"},
		{"", "123456"},
		{"aziz", "123"},
	}

	for _, c := range cases {
		err := CheckCredentials(c.login, c.password)
		switch {
		case err == nil:
			fmt.Println("доступ разрешён")
		case errors.Is(err, ErrEmptyLogin):
			fmt.Println("нужно указать логин")
		case errors.Is(err, ErrShortPassword):
			fmt.Println("пароль слишком короткий")
		}
	}
}
