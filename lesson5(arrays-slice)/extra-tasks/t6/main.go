package main

import "fmt"

func slicesSharedMemAppend() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2] // [a b] len = 2 cap = 4
	fmt.Println("cap x:", cap(x))
	fmt.Println("len y:", len(y), "cap y:", cap(y))
	y = append(y, "z") //[a b z]
	fmt.Println("x:", x)
	fmt.Println("y after appending:", y)
}

func main() {
	slicesSharedMemAppend()
}
