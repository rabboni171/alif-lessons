package main

import "fmt"

func main() {
	a := []int{1} 
	fmt.Printf("a = %v\tlen=%d cap=%d\n", a, len(a), cap(a)) // [1], len = 1 cap = 1

	fmt.Println("\nappend(a, 2)")
	b := append(a, 2) // [1 2] len = 2 cap = 2
	fmt.Printf("a = %v\tlen=%d cap=%d\n", a, len(a), cap(a))
	fmt.Printf("b = %v\tlen=%d cap=%d\n", b, len(b), cap(b))

	fmt.Println("\nb[0] = 0")
	b[0] = 0 // [0 2] 
	fmt.Printf("a = %v\tlen=%d cap=%d\n", a, len(a), cap(a)) // len a = 1 cap = 1
	fmt.Printf("b = %v\tlen=%d cap=%d\n", b, len(b), cap(b)) // len = 2 cap = 2

	fmt.Println("\nappend(b, 3)")
	c := append(b, 3) // [0 2 3] len = 3 cap = 4
 	fmt.Printf("b = %v\tlen=%d cap=%d\n", b, len(b), cap(b))
	fmt.Printf("c = %v\tlen=%d cap=%d\n", c, len(c), cap(c))

	fmt.Println("\nappend(c, 4)")
	d := append(c, 4) // [0 2 3 4] len = 4 cap = 4
	fmt.Printf("c = %v\tlen=%d cap=%d\n", c, len(c), cap(c))
	fmt.Printf("d = %v\tlen=%d cap=%d\n", d, len(d), cap(d))

	fmt.Println("\nc[0] = 1")
	c[0] = 1 // [1 2 3]
	fmt.Printf("b = %v\tlen=%d cap=%d\n", b, len(b), cap(b))
	fmt.Printf("c = %v\tlen=%d cap=%d\n", c, len(c), cap(c))
	fmt.Printf("d = %v\tlen=%d cap=%d\n", d, len(d), cap(d))
}

/*
формула для работы с append
append()

есть свободная capacity?
        │
   ┌────┴────┐
   │         │
  Да         Нет
  │           │
использует   создает
тот же       новый
массив       массив
*/
