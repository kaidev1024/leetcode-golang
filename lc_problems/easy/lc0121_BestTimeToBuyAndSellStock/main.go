func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	lowestPrice := prices[0]
	result := 0
	for _, val := range prices[1:] {
		if lowestPrice > val {
			lowestPrice = val
		} else {
			diff := val - lowestPrice
			if diff > result {
				result = diff
			}
		}
	}
	return result
}