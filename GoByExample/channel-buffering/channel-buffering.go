package main

import "fmt"

func main() {
	// 这里我们 make 了一个字符串通道，最多允许缓存 2 个值。
	messages := make(chan string, 2)

	// 由于此通道是有缓冲的， 因此我们可以将这些值发送到通道中，而无需并发的接收。
	messages <- "buffered"
	messages <- "channel"

	// 然后我们可以正常接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}