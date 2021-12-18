package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	// z := 1.0
	z := x/2
	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Printf("第%v次计算的z值: %v\n", i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(1))
	fmt.Println("***************")
	fmt.Println(Sqrt(2))
	fmt.Println("***************")
	fmt.Println(Sqrt(3))
}