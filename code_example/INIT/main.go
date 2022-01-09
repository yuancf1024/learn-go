package main

import "fmt"

func init() {
	fmt.Println("init1:", a)
}

func init() {
	fmt.Println("init2:", a)
}

var a = 10

const b = 100

func main() {
	fmt.Println("main", a)
}