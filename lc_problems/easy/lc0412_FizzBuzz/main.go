package leetcode

func longestPalindrome(s string) int {
	counts := make([]int, 128)
	for i, l := 0, len(s); i < l; i++ {
		counts[s[i]]++
	}
	hasOdd := false
	result := 0
	for i := 0; i < 128; i++ {
		c := counts[i]
		if c%2 == 1 {
			result += c - 1
			if !hasOdd {
				hasOdd = true
			}
		} else {
			result += c
		}
	}
	if hasOdd {
		result++
	}
	return result
}
