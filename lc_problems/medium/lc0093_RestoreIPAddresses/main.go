func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	backtrack(s, &res, nil)
	return res
}

func backtrack(s string, res *[]string, cur []string) {
	if len(s) == 0 && len(cur) == 4 {
		*res = append(*res, strings.Join(cur, "."))
		return
	}

	if len(cur) > 4 {
		return
	}

	for i := 1; i <= 3 && i <= len(s); i++ {
		if isValid(s[0:i]) {
			backtrack(s[i:], res, append(cur, s[0:i]))
		}
	}
}

func isValid(s string) bool {
	num, _ := strconv.Atoi(s)
	if num > 255 {
		return false
	}

	if strings.HasPrefix(s, "0") && len(s) != 1 {
		return false
	}

	return true
}