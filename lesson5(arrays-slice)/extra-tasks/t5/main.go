package main

import "fmt"

func main() {
	var x []int // создаем пустой интовый слайс
	//fmt.Println("x=", x, "len x:", len(x), "cap x:", cap(x))
	x = append(x, 0) // [0] len = 1 cap = 1
	x = append(x, 1) // [0 1] len = 2 cap = 2
	x = append(x, 2) // [0 1 2] len = 3 cap = 4

	y := append(x, 3) // [0 1 2 4] len = 4 cap = 4
	z := append(x, 4) // [0 1 2 4] len = 4 cap = 4

	fmt.Println(y)
	fmt.Println(z)
}
