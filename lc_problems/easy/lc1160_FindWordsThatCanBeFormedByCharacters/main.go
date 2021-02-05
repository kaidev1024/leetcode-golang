func countCharacters(words []string, chars string) int {
	arr := make([]int, 128)
	for _, c := range chars {
		arr[c]++
	}

	result := 0

	for _, word := range words {
		counts := make([]int, 128)
		flag := true
		for _, v := range word {
			counts[v]++
			if counts[v] > arr[v] {
				flag = false
				break
			}
		}
		if flag {
			result += len(word)
		}
	}
	return result

}