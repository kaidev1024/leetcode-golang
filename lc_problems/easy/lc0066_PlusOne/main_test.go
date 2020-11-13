package main

import (
	"testing"

	"github.com/kaidev1024/leetcode-golang/utils/cmp"
)

var arr = []int{1, 2, 3}

func TestPlusOne(t *testing.T) {

	result := plusOne(arr)
	expectedResult := []int{1, 2, 4}

	if cmp.CmpIntSlice(result, expectedResult) != 0 {
		t.Error("not equal array")
	}
}

func BenchmarkPlusOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		plusOne(arr)
	}
}
