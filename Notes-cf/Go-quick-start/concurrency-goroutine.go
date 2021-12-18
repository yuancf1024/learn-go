package main

import (
	"fmt"
	"time"
)

func a() {
	time.Sleep(3 * time.Second)
	fmt.Println("it's a")
}

func b() {
	time.Sleep(2 * time.Second)
	fmt.Println("it's b")
}

func c() {
	time.Sleep(1 * time.Second)
	fmt.Println("it's c")
}

func main() {

	start := time.Now()

	fmt.Println("程序开始:")
	go a()
	go b()
	go c()
	time.Sleep(1 * time.Second)
	fmt.Println("执行完毕!")
	time.Sleep(2 * time.Second) // 为了等待3秒,以免a() b() 还没运行就结束程序了
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
