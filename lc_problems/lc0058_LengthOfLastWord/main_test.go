package main

import "testing"

var word = "hello world"

func TestLengthOfLastWord(t *testing.T) {
	result := lengthOfLastWord(word)
	if result != 5 {
		t.Error("expect 5")
	}
}

func BenchmarkLengthOfLastWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLastWord(word)
	}
}
