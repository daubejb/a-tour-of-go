package main

import (
	"fmt"
)

// Sqrt - finds the sqrt of x (float64)
func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("try %v: %v\n", i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
