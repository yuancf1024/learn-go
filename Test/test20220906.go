package main

import "fmt"

func main() {
	var s fmt.Stringer
	// var 
	s = "s"
	fmt.Println(s)
}
/*
$ go build test20220906.go 
# command-line-arguments
./test20220906.go:8:4: cannot use "s" (type string) as type fmt.Stringer in assignment:
        string does not implement fmt.Stringer (missing String method)
*/