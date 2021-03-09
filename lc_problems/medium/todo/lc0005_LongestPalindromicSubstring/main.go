func longestPalindrome(s string) string {
	tmp := make([]rune, 0)
	res := make([]rune, 0)

	tmp = append(tmp, '#')
	for _, c := range s {
		tmp = append(tmp, c)
		tmp = append(tmp, '#')
	}

	dp := make([]int, len(tmp))
	pos, maxRight := 0, 0
	center, maxLen := 0, 0

	for i := range dp {
		if i < maxRight {
			dp[i] = min(dp[2*pos-i], maxRight-i)
		} else {
			dp[i] = 1
		}

		for i-dp[i] >= 0 && i+dp[i] < len(tmp) && tmp[i-dp[i]] == tmp[i+dp[i]] {
			dp[i]++
		}

		if dp[i]+i > maxRight {
			maxRight = dp[i] + i
			pos = i
		}

		if maxLen < dp[i]-1 {
			maxLen = dp[i] - 1
			center = i
		}
	}

	for _, c := range string(tmp[center-maxLen : center+maxLen]) {
		if c != '#' {
			res = append(res, c)
		}
	}

	return string(res)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}