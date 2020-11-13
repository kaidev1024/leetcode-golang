package main

import "testing"

func TestClimbStairs(t *testing.T) {
	result := climbStairs(2)
	if result != 2 {
		t.Error("expect 2")
	}
}
