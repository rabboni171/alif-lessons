package main

import "fmt"

//что выведет программа?

func main() {
	a1 := make([]int, 0, 10)
	a1 = append(a1, []int{1, 2, 3, 4, 5}...)
	fmt.Println(a1)
	a2 := append(a1, 6)
	fmt.Println(a2)
	a3 := append(a1, 7)
	fmt.Println(a2)
	fmt.Println(a3)

	fmt.Println(a1, a2, a3)
}
