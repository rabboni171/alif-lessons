package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3, 4, 5} // инициализурем слайс a1
	fmt.Printf("ptr = %p, len = %d, cap =%d\n", &a1, len(a1), cap(a1))
	// len = 5 cap = 5

	a2 := append(a1, 6)
	fmt.Printf("ptr = %p, len = %d, cap =%d\n", &a2, len(a2), cap(a2))
	// len = 6 cap = 10

	a3 := append(a1, 7)
	fmt.Printf("ptr = %p, len = %d, cap =%d\n", &a3, len(a3), cap(a3))

	// len = 7 cap = 10

	// a4 := append(a3, 8, 9, 10, 11)
	// len = 11 cap = 20
	// a4[0] = -100

	fmt.Println(a1, a2, a3)
}
