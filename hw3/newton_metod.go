package main

import (
	"fmt"
	"math"
)

func main() {
	x1 := 0.0
	x2 := 0.0
	eps := 0.00001

	// Using the iteration method
	for {
		x1_new := 0.5 - math.Cos(x2+2)
		x2_new := 2.5 + math.Sin(x1+2)

		if math.Abs(x1-x1_new) < eps && math.Abs(x2-x2_new) < eps {
			break
		}
		x1 = x1_new
		x2 = x2_new
	}
	fmt.Println("The solution obtained by the iteration method: x1 =", x1, " x2 =", x2)

	// Using the Newton's method with the solution obtained by iteration method as the initial guess
	for {
		f1 := math.Sin(x1+2) - x2 - 2.5
		f2 := x1 + math.Cos(x2+2) - 0.5

		df1dx1 := math.Cos(x1 + 2)
		df1dx2 := -1.0
		df2dx1 := 1.0
		df2dx2 := -math.Sin(x2 + 2)

		det := df1dx1*df2dx2 - df1dx2*df2dx1

		x1_new := x1 - (f1*df2dx2-f2*df1dx2)/det
		x2_new := x2 - (f2*df1dx1-f1*df2dx1)/det

		if math.Abs(x1-x1_new) < eps && math.Abs(x2-x2_new) < eps {
			break
		}
		x1 = x1_new
		x2 = x2_new
	}
	fmt.Println("The solution obtained by the Newton's method: x1 =", x1, " x2 =", x2)
}
