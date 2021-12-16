package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	// 打点器和定时器的机制有点相似：使用一个通道来发送数据。
	// 这里我们使用通道内建的 select，等待每 500ms 到达一次的值。
	fmt.Println("Test1") // Test
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// 打点器可以和定时器一样被停止。
	// 打点器一旦停止，将不能再从它的通道中接收到值。
	// 我们将在运行 1600ms 后停止这个打点器。
	fmt.Println("Test2") // Test
	time.Sleep(1600 * time.Millisecond)
	fmt.Println("Test3") // Test
	ticker.Stop()
	fmt.Println("Test4") // Test
	done <- true
	fmt.Println("Ticker stopped")

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
