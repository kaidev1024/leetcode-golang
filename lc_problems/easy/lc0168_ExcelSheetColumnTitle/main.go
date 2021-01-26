package leetcode

import "strings"

func reverse(s string) string {
	var sb strings.Builder
	l := len(s)
	for i := l - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

func convertToTitle(n int) string {
	var sb strings.Builder

	for n > 0 {
		n--
		res := n % 26
		sb.WriteByte(byte('A' + res))
		n /= 26
	}
	return reverse(sb.String())
}
