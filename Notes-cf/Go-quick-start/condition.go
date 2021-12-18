package main

import "fmt"

func main() {

	localStr := "case3" // 是的,还可以通过 := 这种方式直接初始化基础变量
	if localStr == "case3" {
		fmt.Printf("into true logic\n")
	} else {
		fmt.Printf("into false logic\n")
	}

	// 字典初始化
	var dic = map[string]int {
		"apple": 1,
		"watermelon": 2,
	}

	if num, ok := dic["orange"]; ok {
		fmt.Printf("orange num %d\n", num)
	}

	if num, ok := dic["watermelon"]; ok {
		fmt.Printf("watermelon num %d\n", num)
	}

	switch localStr {
	case "case1":
		fmt.Println("case1")
	case "case2":
		fmt.Println("case2")
	case "case3":
		fmt.Println("case3")
	default:
		fmt.Println("default")
	}
}