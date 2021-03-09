func maxProfit(inventory []int, orders int) int {
	sort.Ints(inventory)
	inventory = append([]int{0}, inventory...)

	result := 0
	n := len(inventory)
	curNum := inventory[n-1]
	for i := n - 2; i >= 0; i-- {
		width := n - 1 - i
		height := min(orders/width, curNum-inventory[i])
		if height == 0 {
			continue
		}
		upper := curNum
		lower := curNum - height + 1
		result += (upper + lower) * height * width >> 1
		result %= 1e9 + 7
		curNum -= height
		orders -= width * height
	}
	if orders > 0 {
		result += orders * curNum
	}
	return result % (1e9 + 7)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}