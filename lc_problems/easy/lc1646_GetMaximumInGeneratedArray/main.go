func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	arr := make([]int, n+1)
	arr[1] = 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			arr[i] = arr[i/2]
		} else {
			arr[i] = arr[i/2] + arr[i/2+1]
		}
	}
	result := 0
	for i := 0; i < n+1; i++ {
		if arr[i] > result {
			result = arr[i]
		}
	}
	return result
}