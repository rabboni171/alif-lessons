package main

import "fmt"

type Task struct {
	Title    string
	Priority int
	Done     bool
}

func (t *Task) Complete() {
	t.Done = true
}

func printTasks(tasks []Task) {
	for _, t := range tasks {
		fmt.Printf("  %s (приоритет %d)\n", t.Title, t.Priority)
	}
}

func main() {
	tasks := []Task{
		{Title: "Помыть посуду", Priority: 3},
		{Title: "Сдать отчёт", Priority: 1},
		{Title: "Купить продукты", Priority: 4},
		{Title: "Позвонить клиенту", Priority: 2},
		{Title: "Полить цветы", Priority: 6},
		{Title: "Сходить в спортзал", Priority: 5},
	}

	fmt.Println("До сортировки:")
	printTasks(tasks)

	for i := 0; i < len(tasks); i++ {
		for j := 0; j < len(tasks)-i-1; j++ {
			if tasks[j].Priority > tasks[j+1].Priority {
				tasks[j], tasks[j+1] = tasks[j+1], tasks[j]
			}
		}
	}

	fmt.Println("После сортировки:")
	printTasks(tasks)

	tasks[0].Complete()
	fmt.Println("Выполнена:", tasks[0].Title, "— Done:", tasks[0].Done)
}
