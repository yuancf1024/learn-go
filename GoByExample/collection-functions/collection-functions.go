package main

import (
	"fmt"
	"strings"
)

// Index 返回目标字符串 t 在 vs 中第一次出现位置的索引， 
// 或者在没有匹配值时返回 -1。
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include 如果目标字符串 t 存在于切片 vs 中，则返回 true。
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any 如果切片 vs 中的任意一个字符串满足条件 f，则返回 true。
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All 如果切片 vs 中的所有字符串都满足条件 f，则返回 true。
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter 返回一个新的切片，新切片由原切片 vs 中满足条件 f 的字符串构成。
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map 返回一个新的切片，新切片的字符串由原切片 vs 中的字符串经过函数 f 映射后得到。
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	// 试试各种组合函数。
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p") 
		// HasPrefix tests whether the string s begins with prefix.
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
	// 上面的例子都是用的匿名函数，当前，你也可以使用正确类型的命名函数
}