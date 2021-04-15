func findLongestWord(s string, d []string) string {
	res := ""
	for _, v := range d {
		if isSubSequence(v, s) {
			if len(v) > len(res) || len(v) == len(res) && v < res {
				res = v
			}
		}
	}
	return res
}

func isSubSequence(p string, s string) bool {
	i := 0
	for j := 0; j < len(s) && i < len(p); j++ {
		if p[i] == s[j] {
			i++
		}
	}
	return i == len(p)
}