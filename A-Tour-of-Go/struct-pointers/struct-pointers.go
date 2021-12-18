package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v // 指向结构体v
	p.X = 1e9 // 使用结构体指针操作结构体v中的变量X
	fmt.Println(v)
}