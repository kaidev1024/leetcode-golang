func minimumDeletions(s string) int {
	bCount := 0
	r := 0
	for _, v := range s {
		if v == 'b' {
			bCount++
		} else {
			if bCount > 0 {
				bCount--
				r++
			}
		}
	}
	return r
}