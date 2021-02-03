func finalPrices(prices []int) []int {
	n := len(prices)

	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		ls := len(stack)
		if ls == 0 || prices[stack[ls-1]] < prices[i] {
			stack = append(stack, i)
		} else {
			pos := ls - 1
			for pos >= 0 && prices[stack[pos]] >= prices[i] {
				prices[stack[pos]] -= prices[i]
				pos--
			}
			stack = stack[:pos+1]
			stack = append(stack, i)
		}
	}
	return prices
}