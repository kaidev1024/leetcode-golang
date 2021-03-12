func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	closest := math.MaxInt32
	var helper func(cur, topInd int)
	helper = func(cur, topInd int) {
		d1, d2 := abs(cur-target), abs(closest-target)
		if d1 < d2 || (d1 == d2 && cur < closest) {
			closest = cur
		}
		if cur >= target || topInd >= len(toppingCosts) {
			return
		}
		helper(cur+toppingCosts[topInd], topInd+1)
		helper(cur+toppingCosts[topInd]*2, topInd+1)
		helper(cur, topInd+1)
	}
	for _, v := range baseCosts {
		helper(v, 0)
	}
	return closest
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}