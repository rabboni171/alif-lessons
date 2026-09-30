package main

import "fmt"

func main() {
	parentSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19} // len = 20
	// cap = 20
	fmt.Printf("Parent Slice: %v\n", parentSlice)
	//fmt.Printf("Parent Slice's Address: %p\n\n", parentSlice)
	fmt.Println("len = ", len(parentSlice), "cap = ", cap(parentSlice))
	// len = 20 cap = 20

	parentSlice = append(parentSlice, 20, 21, 22, 23)
	// len = 24 cap = 40
	fmt.Println("parent slice after appending:", parentSlice)
	fmt.Println(len(parentSlice))
	fmt.Println(cap(parentSlice))

	parentSlice1 := append(parentSlice, 24, 25, 26, 27)
	fmt.Println("parent slice1 after appending:", parentSlice1)
	fmt.Println(len(parentSlice1))
	fmt.Println(cap(parentSlice1))

	parentSlice1[0] = -1
	parentSlice1[1] = -2
	parentSlice1[2] = -3

	fmt.Println("parent slice:", parentSlice)
	fmt.Println("parent slice1:", parentSlice1)

	// // новый слайс sliceA
	// sliceA := parentSlice[0:5]
	// fmt.Printf("sliceA: %v\n", sliceA) // [0 1 2 3 4]
	// fmt.Printf("sliceA's Address: %p\n", sliceA)
	// fmt.Printf("sliceA's Length %d\n", len(sliceA)) // len = 5
	// fmt.Printf("sliceA's Capacity %d\n\n", cap(sliceA)) // cap = 20

	// // новый слайс sliceB
	// sliceB := parentSlice[12:15] // [12 13 14]
	// fmt.Printf("sliceB: %v\n", sliceB)
	// fmt.Printf("sliceB's Address: %p\n", sliceB)
	// fmt.Printf("sliceB's Length %d\n", len(sliceB)) // len = 3
	// fmt.Printf("sliceB's Capacity %d", cap(sliceB)) // cap = 8
}
