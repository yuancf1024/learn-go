package main

import (
	"fmt"
)

func main() {
	v := [...]int{1:2, 3:4}
	fmt.Println(len(v))
	fmt.Printf("v= %v ", v)
}

// v := [...]int{1:2, 3:4} 这是声明的数组，而不是切片
// 数组的长度由声明的数据个数自动计算，其中指定了1和3的位置为2和4
// 其他0和2的位置使用零值，所以长度就是4