package main

import "fmt"

type TrafficLight struct {
	color string
}

func NewTrafficLight() *TrafficLight {
	return &TrafficLight{color: "красный"}
}

func (t *TrafficLight) Next() {
	switch t.color {
	case "красный":
		t.color = "жёлтый"
	case "жёлтый":
		t.color = "зелёный"
	case "зелёный":
		t.color = "красный"
	}
}

func (t TrafficLight) Current() string {
	return t.color
}

func (t TrafficLight) String() string {
	return "Светофор: " + t.color
}

func main() {
	light := NewTrafficLight()

	for i := 0; i < 7; i++ {
		fmt.Println(light)
		light.Next()
	}
}
