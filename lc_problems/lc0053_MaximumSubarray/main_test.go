package main

import "testing"

var arr = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

func TestMaxSubArray(t *testing.T) {
	result := maxSubArray(arr)
	if result != 6 {
		t.Error("expect 6")
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSubArray(arr)
	}
}
