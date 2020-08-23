package main

import (
	"fmt"
	"math"
)

func main() {
	PI := 3.14159
	var R, A float64
	fmt.Scanf("%f", &R)
	A = PI * math.Pow(R, 2)
	fmt.Printf("A=%.4f\n", A)
}
