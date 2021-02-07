func reverseWords(s string) string {
	n := len(s)
	arr := []byte(s)
	start := 0
	end := 0

	reverse := func(start, end int) {
		for start < end {
			arr[start], arr[end] = arr[end], arr[start]
			start++
			end--
		}
	}

	for start < n && end < n {
		for start < n && s[start] == ' ' {
			start++
		}
		end = start + 1
		for end < n && s[end] != ' ' {
			end++
		}
		end--
		reverse(start, end)
		start = end + 1

	}

	return string(arr)
}