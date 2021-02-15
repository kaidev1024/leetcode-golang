func mostCommonWord(paragraph string, banned []string) string {

	tmp := ""
	temp := make(map[string]int)
	max := math.MinInt32
	res := ""

	for i, v := range paragraph {
		if 'A' <= v && v <= 'Z' {
			tmp += string(v + 32)
			if i == len(paragraph)-1 {
				temp[tmp]++
			}
		} else if 'a' <= v && v <= 'z' {
			tmp += string(v)
			if i == len(paragraph)-1 {
				temp[tmp]++
			}
		} else {
			if tmp != "" {
				temp[tmp]++
				tmp = ""
			}
		}
	}
	for _, v := range banned {
		if _, ok := temp[v]; ok {
			temp[v] = 0
		}
	}
	for i, v := range temp {
		if v > max {
			max = v
			res = i
		}
	}
	return res
}