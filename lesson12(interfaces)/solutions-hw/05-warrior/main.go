package main

import "fmt"

type Warrior interface {
	Attack() int
	Description() string
}

type Knight struct {
	Name     string
	Strength int
}

func (k Knight) Attack() int {
	return k.Strength * 2
}

func (k Knight) Description() string {
	return "Knight " + k.Name
}

type Archer struct {
	Name      string
	Precision int
}

func (a Archer) Attack() int {
	return a.Precision + 10
}

func (a Archer) Description() string {
	return "Archer " + a.Name
}

type Mage struct {
	Name      string
	ManaPower int
}

func (m Mage) Attack() int {
	return m.ManaPower * 3
}

func (m Mage) Description() string {
	return "Mage " + m.Name
}

func main() {
	squad := []Warrior{
		Knight{Name: "Артур", Strength: 15},
		Archer{Name: "Робин", Precision: 20},
		Mage{Name: "Гэндальф", ManaPower: 8},
	}

	total := 0
	for _, w := range squad {
		fmt.Println(w.Description())
		total += w.Attack()
	}

	fmt.Println("Суммарная сила атаки отряда:", total)
}
