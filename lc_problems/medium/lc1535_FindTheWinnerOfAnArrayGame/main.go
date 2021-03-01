func getWinner(arr []int, k int) int {
	n := len(arr)
	num := arr[0]
	if k > n-1 {
		k = n - 1
	}

	c := 0
	for i := 1; i < n; i++ {
		if num < arr[i] {
			num = arr[i]
			c = 1
		} else {
			c++
		}
		if c == k {
			return num
		}
	}
	return num
}