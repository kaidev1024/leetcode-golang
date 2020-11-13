package main

import "testing"

func TestIsSameTree(t *testing.T) {
	if !isSameTree(t1, t2) {
		t.Error("should be same tree")
	}
	if isSameTree(t1, t3) {
		t.Error("not same tree")
	}
}
