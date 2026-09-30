package main

import (
	"errors"
	"fmt"
)

type Command interface {
	Execute() (string, error)
}

type GreetCommand struct{}

func (c GreetCommand) Execute() (string, error) {
	return "Привет!", nil
}

type DivideCommand struct {
	A, B float64
}

func (c DivideCommand) Execute() (string, error) {
	if c.B == 0 {
		return "", errors.New("деление на ноль")
	}
	return fmt.Sprintf("%.2f", c.A/c.B), nil
}

func main() {
	commands := map[string]Command{
		"greet":     GreetCommand{},
		"divide":    DivideCommand{A: 10, B: 2},
		"divide_by": DivideCommand{A: 10, B: 0},
	}

	names := []string{"greet", "divide", "divide_by", "unknown"}

	for _, name := range names {
		cmd, ok := commands[name]
		if !ok {
			fmt.Printf("команда %q не найдена\n", name)
			continue
		}

		result, err := cmd.Execute()
		if err != nil {
			fmt.Printf("команда %q завершилась ошибкой: %v\n", name, err)
			continue
		}

		fmt.Printf("команда %q вернула: %s\n", name, result)
	}
}
