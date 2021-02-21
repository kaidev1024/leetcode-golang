func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	result := 0
	for i, n := 1, len(points); i < n; i++ {
		result = max(result, points[i][0]-points[i-1][0])
	}

	return result
}