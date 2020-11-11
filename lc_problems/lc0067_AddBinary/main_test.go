package main

import "testing"

var a = "11"
var b = "1"
var r = "100"

func TestAddBinary(t *testing.T) {
	if addBinary(a, b) != r {
		t.Error("wrong result")
	}
}

func BenchmarkAddBinary(bt *testing.B) {
	for i := 0; i < bt.N; i++ {
		addBinary(a, b)
	}
}
