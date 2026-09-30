package main

import "fmt"

type Employee interface {
	Salary() float64
	Position() string
}

type Manager struct {
	Name       string
	BaseSalary float64
}

func (m Manager) Salary() float64 {
	return m.BaseSalary + m.BaseSalary*0.2
}

func (m Manager) Position() string {
	return "Manager " + m.Name
}

type Developer struct {
	Name       string
	BaseSalary float64
	Level      string
}

func (d Developer) Salary() float64 {
	salary := d.BaseSalary
	if d.Level == "senior" {
		salary += d.BaseSalary * 0.3
	}
	return salary
}

func (d Developer) Position() string {
	return "Developer " + d.Name
}

func main() {
	employees := []Employee{
		Manager{Name: "Азиз", BaseSalary: 8000},
		Developer{Name: "Диёр", BaseSalary: 7000, Level: "junior"},
		Developer{Name: "Фарход", BaseSalary: 9000, Level: "senior"},
	}

	best := employees[0]
	for _, e := range employees {
		if e.Salary() > best.Salary() {
			best = e
		}
	}

	fmt.Printf("Самая высокая зарплата: %s — %.2f\n", best.Position(), best.Salary())
}
