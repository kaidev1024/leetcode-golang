func shortestCompletingWord(licensePlate string, words []string) string {

	s := ""
	minLen := 16
	index := make([]int, 26, 26)
	tmp := strings.ToLower(licensePlate)

	for _, v := range tmp {
		if v >= 'a' && v <= 'z' {
			index[v-'a']++
		}
	}
	for _, v := range words {
		if len(v) >= minLen {
			continue
		}
		if !matches(index, v) {
			continue
		}
		minLen = len(v)
		s = v
	}

	return s

}

func matches(index []int, s string) bool {
	temp := make([]int, 26, 26)
	for _, v := range s {
		temp[v-'a']++
	}
	for i := 0; i < 26; i++ {
		if index[i] > temp[i] {
			return false
		}
	}
	return true
}