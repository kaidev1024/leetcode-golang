func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}
	arr := make([]int, n+1)
	arr[1] = 1
	arr[2] = 1
	sum := 2
	for i := 3; i <= n; i++ {
		arr[i] = sum
		sum += sum - arr[i-3]
	}

	return arr[n]
}