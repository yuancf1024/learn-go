package main

import "fmt"

type T string

func (t *T) hello() {
	fmt.Println("hello")
}

func main() {
	var t1 T = "ABC"
	t1.hello() // hello
	// const t2 T = "ABC"
	// t2.hello()
}