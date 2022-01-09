package main

import "testing"

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}

func TestMul(t *testing.T) {
	if ans := Mul(1, 2); ans != 2 {
		t.Errorf("1 + 2 expected be 2, but %d got", ans)
	}

	if ans := Mul(-10, -20); ans != 200 {
		t.Errorf("-10 + -20 expected be 200, but %d got", ans)
	}
}