package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	
	// 在这个例子中，假如我们执行一个外部调用， 
	// 并在 2 秒后使用通道 c1 返回它的执行结果。
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	} ()

	// 这里是使用 select 实现一个超时操作。 
	// res := <- c1 等待结果，<-Time.After 等待超时（1秒钟）以后发送的值。 
	// 由于 select 默认处理第一个已准备好的接收操作，
	// 因此如果操作耗时超过了允许的 1 秒的话，将会执行超时 case。
	select {
	case res := <- c1:
		fmt.Println(res)
	case <- time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// 如果我们允许一个长一点的超时时间：3 秒， 
	// 就可以成功的从 c2 接收到值，并且打印出结果。
	select {
	case res := <- c2:
		fmt.Println(res)
	case <- time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

	cost := time.Since(start)
	fmt.Printf("cost=[%s]",cost)
}