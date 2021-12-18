package main

import "fmt"

/*
斐波那契数，亦称之为斐波那契数列（意大利语： Successione di Fibonacci)，
又称黄金分割数列、费波那西数列、费波拿契数、费氏数列，指的是这样一个数列：
0、1、1、2、3、5、8、13、21、……在数学上，斐波纳契数列以如下被以递归的方法定义：
F0=0，F1=1，Fn=Fn-1+Fn-2（n>=2，n∈N*），用文字来说，就是斐波那契数列列由
 0 和 1 开始，之后的斐波那契数列系数就由之前的两数相加。
*/

// // 返回一个“返回int的函数”
// func fibonacci() func(int) int {
// 	var a [10]int
// 	a[0] = 0
// 	a[1] = 1
// 	return func(x int) int {
// 		a[x] = a[x] + a[x-1]
// 		return a[x]
// 	}
// }

func fibonacci() func() int {
	back1, back2 := 0, 1 // 预先定义好前两个值

	return func() int {
		// 记录（back1）的值
		temp := back1

		// 重新赋值（这个就是核心代码）
		back1, back2 = back2, (back1 + back2)

		// 返回temp
		return temp
	}
}

func main() {
	f := fibonacci() // 返回一个闭包函数
	for i := 0; i < 10; i++ { // 检测下前10个值
		fmt.Println(f())
	}
}
