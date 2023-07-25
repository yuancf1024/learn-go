package main

import (
    "fmt"
)

type student struct {
    Name string
    Age  int
}

func main() {
    //定义map
    m := make(map[string]*student)

    //定义student数组
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }

    //将数组依次添加到map中
    for _, stu := range stus {
        m[stu.Name] = &stu
    } // 这种遍历方式有问题，stu指向最后一个遍历到的结构体

	// 正确的遍历map知识
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}

    //打印map
    for k,v := range m {
        fmt.Println(k ,"=>", v.Name)
    }
}