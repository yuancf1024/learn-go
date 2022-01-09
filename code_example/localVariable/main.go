package main

func foo() *int {
	v := 11
	return &v
}

func main() {
	m := foo()
	println(*m) // 11
}