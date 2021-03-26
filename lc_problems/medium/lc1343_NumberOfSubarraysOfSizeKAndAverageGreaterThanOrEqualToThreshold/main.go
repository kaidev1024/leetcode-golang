func numOfSubarrays(arr []int, k int, threshold int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	sum := 0
	count := 0
	product := k * threshold

	for i, val := range arr {
		sum += val
		if i >= k-1 {
			if sum >= product {
				count++
			}

			// remove the most left element from the sum since it's falling out of threshold
			sum -= arr[i-k+1]
		}
	}

	return count
}