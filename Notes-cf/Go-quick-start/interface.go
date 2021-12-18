package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

// type rect
type Rect struct {
	height float64
	weight float64
}

func (p *Rect) Area() float64 {
	return p.height * p.weight
}

func (p *Rect) Perimeter() float64 {
	return 2 * (p.height + p.weight)
}

func main() {
	var s Shape = &Rect{height: 10, weight: 8}
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}