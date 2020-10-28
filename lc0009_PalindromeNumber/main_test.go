package main

import "testing"

func TestPalindrome(t *testing.T) {
	x1 := 100
	x2 := 121
	if isPalindrome(x1) {
		t.Error("expected not a palindrome")
	}
	if !isPalindrome(x2) {
		t.Error("expected to be a palindrome")
	}
}
