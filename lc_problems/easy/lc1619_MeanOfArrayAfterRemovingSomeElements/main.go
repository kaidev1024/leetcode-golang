func trimMean(arr []int) float64 {
	sort.Ints(arr)

	l := len(arr)
	skip := l / 20

	sum := 0
	for i := skip; i < l-skip; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(l-2*skip)
}