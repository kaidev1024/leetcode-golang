func validMountainArray(arr []int) bool {
	n := len(arr)
	l, r := 0, n-1
	for l+1 < n && arr[l] < arr[l+1] {
		l++
	}
	for r-1 > 0 && arr[r] < arr[r-1] {
		r--
	}
	return l > 0 && r < n-1 && l == r
}