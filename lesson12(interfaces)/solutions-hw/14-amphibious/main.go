package main

import "fmt"

type Flyer interface {
	Fly() string
}

type Swimmer interface {
	Swim() string
}

type Amphibious interface {
	Flyer
	Swimmer
}

type Duck struct {
	Name string
}

func (d Duck) Fly() string {
	return d.Name + " летит"
}

func (d Duck) Swim() string {
	return d.Name + " плывёт"
}

func main() {
	var creature Amphibious = Duck{Name: "Крякуша"}

	fmt.Println(creature.Fly())
	fmt.Println(creature.Swim())
}
