package main

import (
	"fmt"
	"runtime"
)

func query() int {
	ch := make(chan int)
	for i := 0; i < 1000; i++ {
		go func() {ch <- 0}()
	}
	return <- ch
}

func main() {
	for i := 0; i < 4; i++ {
		query()
		fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
	}
}
// goroutines: 1001
// goroutines: 2000
// goroutines: 2999
// goroutines: 3998