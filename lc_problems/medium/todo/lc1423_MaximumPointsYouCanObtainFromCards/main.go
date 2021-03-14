func maxScore(cardPoints []int, k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		sum += cardPoints[i]
	}

	temp := sum

	for i, j := k-1, len(cardPoints)-1; i >= 0; i, j = i-1, j-1 {
		temp -= cardPoints[i]
		temp += cardPoints[j]
		sum = max(sum, temp)
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}