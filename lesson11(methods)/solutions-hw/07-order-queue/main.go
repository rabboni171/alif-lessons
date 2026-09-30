package main

import (
	"errors"
	"fmt"
)

type Order struct {
	ID     int
	Amount float64
}

type OrderQueue struct {
	orders []Order
}

func (q *OrderQueue) Enqueue(o Order) {
	q.orders = append(q.orders, o)
}

func (q *OrderQueue) Dequeue() (Order, error) {
	if q.IsEmpty() {
		return Order{}, errors.New("очередь заказов пуста")
	}
	first := q.orders[0]
	q.orders = q.orders[1:]
	return first, nil
}

func (q OrderQueue) IsEmpty() bool {
	return len(q.orders) == 0
}

func (q OrderQueue) TotalAmount() float64 {
	total := 0.0
	for _, o := range q.orders {
		total += o.Amount
	}
	return total
}

func main() {
	queue := &OrderQueue{}

	queue.Enqueue(Order{ID: 1, Amount: 150})
	queue.Enqueue(Order{ID: 2, Amount: 320})
	queue.Enqueue(Order{ID: 3, Amount: 75})

	fmt.Println("Сумма всех заказов в очереди:", queue.TotalAmount())

	first, err := queue.Dequeue()
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println("Обработан заказ:", first.ID)

	fmt.Println("Осталось в очереди:", queue.TotalAmount())
}
