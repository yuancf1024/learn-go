package main

import "fmt"

func main() {
	var p *int = nil
	var i interface{} = p
	fmt.Println(i == p)
	fmt.Println(p == nil)
	fmt.Println(i == nil)
}