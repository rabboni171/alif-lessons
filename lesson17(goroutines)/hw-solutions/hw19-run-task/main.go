package main

import (
	"fmt"
	"sync"
	"time"
)

func runTask(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Задача", name, "началась")
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Задача", name, "закончена")
}

func main() {
	var wg sync.WaitGroup
	start := time.Now()

	tasks := []string{"отчёт", "письмо", "презентация", "бэкап", "проверка"}

	for _, task := range tasks {
		wg.Add(1)
		go runTask(task, &wg)
	}

	wg.Wait()

	fmt.Println("Выполнено задач:", len(tasks))
	fmt.Println("Всего заняло:", time.Since(start))
}
