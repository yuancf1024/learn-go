package main

func f(s []int) {
	_ = s[2]
	_ = s[1]
	_ = s[0]
}
func main() {}

// Output:
// PS D:\gocf\src\github.com\yuancf1024\learn-go\code_example\BCE> go build -gcflags="-d=ssa/check_bce/debug=1" main.go
// # command-line-arguments
// .\main.go:4:7: Found IsInBounds

// import "fmt"

// func StringSliceEqualBCE(a, b []string) bool {
//     // 两个切片长度不同时
//     if len(a) != len(b) {
//         return false
//     }

//     // a b不是都为空时
//     if (a == nil ) != (b == nil) {
//         return false
//     }

//     // 遍历比较切片中的每一个元素
//     b = b[:len(a)]
//     for i, v := range a {
//         if v != b[i] {
//             return false
//         }
//     }
//     return true
// }

// func main() {
// 	a := []string{"a", "b", "c", "d", "e"}
// 	b := []string{"a", "b", "c", "d", "e"}
// 	fmt.Println(StringSliceEqualBCE(a, b))
// }
