func removeKdigits(num string, k int) string {
	n := len(num)
	if n <= k {
		return "0"
	}
	arr := make([]byte, 0, n)
	for _, charr := range num {
		char := byte(charr)
		if k == 0 {
			arr = append(arr, char)
			continue
		}
		l := len(arr)
		if l == 0 {
			arr = append(arr, char)
			continue
		}
		for l > 0 && arr[l-1] > char && k > 0 {
			l--
			k--
		}
		arr = arr[:l]
		arr = append(arr, char)
	}
	start := 0
	end := len(arr)
	for start < end && arr[start] == '0' {
		start++
	}
	for k > 0 && start < end {
		if arr[end-1] == '0' {
			end--
			continue
		}
		end--
		k--
	}
	if start == end {
		return "0"
	}
	return string(arr[start:end])
}