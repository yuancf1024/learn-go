package main

const cl = 100

var bl = 123

func main()  {
    println(&bl,bl)
    println(&cl,cl) // bug
}
// # command-line-arguments
// Go-aceld/main/test-variable.go:9:13: cannot take the address of cl