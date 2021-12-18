package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// 转换byte，前进13位/后退13位
func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M' || 'a' <= b && b <= 'm':
		b += 13
	case 'M' < b && b <= 'Z' || 'm' < b && b <= 'z':
		b -= 13
	}
	return b
}

// rot13的更精简写法
// func rot13(b byte) byte {
// 	if ('A' <= b && b <= 'M') || ('a' <= b && b <= 'm') {
// 		b += 13
// 	} else if ('N' <= b && b <= 'Z') || ('n' <= b && b <= 'z') {
// 		b -= 13
// 	}
// 	return b
// }

// rot13再精简的写法
// func rot13(b byte) byte {
// 	switch {
// 	case 'A' <= b && b <= 'Z':
// 		b = (b-'A'+13)%26 + 'A'
// 	case 'a' <= b && b <= 'z':
// 		b = (b-'a'+13)%26 + 'a'
// 	}
// }

// 重写Read方法
func (mr rot13Reader) Read(b []byte) (int, error) {
	n, err := mr.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
