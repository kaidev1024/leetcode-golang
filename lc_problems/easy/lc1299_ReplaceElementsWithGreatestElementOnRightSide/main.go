func replaceElements(arr []int) []int {
	l := len(arr)
	maxV := arr[l-1]
	arr[l-1] = -1
	for i := l - 2; i >= 0; i-- {
		arr[i], maxV = maxV, arr[i]
		if arr[i] > maxV {
			maxV = arr[i]
		}
	}
	return arr
}