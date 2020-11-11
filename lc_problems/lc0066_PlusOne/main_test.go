package main

import "testing"

var arr = []int{1, 2, 3}

func equalIntArray(arr1, arr2 []int) bool {
	l1 := len(arr1)
	l2 := len(arr2)
	if l1 != l2 {
		return false
	}

	for i, val := range arr1 {
		if val != arr2[i] {
			return false
		}
	}
	return true
}
func TestPlusOne(t *testing.T) {

	result := plusOne(arr)
	expectedResult := []int{1, 2, 4}
	if !equalIntArray(expectedResult, result) {
		t.Error("not equal array")
	}
}

func BenchmarkPlusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusOne(arr)
	}
}
