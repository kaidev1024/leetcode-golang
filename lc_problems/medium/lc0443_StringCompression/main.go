func compress(chars []byte) int {
	n := len(chars)
	if n == 1 {
		return 1
	}
	iter := 0
	start := 0
	end := 0
	for start < n {
		char := chars[start]
		end = start + 1
		for end < n && char == chars[end] {
			end++
		}
		count := strconv.Itoa(end - start)
		iter++
		if count == "1" {
			start++
			if start != n {
				chars[iter] = chars[start]
			}
			continue
		}
		for i, lc := 0, len(count); i < lc; i++ {
			chars[iter] = count[i]
			iter++
		}
		if end != n {
			chars[iter] = chars[end]
		}

		start = end
	}
	chars = chars[:iter]
	return iter
}