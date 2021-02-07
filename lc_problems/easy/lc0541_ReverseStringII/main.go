func reverseStr(s string, k int) string {
	n := len(s)
	arr := []byte(s)

	reverse := func(start, end int) {
		if n-1 < end {
			end = n - 1
		}
		for start < end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}

	for start := 0; start < n; start += 2 * k {
		reverse(start, start+k-1)
	}

	return string(arr)
}