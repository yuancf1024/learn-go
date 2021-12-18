package main

import "fmt"

func main() {

	// 数组初始化
	var strAry = [10]string{"aa", "bb", "cc", "dd", "ee"}
	// 切片初始化
	var sliceAry = make([]string, 0)
	sliceAry = strAry[1:3]
	// 字典初始化
	var dic = map[string]int {
		"apple": 1,
		"watermelon": 2,
	}

	fmt.Printf("strAry %+v\n", strAry)
	fmt.Printf("sliceAry %+v\n", sliceAry)
	fmt.Printf("dic %+v\n", dic)
}