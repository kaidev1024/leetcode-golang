package leetcode

func firstUniqChar(s string) int {
	l := len(s)
	if l == 1 {
		return 0
	}

	arr := make([]int, 128)

	for i := 0; i < l; i++ {
		pos := s[i]
		if arr[pos] == 0 {
			arr[pos] = i + 1
			continue
		}
		if arr[pos] > 0 {
			arr[pos] = -1
		}
	}
	result := l + 1
	for i := 0; i < 128; i++ {
		if arr[i] > 0 && arr[i] < result {
			result = arr[i]
		}
	}
	result--
	if result == l {
		return -1
	}
	return result
}
