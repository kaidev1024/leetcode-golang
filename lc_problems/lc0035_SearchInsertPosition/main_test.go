package main

import "testing"

func TestSearchInsert(t *testing.T) {
	arr := []int{1, 3, 5, 6}
	target1 := 2
	target2 := 5

	result1 := searchInsert(arr, target1)
	result2 := searchInsert(arr, target2)
	if result1 != 1 {
		t.Error("wrong position")
	}
	if result2 != 2 {
		t.Error("wrong position")
	}
}
