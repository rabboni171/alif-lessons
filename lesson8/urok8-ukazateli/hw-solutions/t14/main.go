// Задача 14: круги гонки - два указателя меняются в одной функции.
package main

import "fmt"

func advanceLap(currentLap *int, totalLaps int, finished *bool) {
	if *finished {
		return
	}

	*currentLap++
	if *currentLap >= totalLaps {
		*currentLap = totalLaps
		*finished = true
	}
}

func main() {
	lap := 0
	finished := false
	totalLaps := 3

	for i := 0; i < 5; i++ {
		advanceLap(&lap, totalLaps, &finished)
		fmt.Println("Круг:", lap, "| финиш:", finished)
	}
}
