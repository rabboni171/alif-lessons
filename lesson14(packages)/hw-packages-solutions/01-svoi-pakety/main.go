package main

import (
	"fmt"

	"svoipakety/geometry"
	"svoipakety/strutil"
	"svoipakety/validator"
)

func main() {
	fmt.Println("=== Задача 1: geometry ===")
	type rect struct {
		width, height float64
	}
	rects := []rect{
		{width: 3, height: 4},
		{width: 5.5, height: 2},
		{width: 10, height: 10},
	}
	for _, r := range rects {
		area := geometry.RectangleArea(r.width, r.height)
		perimeter := geometry.RectanglePerimeter(r.width, r.height)
		fmt.Printf("прямоугольник %vx%v: площадь=%v, периметр=%v\n", r.width, r.height, area, perimeter)
	}

	fmt.Println("\n=== Задача 2: strutil ===")
	phrases := []string{"топот", "шалаш", "привет", "А роза упала на лапу Азора"}
	for _, p := range phrases {
		if strutil.IsPalindrome(p) {
			fmt.Printf("%q - палиндром\n", p)
		} else {
			fmt.Printf("%q - не палиндром\n", p)
		}
	}

	fmt.Println("\n=== Задача 3: validator ===")
	type user struct {
		email string
		age   int
	}
	users := []user{
		{email: "ali@mail.tj", age: 21},
		{email: "farhod@edu.tj", age: 16},
		{email: "no-at-sign.com", age: 30},
		{email: "student@mail.", age: 19},
	}
	for _, u := range users {
		emailOK := validator.IsValidEmail(u.email)
		adultOK := validator.IsAdult(u.age)
		if emailOK && adultOK {
			fmt.Printf("%s (возраст %d) - обе проверки пройдены\n", u.email, u.age)
		} else {
			fmt.Printf("%s (возраст %d) - не прошёл: email=%t, adult=%t\n", u.email, u.age, emailOK, adultOK)
		}
	}
}
