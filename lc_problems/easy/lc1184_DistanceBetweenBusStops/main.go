func distanceBetweenBusStops(distance []int, start int, destination int) int {
	n := len(distance)
	sums := make([]int, n+1)

	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + distance[i-1]
	}
	if destination < start {
		start, destination = destination, start
	}
	d1 := sums[destination] - sums[start]
	d2 := sums[n] - d1
	if d1 < d2 {
		return d1
	}
	return d2
}