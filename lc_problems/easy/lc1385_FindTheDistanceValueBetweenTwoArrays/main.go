func getLimit(a, d int) (int, int) {
	return a - d, a + d
}

func getLimitArr(arr []int, d int) (int, int) {
	minVal, maxVal := arr[0], arr[0]

	for _, val := range arr {
		l, u := getLimit(val, d)
		if l < minVal {
			minVal = l
		}
		if u > maxVal {
			maxVal = u
		}
	}
	return minVal, maxVal
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	minVal, maxVal := getLimitArr(arr2, d)

	r := 0
	for _, val := range arr1 {
		if val < minVal && val > maxVal {
			r++
		}
	}
	return r
}