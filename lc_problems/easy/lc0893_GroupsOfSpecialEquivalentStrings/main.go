func numSpecialEquivGroups(A []string) int {
	helper := func(str string) string {
		odd := make([]int, 26)
		even := make([]int, 26)
		for i, l := 0, len(str); i < l; i++ {
			if i%2 == 0 {
				even[str[i]-'a']++
			} else {
				odd[str[i]-'a']++
			}
		}

		result := ""
		for i := 0; i < 26; i++ {
			if odd[i] > 0 {
				result += fmt.Sprintln("%c%v", 'a'+i, odd[i])
			}
		}
		result += "-"
		for i := 0; i < 26; i++ {
			if even[i] > 0 {
				result += fmt.Sprintln("%c%v", 'a'+i, even[i])
			}
		}
		return result
	}

	counts := make(map[string]int)
	for i, l := 0, len(A); i < l; i++ {
		counts[helper(A[i])]++
	}
	return len(counts)
}