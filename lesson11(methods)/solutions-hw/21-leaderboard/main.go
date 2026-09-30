package main

import "fmt"

type Player struct {
	Name  string
	Score int
}

func (p *Player) AddPoints(p2 int) {
	p.Score += p2
}

func TopPlayer(players []Player) Player {
	top := players[0]
	for _, p := range players {
		if p.Score > top.Score {
			top = p
		}
	}
	return top
}

func main() {
	players := []Player{
		{Name: "Алишер"},
		{Name: "Диана"},
		{Name: "Марат"},
		{Name: "Санжар"},
	}

	players[1].AddPoints(50)
	players[0].AddPoints(30)
	players[3].AddPoints(80)
	players[2].AddPoints(20)
	players[3].AddPoints(10)

	winner := TopPlayer(players)
	fmt.Printf("Победитель: %s (%d очков)\n", winner.Name, winner.Score)
}
