package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	// 我们将要在协程中运行这个函数。 
	// done 通道将被用于通知其他协程这个函数已经完成工作。
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值来通知我们已经完工啦。
	done <- true // 从一个协程将值发送到通道
}

func main() {

	// 运行一个 worker 协程，并给予用于通知的通道。
	done := make(chan bool, 1)
	go worker(done)
	fmt.Println("Test1")

	// 程序将一直阻塞，直至收到 worker 使用通道发送的通知。
	<-done // 从通道接受
	// 如果你把 <- done 这行代码从程序中移除， 程序甚至可能在 worker 开始运行前就结束了。
	fmt.Println("Test2")
}