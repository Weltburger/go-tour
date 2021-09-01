package main

import "fmt"

type ErrNegativeSqrt  float64

func (e ErrNegativeSqrt ) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0.0, ErrNegativeSqrt(x)
	}

	z, prev := 1.0, 0.0
	for {
		prev, z = z, z - (z*z - x) / (2*z)
		fmt.Println(z)
		if abs(prev - z) < 1e-15 {
			break
		}
	}
	return z, nil
}

func abs(n float64) float64 {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
