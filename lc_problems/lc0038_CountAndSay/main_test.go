package main

import "testing"

func TestCountAndSay(t *testing.T) {
	result := countAndSay(4)

	if result != "1211" {
		t.Error("wrong result")
	}
}
