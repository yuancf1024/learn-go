package main

import (
	"testing"
	"fmt"
)

func BenchmarkHello(b *testing.B) {
	... // 耗时操作
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}