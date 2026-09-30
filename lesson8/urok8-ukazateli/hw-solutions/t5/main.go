// Задача 5: double по значению (не меняет оригинал) vs doublePtr через указатель.
package main

import "fmt"

func double(n int) int {
	return n * 2
}

func doublePtr(n *int) {
	*n = *n * 2
}

func main() {
	x := 5
	result := double(x)
	fmt.Println("double: x =", x, ", result =", result) // x = 5, result = 10

	y := 5
	doublePtr(&y)
	fmt.Println("doublePtr: y =", y) // y = 10
}
