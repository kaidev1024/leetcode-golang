func reverseWords(s string) string {
	var sb strings.Builder
	for i, j := len(s)-1, len(s)-1; i >= 0 && j >= 0; j = i {
		for ; j >= 0 && s[j] == ' '; j-- {
		}
		for i = j; i >= 0 && s[i] != ' '; i-- {
		}
		if i != j {
			sb.WriteString(" " + s[i+1:j+1])
		}
	}
	return sb.String()[1:]
}