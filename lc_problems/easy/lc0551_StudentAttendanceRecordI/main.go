func checkRecord(s string) bool {
	ca := 0
	cl := 0
	for i, n := 0, len(s); i < n; i++ {
		if s[i] == 'A' {
			ca++
			if ca > 1 {
				return false
			}
			cl = 0
		} else if s[i] == 'L' {
			cl++
			if cl > 2 {
				return false
			}
		} else {
			cl = 0
		}
	}
	return true
}