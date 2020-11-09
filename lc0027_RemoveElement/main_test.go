package main

import "testing"

func TestRemoveElement(t *testing.T) {
	arr := []int{3, 2, 2, 3}
	ele := 3

	result := removeElement(arr, ele)
	if result != 2 {
		t.Error("wrong result")
	}
}
