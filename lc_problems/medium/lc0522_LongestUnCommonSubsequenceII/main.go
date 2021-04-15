func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		li := len(strs[i])
		lj := len(strs[j])
		if li == lj {
			return strs[i] < strs[j]
		}
		return li > lj
	})
	n := len(strs)
	for start := 0; start < n; {
		end := start + 1
		for end < n {
			if strs[start] == strs[end] {
				end++
			} else {
				break
			}
		}
		if end-start > 1 {
			start = end
			continue
		}
		flag := true
		for pre := start - 1; pre >= 0; pre-- {
			if isSub(strs[start], strs[pre]) {
				flag = false
				start = end
				break
			}
		}
		if flag {
			return len(strs[start])
		}
	}
	return -1
}

func isSub(a, b string) bool {
	ia := 0
	na := len(a)
	for i := 0; i < len(b); i++ {
		if b[i] == a[ia] {
			ia++
			if ia == na {
				return true
			}
		}
	}
	return false
}