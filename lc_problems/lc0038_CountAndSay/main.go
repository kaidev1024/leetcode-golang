package main

import (
	"fmt"
	"strings"
)

func countAndSay(n int) string {
	result := "1"
	for ; n > 1; n-- {
		result = countUtil(result)
	}

	return result
}

func countUtil(s string) string {
	cur := s[0]
	count := 1
	l := len(s)
	var result strings.Builder

	for i := 1; i < l; i++ {
		if s[i] == cur {
			count++
		} else {
			result.WriteByte(byte(count + '0'))
			result.WriteByte(cur)
			count = 1
			cur = s[i]
		}
	}
	result.WriteByte(byte(count + '0'))
	result.WriteByte(cur)

	return result.String()
}

func main() {
	result := countAndSay(4)
	fmt.Println(result)
}
