// Кавказов Алик БПО-21-01

package main

import (
	"fmt"
	"math"
)

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func exact_value(x float64) float64 {
	return math.Exp(x) - 3*x
}

func approximate_value(x, accuracy float64) (float64, int) {
	yt := exact_value(x)
	yp := 1 - 2*x
	n := 2 // n - одновременно и показатель степени и значение факториала и колво слагаемых
	for ; math.Abs(yt-yp) >= accuracy; n++ {
		yp += math.Pow(x, float64(n)) / float64(factorial(n))
	}
	return yp, n
}
func main() {
	accuracy := 0.0001
	fmt.Print("\n\nКАВКАЗОВ АЛИК БПО-21-01\n\n")
	fmt.Print("Аргумент | Точное значение | Приближенное значение | Количество слагаемыx | Ошибка\n\n")
	for value := 0.1; value <= 0.6; value += 0.1 {
		yt := exact_value(value)
		yp, n := approximate_value(value, accuracy)
		fmt.Printf("%0.1f\t |\t%f   |\t%f\t   |\t\t%d\t  | %f\n", value, yt, yp, n, yt-yp)
	}
}
