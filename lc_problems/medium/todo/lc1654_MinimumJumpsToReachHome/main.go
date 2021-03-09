func minimumJumps(forbidden []int, a int, b int, x int) int {
	n := 6001
	dp := make([][]int, 2)
	dp[0] = make([]int, n)
	dp[1] = make([]int, n)
	for i := 0; i < n; i++ {
		dp[0][i] = math.MaxInt32
		dp[1][i] = math.MaxInt32
	}

	for _, pos := range forbidden {
		dp[0][pos] = -1
		dp[1][pos] = -1
	}

	queue := make([]entry, 0)
	queue = append(queue, entry{0, 0})

	lq := 1
	step := 0
	dp[0][0] = 0
	for lq > 0 {
		for i := 0; i < lq; i++ {
			cur := queue[i]
			if cur.pos == x {
				return min(dp[0][x], dp[1][x])
			}
			next := cur.pos + a
			if next < n && dp[1][next] == math.MaxInt32 {
				queue = append(queue, entry{next, 1})
				dp[1][next] = step + 1
			}
			next = cur.pos - b
			if cur.dir == 1 && next >= 0 && dp[0][next] == math.MaxInt32 {
				queue = append(queue, entry{next, 0})
				dp[0][next] = step + 1
			}
		}
		queue = queue[lq:]
		lq = len(queue)
		step++
	}

	result := min(dp[0][x], dp[1][x])
	if result == math.MaxInt32 {
		return -1
	}
	return result
}

type entry struct {
	pos int
	dir int
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}