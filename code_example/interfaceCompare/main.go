package main

import "fmt"

type Stu struct {
	Name string
}

type StuInt interface{}

func main() {
	var stu1, stu2 StuInt = &Stu{"Tom"}, &Stu{"Tom"}
	var stu3, stu4 StuInt = Stu{"Tom"}, Stu{"Tom"}
	fmt.Println(stu1 == stu2) // false
	fmt.Println(stu3 == stu4) // true
}