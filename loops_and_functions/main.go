package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z, prev := 1.0, 0.0
	for {
		prev, z = z, z - (z*z - x) / (2*z)
		fmt.Println(z)
		if abs(prev - z) < 1e-15 {
			break
		}
	}
	return z
}

func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Println(Sqrt(2))
}
