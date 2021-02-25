func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isAdditiveNumber(num string) bool {
	n := len(num)
	half := n >> 1

	for l1 := 1; l1 <= half; l1++ {
		if !isValidNumber(num[:l1]) {
			continue
		}
		for l2 := 1; n-l1-l2 >= max(l1, l2); l2++ {
			if !isValidNumber(num[l1 : l1+l2]) {
				continue
			}
			if isValid(num[:l1], num[l1:l1+l2], num[l1+l2:]) {
				return true
			}
		}
	}
	return false
}

func isValidNumber(n string) bool {
	if len(n) > 1 && n[0] == '0' {
		return false
	}
	return true
}

func isValid(a, b, c string) bool {
	sum := add(a, b)
	if sum == c {
		return true
	}
	n := len(sum)
	if len(c) < n {
		return false
	}
	if sum != c[:n] {
		return false
	}
	return isValid(b, sum, c[n:])
}

func add(a, b string) string {
	na := len(a)
	nb := len(b)
	if na < nb {
		return add(b, a)
	}

	result := make([]byte, na+1)
	var sum byte = 0
	for i := 1; i <= na; i++ {
		sum += a[na-i] - '0'
		if nb-i >= 0 {
			sum += b[nb-i] - '0'
		}
		result[na+1-i] = sum%10 + '0'
		sum /= 10
	}
	if sum == 0 {
		return string(result[1:])
	}
	result[0] = '1'

	return string(result)
}