package leetcode

func isAnagram(s string, t string) bool {
	l := len(s)
	if l != len(t) {
		return false
	}

	counts := make([]int, 128)

	for i := 0; i < l; i++ {
		counts[s[i]]++
		counts[t[i]]--
	}

	for i := 0; i < 128; i++ {
		if counts[i] != 0 {
			return false
		}
	}
	return true
}
