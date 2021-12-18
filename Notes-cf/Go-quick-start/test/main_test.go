package main

import "testing"

func TestRect(t *testing.T) {
	var s Shape = &Rect{height: 10, weight: 8}
	if s.Area() != 80 {
		t.Errorf("area %f\n", s.Area())
	}

	if s.Perimeter() != 30 {
		t.Errorf("perimeter %f\n", s.Perimeter())
	}
}