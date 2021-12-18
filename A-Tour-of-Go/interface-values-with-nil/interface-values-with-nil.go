package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型T实现了接口I，但我们无需显式声明此事。
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
