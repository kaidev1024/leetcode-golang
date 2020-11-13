package main

import (
	"testing"
)

var arr = []int{1, 2, 2, 3, 5, 6}

func TestMerge(t *testing.T) {
	merge(arr1, 3, arr2, 3)
	result := slicecmp.CmpIntSlice(arr, arr1)
	if result != 0 {
		t.Error("not equal")
	}
}
