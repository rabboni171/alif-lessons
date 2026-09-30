package main

import "fmt"

type Box struct {
	Value any
}

func main() {
	boxes := []Box{
		{Value: 42},
		{Value: "строка"},
		{Value: 3.14},
		{Value: []int{1, 2, 3}},
		{Value: Box{Value: 1}},
	}

	for _, b := range boxes {
		fmt.Printf("содержимое: %v, тип: %T\n", b.Value, b.Value)
	}
}
