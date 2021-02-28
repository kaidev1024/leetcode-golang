func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]struct{})
	for i, n := 0, len(wordDict); i < n; i++ {
		words[wordDict[i]] = struct{}{}
	}

	dp := make(map[int]bool)

	var helper func(s string, pos int) bool
	helper = func(s string, pos int) bool {
		if len(s) == 0 {
			return true
		}

		if v, ok := dp[pos]; ok {
			return v
		}

		for i, n := 1, len(s); i <= n; i++ {
			if _, ok := words[s[:i]]; ok {
				if helper(s[i:], pos+i) {
					dp[pos+i] = true
					return true
				} else {
					dp[pos+i] = false
				}
			}
		}
		return false
	}
	return helper(s, 0)
}