func countSubstrings(s string, t string) int {
	var res int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			x := i
			y := j
			var diff int
			for x < len(s) && y < len(t) && diff <= 1 {
				if s[x] != t[y] {
					diff++
				}
				if diff == 1 {
					res++
				}
				x++
				y++
			}
		}
	}
	return res
}