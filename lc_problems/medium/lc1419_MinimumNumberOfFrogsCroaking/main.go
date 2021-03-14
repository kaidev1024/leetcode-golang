func minNumberOfFrogs(croakOfFrogs string) int {
	result := 0
	counts := make([]int, 5)
	for _, v := range croakOfFrogs {
		if v == 'c' {
			counts[0]++
			if counts[0] > result {
				result = counts[0]
			}
		} else if v == 'r' {
			counts[1]++
			if counts[1] > counts[0] {
				return -1
			}
		} else if v == 'o' {
			counts[2]++
			if counts[2] > counts[0] || counts[2] > counts[1] {
				return -1
			}
		} else if v == 'a' {
			counts[3]++
			if counts[3] > counts[0] || counts[3] > counts[1] || counts[3] > counts[2] {
				return -1
			}
		} else if v == 'k' {
			counts[4]++
			if counts[4] > counts[3] || counts[4] > counts[0] || counts[4] > counts[1] || counts[4] > counts[2] {
				return -1
			}
			for i := 0; i < 5; i++ {
				counts[i]--
			}
		} else {
			return -1
		}
	}
	for _, v := range counts {
		if v > 0 {
			return -1
		}
	}
	return result
}