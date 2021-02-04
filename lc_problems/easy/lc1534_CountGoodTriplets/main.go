func countGoodTriplets(arr []int, a int, b int, c int) int {
	dist := func(v1, v2 int) int {
		if v1 > v2 {
			return v1 - v2
		}
		return v2 - v1
	}

	r := 0
	l := len(arr)
	for i := 0; i < l-2; i++ {
		for j := i + 1; j < l-1; j++ {
			if dist(arr[i], arr[j]) > a {
				continue
			}
			for k := j + 1; k < l; k++ {
				if dist(arr[j], arr[k]) > b {
					continue
				}
				if dist(arr[i], arr[k]) <= c {
					r++
				}
			}
		}
	}
	return r
}