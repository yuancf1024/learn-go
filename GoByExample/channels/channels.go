package main

import "fmt"

func main() {
	// 使用 make(chan val-type) 创建一个新的通道。 
	// 通道类型就是他们需要传递值的类型。
	messages := make(chan string)
	fmt.Println("Test1") // Test

	go func() {
		// 使用 channel <- 语法 发送 一个新的值到通道中。 
		// 这里我们在一个新的协程中发送 "ping" 到上面创建的 messages 通道中。
		messages <- "ping"
	}()

	// 使用 <-channel 语法从通道中 接收 一个值。 
	// 这里我们会收到在上面发送的 "ping" 消息并将其打印出来。
	fmt.Println("Test2") // Test
	msg := <- messages
	fmt.Println("Test3") // Test
	fmt.Println(msg)
	fmt.Println("Test4") // Test
}