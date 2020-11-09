package main

import "testing"

func TestStrStr(t *testing.T) {
	haystack := "hello"
	needle := "ll"

	ind := strStr(haystack, needle)
	if ind != 2 {
		t.Error("wrong index")
	}
}
