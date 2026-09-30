package main

import "fmt"

type Character struct {
	Name   string
	HP     int
	Attack int
}

func (c *Character) IsAlive() bool {
	return c.HP > 0
}

func (c *Character) TakeDamage(dmg int) {
	c.HP -= dmg
	if c.HP < 0 {
		c.HP = 0
	}
}

func (c *Character) AttackTarget(target *Character) {
	target.TakeDamage(c.Attack)
}

func main() {
	hero := &Character{Name: "Герой", HP: 30, Attack: 5}
	dragon := &Character{Name: "Дракон", HP: 40, Attack: 4}

	attacker, defender := hero, dragon

	for hero.IsAlive() && dragon.IsAlive() {
		attacker.AttackTarget(defender)
		fmt.Printf("%s атакует %s, у %s осталось %d HP\n", attacker.Name, defender.Name, defender.Name, defender.HP)
		attacker, defender = defender, attacker
	}

	if hero.IsAlive() {
		fmt.Println("Победил:", hero.Name)
	} else {
		fmt.Println("Победил:", dragon.Name)
	}
}
