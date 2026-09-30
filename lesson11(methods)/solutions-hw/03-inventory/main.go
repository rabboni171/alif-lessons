package main

import (
	"fmt"
)

type Inventory struct {
	items map[string]int
}

func (inv *Inventory) Add(item string, qty int) {
	if inv.items == nil {
		inv.items = make(map[string]int)
	}
	inv.items[item] += qty
}

func (inv *Inventory) Remove(item string, qty int) error {
	have, ok := inv.items[item]
	if !ok {
		return fmt.Errorf("товара %q нет на складе", item)
	}
	if have < qty {
		return fmt.Errorf("не хватает %q: есть %d, а нужно %d", item, have, qty)
	}
	inv.items[item] -= qty
	return nil
}

func (inv Inventory) Count(item string) int {
	return inv.items[item]
}

func (inv Inventory) Total() int {
	total := 0
	for _, qty := range inv.items {
		total += qty
	}
	return total
}

func main() {
	inv := &Inventory{}

	inv.Add("яблоко", 50)
	inv.Add("банан", 30)
	inv.Add("груша", 20)

	fmt.Println("Яблок на складе:", inv.Count("яблоко"))
	fmt.Println("Всего товаров:", inv.Total())

	if err := inv.Remove("банан", 10); err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println("Бананов после списания:", inv.Count("банан"))

	if err := inv.Remove("апельсин", 5); err != nil {
		fmt.Println("Ошибка:", err)
	}

	if err := inv.Remove("груша", 100); err != nil {
		fmt.Println("Ошибка:", err)
	}
}
