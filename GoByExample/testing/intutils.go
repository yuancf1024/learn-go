package main

// 我们要测试下面这个简单的函数——返回最小值。 
// 一般地，需要被测试的代码应该在类似于 intutils.go 的文件下， 
// 其对应的测试文件应该被命名为 intutils_test.go。

func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
