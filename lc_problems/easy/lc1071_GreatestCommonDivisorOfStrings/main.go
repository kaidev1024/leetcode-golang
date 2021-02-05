func gcd(a, b int) int {
	if a < b {
		return gcd(b, a)
	}
	res := a % b
	if res == 0 {
		return b
	}

	return gcd(b, res)
}

func match(str1, str2 string) bool {
	l1 := len(str1)
	l2 := len(str2)

	for i, j := 0, 0; i < l1; {
		for j < l2 {
			if str1[i] == str2[j] {
				i++
				j++
			} else {
				return false
			}
		}
		j = 0
	}
	return true
}

func gcdOfStrings(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	if l1 == 0 || l2 == 0 {
		return ""
	}

	l := gcd(l1, l2)
	common := str1[0:l]

	if !match(str1, common) || !match(str2, common) {
		return ""
	}

	return common
}