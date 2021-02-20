func maximumScore(a int, b int, c int) int {
	arr := []int{a, b, c}
	sort.Ints(arr)

	if arr[0]+arr[1] <= arr[2] {
		return arr[0] + arr[1]
	}

	return (arr[0] + arr[1] + arr[2]) >> 1
}