package main

import "fmt"

type Handler interface {
	Handle(event any)
}

type PrintHandler struct{}

func (h PrintHandler) Handle(event any) {
	fmt.Println("PrintHandler получил событие:", event)
}

type CountingHandler struct {
	count int
}

func (h *CountingHandler) Handle(event any) {
	h.count++
	fmt.Printf("CountingHandler получил событие #%d: %v\n", h.count, event)
}

var subscribers = map[string][]Handler{}

func Subscribe(eventName string, handler Handler) {
	subscribers[eventName] = append(subscribers[eventName], handler)
}

func Publish(eventName string, event any) {
	for _, handler := range subscribers[eventName] {
		handler.Handle(event)
	}
}

func main() {
	printer := PrintHandler{}
	counter := &CountingHandler{}

	Subscribe("order_created", printer)
	Subscribe("order_created", counter)

	Publish("order_created", "заказ #1")
	Publish("order_created", "заказ #2")
	Publish("order_created", "заказ #3")
}
