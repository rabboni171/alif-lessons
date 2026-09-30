package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

func (s Student) Average() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	total := 0
	for _, g := range s.Grades {
		total += g
	}
	return float64(total) / float64(len(s.Grades))
}

func (s Student) String() string {
	return fmt.Sprintf("%s (средний балл: %.2f)", s.Name, s.Average())
}

func main() {
	students := []Student{
		{Name: "Али", Grades: []int{4, 5, 5, 4}},
		{Name: "Бек", Grades: []int{3, 4, 3, 5}},
		{Name: "Дана", Grades: []int{5, 5, 5, 5}},
		{Name: "Ержан", Grades: []int{4, 3, 4, 4}},
		{Name: "Санжар", Grades: []int{5, 4, 5, 4}},
	}

	headOfClass := students[0]
	for _, s := range students {
		if s.Average() > headOfClass.Average() {
			headOfClass = s
		}
	}

	fmt.Println("Староста класса:", headOfClass)
}
