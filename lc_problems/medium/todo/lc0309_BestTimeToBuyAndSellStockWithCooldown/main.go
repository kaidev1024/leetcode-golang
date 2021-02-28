func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	prevBuy, buy, prevSell, sell := 0, -prices[0], 0, 0
	for i := 1; i < n; i++ {
		prevBuy = buy
		buy = max(prevSell-prices[i], buy)
		prevSell = sell
		sell = max(sell, prevBuy+prices[i])
	}
	return sell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}