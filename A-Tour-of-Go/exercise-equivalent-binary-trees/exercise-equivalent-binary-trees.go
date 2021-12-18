package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// 发送value，结束后关闭channel
// Walk 步进 tree t 将所有的值从tree 发送到channel ch
func Walk(t *tree.Tree, ch chan int) {
	snedValue(t, ch)
	close(ch)
}

// 递归向channel传值
func snedValue(t *tree.Tree, ch chan int) {
	if t != nil {
		snedValue(t.Left, ch)
		ch <- t.Value
		snedValue(t.Right, ch)
	}
}

// 使用写好的walk函数来确定两个tree对象 是否一样 原理还是判断value值
// Same 检测树t1 和 t2是否含有相同的值
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		if i != <- ch2 {
			return false
		}
	}
	return true
}

func main() {
	// 打印tree.New(1)的值
	var ch = make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v)
	}

	// 比较两个tree的value值是否相等
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
