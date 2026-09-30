package main

import "fmt"

type Task struct {
	Title string
	Done  bool
}

func (t Task) String() string {
	if t.Done {
		return "[x] " + t.Title
	}
	return "[ ] " + t.Title
}

func main() {
	tasks := []Task{
		{Title: "Купить хлеб", Done: true},
		{Title: "Написать код", Done: false},
		{Title: "Сдать домашку", Done: true},
		{Title: "Погулять", Done: false},
		{Title: "Прочитать книгу", Done: false},
	}

	for _, t := range tasks {
		fmt.Println(t)
	}
}
