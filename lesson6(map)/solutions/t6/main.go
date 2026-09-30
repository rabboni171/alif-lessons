// Задача 6: лайки видео через map[string]int и оператор ++ (без среза).
package main

import "fmt"

func main() {
	likes := make(map[string]int)

	likes["Урок про map"]++
	likes["Урок про map"]++
	likes["Урок про map"]++

	likes["Урок про срезы"]++
	likes["Урок про срезы"]++

	fmt.Println(likes)
}
