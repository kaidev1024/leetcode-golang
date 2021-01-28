package leetcode

func countSegments(s string) int {
	l := len(s)

	start := 0
	end := l - 1
	for ; start <= end && s[start] == ' '; start++ {
	}
	for ; start <= end && s[end] == ' '; end-- {
	}
	if start > end {
		return 0
	}
	count := 1
	for i := start + 1; i <= end; i++ {
		if s[i-1] == ' ' && s[i] != ' ' {
			count++
		}
	}
	return count
}
