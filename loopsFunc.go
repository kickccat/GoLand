package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	i := 0
	for temp := z - (z*z-x)/(2*z); math.Abs(temp-z) > 1.0e-10; {
		z = temp
		temp = z - (z*z-x)/(2*z)
		i++
	}

	fmt.Printf("i = %g\n", i)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
