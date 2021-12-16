package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	start := time.Now()

	// 我们将使用一个无符号整型（永远是正整数）变量来表示这个计数器。
	var ops uint64

	// WaitGroup 帮助我们等待所有协程完成它们的工作。
	var wg sync.WaitGroup

	// 我们会启动 50 个协程，并且每个协程会将计数器递增 1000 次。
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				// 使用 AddUint64 来让计数器自动增加，
				// 使用 & 语法给定 ops 的内存地址。
				atomic.AddUint64(&ops, 1)

				// 由于多个协程会互相干扰，运行时值会改变，可能会导致我们得到一个不同的数字
				// ops++
			}
			wg.Done()
		}()
	}

	// 等待，直到所有协程完成。
	wg.Wait()

	// 现在可以安全的访问 ops，因为我们知道，此时没有协程写入 ops， 此外，
	// 还可以使用 atomic.LoadUint64 之类的函数，在原子更新的同时安全地读取它们。
	fmt.Println("ops:", ops)

	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
