func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 1 {
		return 1
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	result := 0
	for i := 0; i < n; {

		end := points[i][1]
		for j := i + 1; ; j++ {
			if j < n && points[j][0] <= end {
				continue
			}
			result++
			i = j
			break
		}
	}
	return result
}