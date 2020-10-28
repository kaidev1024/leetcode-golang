package main

import "testing"

func TestRomanToInt(t *testing.T) {
	s1 := "III"
	s2 := "IV"
	if romanToInt(s1) != 3 {
		t.Error("should be equal")
	}
	if romanToInt(s2) != 4 {
		t.Error("should be equal")
	}
}
