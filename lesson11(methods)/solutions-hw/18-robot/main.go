package main

import "fmt"

type Robot struct {
	x, y int
}

func (r *Robot) MoveUp() {
	r.y++
}

func (r *Robot) MoveDown() {
	r.y--
}

func (r *Robot) MoveLeft() {
	r.x--
}

func (r *Robot) MoveRight() {
	r.x++
}

func (r Robot) Position() (int, int) {
	return r.x, r.y
}

func (r Robot) String() string {
	return fmt.Sprintf("Robot @ (%d, %d)", r.x, r.y)
}

func main() {
	robot := &Robot{}

	robot.MoveUp()
	fmt.Println(robot)

	robot.MoveRight()
	fmt.Println(robot)

	robot.MoveRight()
	fmt.Println(robot)

	robot.MoveDown()
	fmt.Println(robot)

	robot.MoveLeft()
	fmt.Println(robot)
}
