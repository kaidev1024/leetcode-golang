func magicalString(n int) int {
	if n == 0 {
		return 0
	}
	str := []byte{'1', '2', '2'}
	if n <= 3 {
		return 1
	}
	var char byte = '1'
	index := 2
	for len(str) < n {
		count := str[index] - '0'
		str = append(str, char)
		if count == 2 {
			str = append(str, char)
		}
		index++
		if char == '1' {
			char = '2'
		} else {
			char = '1'
		}
	}
	ret := 0
	for i := 0; i < n; i++ {
		if str[i] == '1' {
			ret++
		}
	}
	return ret
}

