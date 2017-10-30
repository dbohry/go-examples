package main

import "fmt"

func sqrt(x float64) float64 {
	z:= 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}

func main() {
	val := 9987665555.0
	fmt.Println(sqrt(val))
}