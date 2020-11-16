func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	result := 0
	prev := prices[0]
	for _, val := range prices[1:] {
		if val > prev {
			result += val - prev
		}
		prev = val
	}
	return result
}