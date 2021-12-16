package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	// 首先，我们将看一个基本的速率限制。
	// 假设我们想限制对收到请求的处理，我们可以通过一个通道处理这些请求
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter 通道每 200ms 接收一个值。 这是我们任务速率限制的调度器。
	limiter := time.Tick(200 * time.Millisecond)

	// 通过在每次请求前阻塞 limiter 通道的一个接收，
	// 可以将频率限制为，每 200ms 执行一次请求。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们可能希望在速率限制方案中允许短暂的并发请求，
	// 并同时保留总体速率限制。 我们可以通过缓冲通道来完成此任务。
	// burstyLimiter 通道允许最多 3 个爆发（bursts）事件。
	burstLimiter := make(chan time.Time, 3)

	// 填充通道，表示允许的爆发（bursts）。
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	// 每 200ms 我们将尝试添加一个新的值到 burstyLimiter中， 直到达到 3 个的限制。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstLimiter <- t
		}
	}()

	// 现在，模拟另外 5 个传入请求。 受益于 burstyLimiter 的爆发（bursts）能力，
	// 前 3 个请求可以快速完成。
	burstRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)
	for req := range burstRequests {
		<-burstLimiter
		fmt.Println("request", req, time.Now())
	}

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
