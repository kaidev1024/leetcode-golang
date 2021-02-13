func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	first := cost[0]
	second := cost[1]
	for i := 2; i < n; i++ {
		temp := cost[i] + min(first, second)
		first = second
		second = temp
	}
	return min(first, second)
}