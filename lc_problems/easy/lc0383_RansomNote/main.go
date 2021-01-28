package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	counts := make([]int, 128)
	for _, c := range magazine {
		counts[c]++
	}

	for _, c := range ransomNote {
		if counts[c] == 0 {
			return false
		}
		counts[c]--
	}
	return true
}
