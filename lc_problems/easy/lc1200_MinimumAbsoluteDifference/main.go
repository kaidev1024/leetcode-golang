func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	n := len(arr)

	result := make([][]int, 0, n-1)
	minD := arr[1] - arr[0]
	result = append(result, []int{arr[0], arr[1]})

	for i := 2; i < n; i++ {
		d := arr[i] - arr[i-1]
		if d == minD {
			result = append(result, []int{arr[i-1], arr[i]})
		} else if d < minD {
			minD = d
			result = result[0:0]
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}