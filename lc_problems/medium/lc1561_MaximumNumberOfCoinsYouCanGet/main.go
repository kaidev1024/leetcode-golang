func maxCoins(piles []int) int {
	sort.Ints(piles)
	n := len(piles)
	result := 0
	count := n / 3
	for i, j := n-2, 0; j < count; i -= 2 {
		result += piles[i]
		j++
	}
	return result
}