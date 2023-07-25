package main

import "fmt"

func main() {
  
	list := new([]int)
  
	*list = append(*list, 1) // list类型为指针
  
	fmt.Println(list)

	list1 := make([]int, 0)
  
	list1 = append(list1, 1) // list1类型为切片
  
	fmt.Println(list1)
}