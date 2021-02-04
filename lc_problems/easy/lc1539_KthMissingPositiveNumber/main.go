func findKthPositive(arr []int, k int) int {
	l := len(arr)

	count := arr[0] - 1
	if k <= count {
		return k
	}
	k -= count
	for i := 1; i < l; i++ {
		count = arr[i] - arr[i-1] - 1
		if k <= count {
			return arr[i-1] + k
		}
		k -= count
	}
	return arr[l-1] + k
}