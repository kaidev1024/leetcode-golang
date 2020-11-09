package main

import (
	"testing"

	deepequal "github.com/kaidev1024/go-deep-equal"
)

var nums = []int{2, 7, 11, 15}
var target = 9

func TestTwoSum(t *testing.T) {
	result := twoSum(nums, target)
	expectedResult := []int{1, 0}

	if !deepequal.IsSliceEqual(result, expectedResult) {
		t.Error("Expected [1 0] but got ", result)
	}
}
