func isPrefixOfWord(sentence string, searchWord string) int {
	n := len(sentence)
	ns := len(searchWord)

	if n < ns {
		return -1
	}

	result := 0

	for i := 0; i < n; {
		for i < n && sentence[i] == ' ' {
			i++
		}
		if i == n {
			return -1
		}
		result++

		j := 0
		for j < ns {
			if sentence[i] == searchWord[j] {
				i++
				j++
			} else {
				break
			}
		}
		if j == ns {
			return result
		}

		for i < n && sentence[i] != ' ' {
			i++
		}
	}
	return -1
}