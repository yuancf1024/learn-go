package main

import (
	"fmt"
	"testing"
)

// 通常编写一个名称以 Test 开头的函数来创建测试。
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error* 会报告测试失败的信息，然后继续运行测试。 
		// t.Fail* 会报告测试失败的信息，然后立即终止测试。
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// 单元测试可以重复，所以会经常使用 表驱动 风格编写单元测试， 
// 表中列出了输入数据，预期输出，使用循环，遍历并执行测试逻辑。
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	} {
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// t.Run 可以运行一个 “subtests” 子测试，
	// 一个子测试对应表中一行数据。 
	// 运行 go test -v 时，他们会分开显示。

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func (t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}