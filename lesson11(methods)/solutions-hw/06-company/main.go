package main

import "fmt"

type Employee struct {
	Name        string
	Salary      float64
	YearsWorked int
}

func (e Employee) Bonus() float64 {
	fiveYearPeriods := e.YearsWorked / 5
	return e.Salary * 0.1 * float64(fiveYearPeriods)
}

type Company []Employee

func (c Company) TotalPayroll() float64 {
	total := 0.0
	for _, e := range c {
		total += e.Salary + e.Bonus()
	}
	return total
}

func (c Company) TopEarner() Employee {
	top := c[0]
	for _, e := range c {
		if e.Salary+e.Bonus() > top.Salary+top.Bonus() {
			top = e
		}
	}
	return top
}

func main() {
	company := Company{
		{Name: "Алишер", Salary: 3000, YearsWorked: 12},
		{Name: "Диана", Salary: 3500, YearsWorked: 4},
		{Name: "Марат", Salary: 2800, YearsWorked: 21},
	}

	fmt.Printf("Общий фонд оплаты труда: %.2f\n", company.TotalPayroll())
	top := company.TopEarner()
	fmt.Printf("Больше всех получает: %s (%.2f с учётом бонуса)\n", top.Name, top.Salary+top.Bonus())
}
