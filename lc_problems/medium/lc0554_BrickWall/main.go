func leastBricks(wall [][]int) int {
	counts := make(map[int]int)
	for _, layer := range wall {
		sum := 0
		n := len(layer)
		for i := 0; i < n-1; i++ {
			sum += layer[i]
			counts[sum]++
		}
	}

	max := 0
	for _, v := range counts {
		if v > max {
			max = v
		}
	}
	return len(wall) - max
}