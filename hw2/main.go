package main

import (
	"fmt"
	"math"
)

func y(x float64) float64 {
	return math.Pow(math.Abs(math.Sin(math.Pow(x, 2)+1)), 1.0/3.0) - x
}

func main() {
	a, b := 0.0, 1.0
	e := 0.000001
	c := 0.0
	for ; math.Abs(b-a) > e; c = (a + b) / 2 {
		if y(a)*y(c) < 0 {
			b = c
		} else {
			a = c
		}
	}
	fmt.Println("Метод половинного деления Кавказов Алик")
	fmt.Println(c)
}

// Кавказов Алик БПО-21-01. Метод половинного деления
