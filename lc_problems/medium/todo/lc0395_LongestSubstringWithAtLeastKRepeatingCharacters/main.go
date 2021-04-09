func longestSubstring(s string, k int) int {
	n := len(s)
	if n < k {
		return 0
	}
	arr := []byte(s)

	var helper func(start, end int) int
	helper = func(start, end int) int {
		counts := make([]int, 26)
		for i := start; i <= end; i++ {
			counts[arr[i]-'a']++
		}

		ret := 0
		starti := start
		endi := starti
		isSplit := false
		for endi <= end {
			if counts[arr[endi]-'a'] < k {
				ret = max(ret, helper(starti, endi-1))
				endi++
				starti = endi
				isSplit = true
			} else {
				endi++
			}
		}
		if !isSplit {
			return end - start + 1
		}
		return max(ret, helper(starti, endi-1))
	}

	return helper(0, n-1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}