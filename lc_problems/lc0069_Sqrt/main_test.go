package main

import "testing"

func TestMySqrt(t *testing.T) {
	x := 8
	if mySqrt(x) != 2 {
		t.Error("expect 2")
	}
}
