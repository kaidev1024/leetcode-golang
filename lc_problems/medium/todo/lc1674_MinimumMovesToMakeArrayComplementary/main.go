func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minMoves(nums []int, limit int) int {
	counts := make([]int, 2+2*limit)
	n := len(nums)
	n2 := n >> 1

	for i := 0; i < n2; i++ {
		a, b := nums[i], nums[n-1-i]
		counts[2] += 2
		counts[min(a, b)+1]--
		counts[a+b]--
		counts[a+b+1]++
		counts[max(a, b)+limit+1]++
	}

	result := 2 * n
	sum := 0
	for i := 2; i < 2+2*limit; i++ {
		sum += counts[i]
		result = min(result, sum)
	}
	return result
}