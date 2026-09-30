package main

import "fmt"

type Animal struct {
	Name  string
	Sound string
}

func (a Animal) MakeSound() string {
	return fmt.Sprintf("%s говорит %s", a.Name, a.Sound)
}

type Dog struct {
	Animal
}

func (d Dog) MakeSound() string {
	return d.Animal.MakeSound() + " (виляет хвостом)"
}

type Cat struct {
	Animal
}

func (c Cat) MakeSound() string {
	return c.Animal.MakeSound() + " (жмурится)"
}

type Sounder interface {
	MakeSound() string
}

func main() {
	animals := []Sounder{
		Dog{Animal{Name: "Bobik", Sound: "Гав"}},
		Cat{Animal{Name: "Murka", Sound: "Мяу"}},
		Dog{Animal{Name: "Rex", Sound: "Гав"}},
	}

	for _, a := range animals {
		fmt.Println(a.MakeSound())
	}
}
