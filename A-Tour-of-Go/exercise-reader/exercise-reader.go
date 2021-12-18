package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: 给 MyReader 添加一个Read([]byte) (int, error) 方法
// func Read([]byte) (int, error) {
// 	return fmt.Sprintf("%v", )
// }

func (r MyReader) Read(b []byte) (int, error) {
	// 赋值并返回
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
