func findSubstringInWraproundString(p string) int {
	counts := make([]int, 26)
	i := 0
	n := len(p)
	for i < n {
		j := i + 1
		for j < n {
			if p[j] == p[j-1]+1 || (p[j] == 'a' && p[j-1] == 'z') {
				j++
			} else {
				break
			}
		}

		for k := i; k < j; k++ {
			counts[p[k]-'a'] = max(counts[p[k]-'a'], k+1-i)
		}
		i = j
	}
	ret := 0
	for _, v := range counts {
		ret += v
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}