func numTrees(n int) int {
	counts := make([]int, n+1)
	counts[0] = 1
	counts[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			counts[i] += counts[j] * counts[i-1-j]
		}
	}
	return counts[n]
}