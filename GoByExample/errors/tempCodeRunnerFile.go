package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New 使用给定的错误信息构造一个基本的 error 值。
		return -1, errors.New("can't work with 42")
	}
	
	return arg + 3, nil // 返回错误值为 nil 代表没有错误。
}

type argError struct {
	arg int
	prob string
}

// 你还可以通过实现 Error() 方法来自定义 error 类型。 
// 这里使用自定义错误类型来表示上面例子中的参数错误。
func (e *argError) Error() string { // * 解引用指针，从对应地址获取值
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// 在这个例子中，我们使用 &argError 语法来建立一个新的结构体， 
// 并提供了 arg 和 prob 两个字段的值。
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// 下面的两个循环测试了每一个会返回错误的函数。 
	// 注意，在 if 的同一行进行错误检查，是 Go 代码中的一种常见用法。
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
		fmt.Println("***Test***") // test
		fmt.Println(ok) // test
		fmt.Println(e.(*argError))
		fmt.Println(e)

	}
}