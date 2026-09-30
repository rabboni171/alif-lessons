package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Info() string {
	return fmt.Sprintf("%s, %d лет", p.Name, p.Age)
}

type Student struct {
	Person
	University string
}

func (s Student) Info() string {
	return s.Person.Info() + ", учится в " + s.University
}

func main() {
	s := Student{
		Person:     Person{Name: "Аида", Age: 20},
		University: "МГУ",
	}

	fmt.Println(s.Info())
}
