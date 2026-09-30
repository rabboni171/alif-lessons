package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []int
}

func NewQueue() *Queue {
	return &Queue{items: []int{}}
}

func (q *Queue) Enqueue(v int) {
	q.items = append(q.items, v)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, errors.New("очередь пуста")
	}
	first := q.items[0]
	q.items = q.items[1:]
	return first, nil
}

func main() {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	v, err := queue.Dequeue()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Достали из очереди:", v)
}
