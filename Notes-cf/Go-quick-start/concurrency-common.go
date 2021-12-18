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
	a()
	b()
	c()
	time.Sleep(1 * time.Second)
	fmt.Println("执行完毕!")
	cost := time.Since(start)
	fmt.Printf("cost=[%s]", cost)
}
