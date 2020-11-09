package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	size := removeDuplicates(nums)

	if size != 2 {
		t.Error("wrong array")
	}
}
