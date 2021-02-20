func minimumLength(s string) int {
	n := len(s)
	start := 0
	end := n - 1
	for start < end {
		if s[start] != s[end] {
			return end - start + 1
		}
		prefixInd := start + 1
		for ; s[prefixInd] == s[start] && prefixInd < end; prefixInd++ {
		}
		start = prefixInd
		if prefixInd == end {
			return 0
		}
		suffixInd := end - 1
		for ; s[suffixInd] == s[end] && prefixInd < suffixInd; suffixInd-- {
		}
		end = suffixInd

	}
	return 1
}