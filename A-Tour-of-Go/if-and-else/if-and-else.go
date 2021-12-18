package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用v了
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 1, 3),
		pow(3, 3, 20),
	)
}