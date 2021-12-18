package main

import "fmt"

var globalStr string
var globalInt int

func main() {
	var localStr string
	var localInt int
	localStr = "first local"
	localInt = 2021
	globalInt = 1024
	globalStr = "first global"
	fmt.Printf("globalStr is %s\n", globalStr) // globalStr is first global
	fmt.Printf("globalInt is %d\n", globalInt) // globalInt is 1024
	fmt.Printf("localStr is %s\n", localStr) // localStr is first local
	fmt.Printf("localInt is %d\n", localInt) // localInt is 2021

}