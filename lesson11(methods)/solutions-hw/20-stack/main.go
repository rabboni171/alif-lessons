package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("стек пуст")
	}
	last := len(s.items) - 1
	v := s.items[last]
	s.items = s.items[:last]
	return v, nil
}

func (s Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := &Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	v, _ := stack.Pop()
	fmt.Println("Достали:", v)

	v, _ = stack.Pop()
	fmt.Println("Достали:", v)

	stack.Pop()

	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
}
