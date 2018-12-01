package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1
	y := z
	delta := z
	for delta > 10E-5 {
		z -= (z*z - x) / (2 * z)
		delta = math.Abs(z - y)
		y = z
		fmt.Println(delta)
	}
	return z
}

func main() {
	srm := Sqrt(2)
	srl := math.Sqrt(2)
	fmt.Println(srm, srl, srm-srl)
}
