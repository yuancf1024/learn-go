package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Printf("current i %d\n", i)
	}

	j := 0
	for {
		if j == 5 {
			break
		}
		fmt.Printf("current j %d\n", j)
		j++
	}

	var strAry = []string{"aa", "bb", "cc", "dd", "ee"} // 是的,不指定初始个数也ok

	// 切片初始化
	var sliceAry = make([]string, 0)
	sliceAry = strAry[1:3]
	for i, str := range sliceAry {
		fmt.Printf("slice i %d, str %s\n", i, str)
	}

	// 字典初始化
	var dic = map[string]int {
		"apple": 1,
		"google": 2,
		"facebook": 3,
		"amazon": 4,
	}

	for k, v := range dic {
		fmt.Printf("key %s, value %d\n", k, v)
	}
}