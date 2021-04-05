func findAnagrams(s string, p string) []int {
	result := make([]int, 0)
	np := len(p)
	ns := len(s)
	if np > ns {
		return nil
	}

	countP := make([]int, 26)
	target := 0
	for _, v := range p {
		index := v - 'a'
		if countP[index] == 0 {
			target++
		}
		countP[index]++
	}

	countS := make([]int, 26)
	count := 0
	for i := 1; i < np; i++ {
		index := s[ns-i] - 'a'
		if countP[index] == 0 {
			continue
		}
		countS[index]++
		if countP[index] == countS[index] {
			count++
		} else if countP[index]+1 == countS[index] {
			count--
		}
	}
	for i, j := np, 1; i <= ns; i++ {
		index := s[ns-i] - 'a'
		if countP[index] != 0 {
			countS[index]++
			if countP[index] == countS[index] {
				count++
				if count == target {
					result = append(result, ns-i)
				}
			} else if countP[index]+1 == countS[index] {
				count--
			}
		}

		index = s[ns-j] - 'a'
		j++
		if countP[index] == 0 {
			continue
		}
		countS[index]--
		if countP[index] == countS[index] {
			count++
		} else if countP[index]-1 == countS[index] {
			count--
		}
	}
	return result
}