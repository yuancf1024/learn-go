package main

import (
	"fmt"
	"time"
)

var ch chan int

func a() {
	time.Sleep(3 * time.Second)
	a := 5
	ch <- a
	fmt.Println("out of a")
}

func b() {
	time.Sleep(1 * time.Second)
}

func main() {
	ch = make(chan int, 1) // 有缓冲管道
	go a()
	go b()
	time.Sleep(5 * time.Second)
	fmt.Println("out of main")
}