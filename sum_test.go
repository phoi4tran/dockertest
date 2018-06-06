package main

import "testing"

func TestSums(t *testing.T){
	var a, b, c int
	a = add(1, 2)
	b = sub(10,5)
	c = zero(4)
	if a != 3 {
		t.Error("Expected 3, got ", a)
	}
	if b != 5 {
		t.Error("Expected 5, got ", b)
	}
	if c != 0{
		t.Error("Expected 0, got ", c)
	}
}
