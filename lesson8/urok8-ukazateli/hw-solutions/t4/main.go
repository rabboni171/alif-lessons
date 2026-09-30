// Задача 4: трата энергии через указатель без ухода в минус.
package main

import "fmt"

func spendEnergy(energy *int, cost int) {
	if cost > *energy {
		*energy = 0
		return
	}
	*energy -= cost
}

func main() {
	energy := 100

	spendEnergy(&energy, 30)
	fmt.Println("Энергия:", energy) // 70

	spendEnergy(&energy, 50)
	fmt.Println("Энергия:", energy) // 20

	spendEnergy(&energy, 90)
	fmt.Println("Энергия:", energy) // 0, а не -70
}
