package leetcode

func reverseString(s []byte) {
	l := 0
	r := len(s) - 1
	for l < r {
		cache := s[l]
		s[l] = s[r]
		s[r] = cache
		l++
		r--
	}
}
