package main

import (
	"fmt"

	"taskmanager/task"
)

func main() {
	tasks := []task.Task{
		task.New("Сделать домашку по пакетам"),
		task.New("Сходить в магазин"),
		task.New("Прочитать главу учебника"),
		task.New("Позвонить преподавателю"),
	}

	tasks[0].Done = true
	tasks[2].Done = true

	for _, t := range tasks {
		mark := "[ ]"
		if t.Done {
			mark = "[x]"
		}
		fmt.Printf("%s %s\n", mark, t.Title)
	}
}
