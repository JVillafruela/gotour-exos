package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Error unable to compute sqrt(%v)", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0,ErrNegativeSqrt(x)
	}
	z := 1.0
	y := z
	delta := z
	for delta > 10E-5 {
		z -= (z*z - x) / (2 * z)
		delta = math.Abs(z - y)
		y = z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
