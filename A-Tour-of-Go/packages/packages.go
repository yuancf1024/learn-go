package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand.Seed(1) 设置随机种子,默认种子里面的数是1
	// 输出我最喜欢的数字
	fmt.Println("My favorite number is", rand.Intn(10))

	rand.Seed(time.Now().UnixNano())
	fmt.Println(randomString(10)) // 返回一个随机的10字母字符串
}

func randomString(l int) string {
	bytes := make([] byte, l)
	for i := 0; i<l; i++ {
		bytes[i] = byte(randInt(65, 90)) // 从大写字母中获取
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}