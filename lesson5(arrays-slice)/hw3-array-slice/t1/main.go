package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50} // массив — размер фиксирован, всегда 5

	fmt.Println("массив:", arr)
	fmt.Println("len:", len(arr))
}
