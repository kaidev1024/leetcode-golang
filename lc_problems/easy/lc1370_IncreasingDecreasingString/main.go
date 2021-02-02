func sortString(s string) string {
	n := len(s)

	counts := make([]int, 26)
	for i := 0; i < n; i++ {
		counts[s[i]-'a']++
	}

	var sb strings.Builder
	count := 0
	for n != count {
		for i := 0; i < 26; i++ {
			if counts[i] > 0 {
				sb.WriteByte('a' + byte(i))
				counts[i]--
				count++
			}
		}
		for i := 25; i >= 0; i-- {
			if counts[i] > 0 {
				sb.WriteByte('a' + byte(i))
				counts[i]--
				count++
			}
		}
	}
	return sb.String()
}