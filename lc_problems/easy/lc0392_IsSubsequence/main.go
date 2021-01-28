package leetcode

func isSubsequence(s string, t string) bool {
	ls := len(s)
	if ls == 0 {
		return true
	}
	lt := len(t)
	if ls > lt {
		return false
	}

	spos := 0

	for i := 0; i < lt; i++ {
		if s[spos] == t[i] {
			spos++
			if spos == ls {
				return true
			}
		}
	}

	return false
}
