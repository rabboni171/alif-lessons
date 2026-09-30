package main

import (
	"errors"
	"fmt"
)

var ErrUserExists = errors.New("логин уже занят")

func Register(login string, users map[string]bool) error {
	if users[login] {
		return ErrUserExists
	}
	users[login] = true
	return nil
}

func RegisterWithRetry(baseLogin string, users map[string]bool) string {
	login := baseLogin
	suffix := 2

	for {
		err := Register(login, users)
		if err == nil {
			return login
		}
		if !errors.Is(err, ErrUserExists) {
			return login
		}

		login = fmt.Sprintf("%s_%d", baseLogin, suffix)
		suffix++
	}
}

func main() {
	users := map[string]bool{
		"aziz":   true,
		"aziz_2": true,
		"aziz_3": true,
	}

	login := RegisterWithRetry("aziz", users)
	fmt.Println("зарегистрирован как:", login)
}
