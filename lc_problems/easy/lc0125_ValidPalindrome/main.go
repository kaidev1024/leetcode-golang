package leetcode

import "unicode"

func isPalindrome(s string) bool {
	n := len(s)
	if n < 2 {
		return true
	}
	left := 0
	right := n - 1
	for left < right {
		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsNumber(rune(s[left])) {
			left++
		}
		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsNumber(rune(s[right])) {
			right--
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}
