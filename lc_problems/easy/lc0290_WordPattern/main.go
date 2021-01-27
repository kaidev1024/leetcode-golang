package leetcode

import "strings"

func wordPattern(pattern string, s string) bool {
	l := len(pattern)
	arr := strings.Fields(s)
	if l != len(arr) {
		return false
	}

	mapping := make(map[byte]string)
	set := make(map[string]struct{})

	for i := 0; i < l; i++ {
		char := pattern[i]
		str := arr[i]
		if val, ok := mapping[char]; ok {
			if val != str {
				return false
			}
			continue
		}
		if _, ok := set[str]; ok {
			return false
		}

		mapping[char] = str
		set[str] = struct{}{}
	}

	return true
}
