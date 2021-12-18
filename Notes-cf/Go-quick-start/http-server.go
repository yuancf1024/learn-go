package main

import (
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello")) // 以字符串"hello"作为返回包
}

func main() {
	http.HandleFunc("/say_hello", sayHello)
	err := http.ListenAndServe(":8080", nil) // 开启一个http服务
	if err != nil {
		log.Print("ListenAndServe:", err)
		return
	}
}