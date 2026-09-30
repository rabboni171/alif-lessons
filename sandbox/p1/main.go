package main

import (
	"errors"
	"fmt"
)

// не используя пакет errors(чтоб в import не было) создать собственный тип ошибки

type MyError struct {
	Message string
}

func (e MyError) Error() string {
	return e.Message
}

func someOperation(ok bool) error {
	if !ok {
		//return MyError{Message: "какая-то ошибка в базе данных например"}
		return errors.New("какая-то ошибка в базе данных например")
	}
	return nil
}

func doSomething() error {
	if err := someOperation(false); err != nil {
		return MyError{
			Message: "что-то пошло не так",
		}
	}
	return nil
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
	}
}
