package main

import (
	"fmt"
	"math"
)

func main() {
	x1 := 0.0
	x2 := 0.0
	eps := 0.00001

	for {
		x1_new := 0.5 - math.Cos(x2+2)
		x2_new := 2.5 + math.Sin(x1+2)

		if math.Abs(x1-x1_new) < eps && math.Abs(x2-x2_new) < eps {
			break
		}
		x1 = x1_new
		x2 = x2_new
	}
	fmt.Println("x1 =", x1, " x2 =", x2)
	fmt.Println("Кавказов Алик БПО-21-01")
}
