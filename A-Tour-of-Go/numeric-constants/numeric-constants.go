package main

import "fmt"

const (
	// 将1左移100位来创建一个非常大的数字，2^100
	// 即这个数的二进制是1后面跟着100个0
	Big = 1 << 100
	// 再往右移99位，即Small = 1 << 1, 或者说Small = 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x*0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	// fmt.Println("Big = %e", Big)
	fmt.Println(needFloat(Big))
	// fmt.Println(needInt(Big)) // overflows int
}