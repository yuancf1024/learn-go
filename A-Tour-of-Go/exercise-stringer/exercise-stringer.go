package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

// TODO: 给IPAddr 添加一个 "String() string" 方法

// 实现1
// func (ip IPAddr) String() string{
// 	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
// }

// 实现2
func (ip IPAddr) String() string {
	lastIndex := len(ip) - 1
	var s string
	for i, v := range ip {
		s += strconv.Itoa(int(v)) // 数字文字间相互转用strconv类
		if i != lastIndex {
			s += "."
		}
	}
	return s 
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
