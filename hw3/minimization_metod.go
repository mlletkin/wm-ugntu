package main

import (
	"fmt"
	"math"
)

func main() {
	x1 := 0.0
	x2 := 0.0
	eps := 0.00001
	alpha := 0.01 // learning rate

	for {
		f1 := math.Sin(x1+2) - x2 - 2.5
		f2 := x1 + math.Cos(x2+2) - 0.5

		df1dx1 := math.Cos(x1 + 2)
		df1dx2 := -1.0
		df2dx1 := 1.0
		df2dx2 := -math.Sin(x2 + 2)

		x1_new := x1 - alpha*f1*df1dx1 - alpha*f2*df1dx2
		x2_new := x2 - alpha*f1*df2dx1 - alpha*f2*df2dx2

		if math.Abs(x1-x1_new) < eps && math.Abs(x2-x2_new) < eps {
			break
		}
		x1 = x1_new
		x2 = x2_new
	}
	fmt.Println("The solution is: x1 =", x1, " x2 =", x2)
}
