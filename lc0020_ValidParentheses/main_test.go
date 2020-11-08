package main

import "testing"

func IsValidTest(t *testing.T) {
	data := "()"
	if !isValid(data) {
		t.Error("expect valid, actually invalid")
	}
}
