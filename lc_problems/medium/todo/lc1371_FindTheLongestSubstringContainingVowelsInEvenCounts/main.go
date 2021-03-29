func findTheLongestSubstring(s string) int {
	n := len(s)
	counts := make([][5]int, n+1)
	a, e, i, o, u := 0, 0, 0, 0, 0

	for j := 0; j < n; j++ {
		switch s[j] {
		case 'a':
			a++
		case 'e':
			e++
		case 'i':
			i++
		case 'o':
			o++
		case 'u':
			u++
		}
		counts[j+1][0] = a
		counts[j+1][1] = e
		counts[j+1][2] = i
		counts[j+1][3] = o
		counts[j+1][4] = u
	}

	for width := n; width >= 0; width-- {
		for start := 0; start+width <= n; start++ {
			end := start + width
			if (counts[end][0]-counts[start][0])%2 == 1 {
				continue
			}
			if (counts[end][1]-counts[start][1])%2 == 1 {
				continue
			}
			if (counts[end][2]-counts[start][2])%2 == 1 {
				continue
			}
			if (counts[end][3]-counts[start][3])%2 == 1 {
				continue
			}
			if (counts[end][4]-counts[start][4])%2 == 1 {
				continue
			}
			return width
		}
	}
	return 0
}