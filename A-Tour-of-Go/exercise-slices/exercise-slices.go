package main

import (
	// "math"
	"golang.org/x/tour/pic"
)

/*分析：
* 外层切片的长度为dy；
* 内层切片的长度为dx；
* 内层切片中的每个元素值为(x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)
* 使用嵌套循环的方式计算颜色值
 */

func Pic(dx, dy int) [][]uint8 {
	// temp := make([][]uint8,dy)
	// for i := 1; i < dy; i++ {
	// 	temp[i] = []uint8
	// 	append(temp, [])

	a := make([][]uint8, dy) // 外层切片
	for x := range a {
		b := make([]uint8, dx) // 内层切片
		for y := range b {
			//b[y] = uint8((x+y)/2) // 给内层里的每一个元素赋值。
			//b[y] = uint8(x*y) // 给内层里的每一个元素赋值。
			//b[y] = uint8(x^y) // 给内层里的每一个元素赋值。
			//b[y] = uint8(float64(x)*math.Log(float64(y))) // 给内层里的每一个元素赋值。
			b[y] = uint8(x % (y + 1)) // 给内层里的每一个元素赋值。
		}
		a[x] = b // 给外层切片里的每一个元素赋值
	}
	return a
}

func main() {
	pic.Show(Pic)
}
