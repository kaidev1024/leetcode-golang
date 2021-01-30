func totalMoney(n int) int {
	if n <= 7 {
		return (1 + n) * n / 2
	}

	nWeeks := n / 7
	nDays := n % 7

	return (28+28+(nWeeks-1)*7)*nWeeks/2 + (nWeeks+1+nWeeks+nDays)*nDays/2
}