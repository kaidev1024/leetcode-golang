func maximumGain(s string, x int, y int) int {
	ret := 0
	countA := 0
	countB := 0
	for _, char := range s {
		if char == 'a' {
			if x >= y {
				countA++
			} else {
				if countB > 0 {
					countB--
					ret += y
				} else {
					countA++
				}
			}
		} else if char == 'b' {
			if y >= x {
				countB++
			} else {
				if countA > 0 {
					countA--
					ret += x
				} else {
					countB++
				}
			}
		} else {
			ret += min(countA, countB) * min(x, y)
			countA = 0
			countB = 0
		}
	}
	ret += min(countA, countB) * min(x, y)
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}