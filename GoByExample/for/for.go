package main

import "fmt"

func main() {
	
	// 最基础的方式，单个循环条件
	i := 1
	for i <= 3 {
		fmt.Println(i)
		 i = i + 1
	}

	// 经典的初始/条件/后续 for循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 不带条件的for循环将一直重复执行，直到在循环体内
	// 使用了break或者return跳出循环
	for {
		fmt.Println("loop")
		break
	}

	// 你也可以使用continue直接进入下一次循环
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}