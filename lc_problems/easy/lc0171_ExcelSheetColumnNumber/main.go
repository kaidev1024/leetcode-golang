package leetcode

func titleToNumber(s string) int {
	sum := 0
	n := len(s)
	for i := 0; i < n; i++ {
		sum = sum*26 + int(s[i]-'A'+1)
	}
	return sum
}
