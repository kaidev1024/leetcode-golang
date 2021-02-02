func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)

	sumArr := make([]int, n+1)

	for i := 1; i <= n; i++ {
		sumArr[i] = sumArr[i-1] + arr[i-1]
	}

	sum := sumArr[n]

	for i := 3; i <= n; i += 2 {
		for start := 0; start+i-1 < n; start++ {
			sum += sumArr[start+i] - sumArr[start]
		}
	}

	return sum
}