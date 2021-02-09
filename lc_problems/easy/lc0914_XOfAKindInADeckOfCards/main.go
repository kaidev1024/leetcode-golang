func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	if b%a == 0 {
		return a
	}

	if a > b {
		return gcd(a%b, b)
	}
	return gcd(a, b%a)
}

func hasGroupsSizeX(deck []int) bool {
	n := len(deck)
	if n < 2 {
		return false
	}

	counts := make(map[int]int)

	for i := 0; i < n; i++ {
		counts[deck[i]]++
	}

	div := counts[deck[0]]
	for _, v := range counts {
		div = gcd(div, v)
		if div == 1 {
			return false
		}
	}
	return true
}