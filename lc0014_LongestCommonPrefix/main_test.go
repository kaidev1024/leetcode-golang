package main

import "testing"

func LongestCommonPrefixText(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}

	result := longestCommonPrefix(strs)
	if result != "fl" {
		t.Error("should be equal")
	}
}
