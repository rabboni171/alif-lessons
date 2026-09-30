package main

import "fmt"

type Instrument interface {
	Play() string
}

type Guitar struct {
	Name string
}

func (g Guitar) Play() string {
	return "Guitar " + g.Name + ": дзынь-дзынь"
}

type Piano struct {
	Name string
}

func (p Piano) Play() string {
	return "Piano " + p.Name + ": тим-тим-тим"
}

type Drum struct {
	Name string
}

func (d Drum) Play() string {
	return "Drum " + d.Name + ": бум-бум"
}

func main() {
	orchestra := []Instrument{
		Guitar{Name: "Fender"},
		Piano{Name: "Yamaha"},
		Drum{Name: "Pearl"},
		Guitar{Name: "Ibanez"},
	}

	for _, instrument := range orchestra {
		fmt.Println(instrument.Play())
	}
}
