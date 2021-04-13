func lengthLongestPath(input string) int {
	n := len(input)
	word := make([]byte, 0, n)
	var separator byte = '/'
	curLevel := 0
	ret := 0

	pos := 0
	for pos < n {
		isFile := false
		for pos < n && input[pos] != '\n' {
			if input[pos] == '.' {
				isFile = true
			}
			word = append(word, input[pos])
			pos++
		}
		if isFile {
			ret = max(ret, len(word))
		}

		if pos == '\n' {
			return ret
		}
		pos++
		level := 0
		for pos < n && input[pos] == '\t' {
			level++
			pos++
		}
		difLevel := level - curLevel
		if difLevel > 0 {
			word = append(word, separator)
		} else {
			for difLevel <= 0 {
				index := lastIndex(word, separator)
				if difLevel == 0 {
					word = word[:index+1]
				} else {
					word = word[:index]
				}
				difLevel++
			}
		}
		curLevel = level
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lastIndex(arr []byte, char byte) int {
	i := len(arr) - 1

	for i >= 0 {
		if arr[i] == char {
			return i
		}
		i--
	}
	return -1
}