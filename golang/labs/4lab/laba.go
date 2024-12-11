package lab4

import (
	"fmt"
	"math"
)

func calculate(a, x float64) float64 {
	y := math.Pow(a, math.Pow(x, 2)-1) - math.Log10(math.Pow(x, 2)-1) + math.Cbrt(math.Pow(x, 2)-1)
	return y
}
func executeA(a float64) []float64 {
	var result []float64
	for x := 1.2; x < 2.7; x += 0.3 {
		result = append(result, calculate(a, x))
	}
	return result
}
func executeB(mass []float64, a float64) []float64 {
	var result []float64
	for _, m := range mass {
		result = append(result, calculate(a, m))
	}
	return result
}
func Executelab4() {
	var a float64 = 2.25
	var massx []float64 = []float64{1.31, 1.39, 1.44, 1.56, 1.92}
	fmt.Println(executeA(a))
	fmt.Println(executeB(massx, a))
}
