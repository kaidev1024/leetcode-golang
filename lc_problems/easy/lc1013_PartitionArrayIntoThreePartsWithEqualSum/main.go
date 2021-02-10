func canThreePartsEqualSum(arr []int) bool {
	sum := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	if sum%3 != 0 {
		return false
	}
	sum3 := sum / 3

	pos := 0
	part := 0
	for ; pos < n-2; pos++ {
		part += arr[pos]
		if part == sum3 {
			break
		}
	}
	if part != sum3 {
		return false
	}
	pos++
	part = 0
	for ; pos < n-1; pos++ {
		part += arr[pos]
		if part == sum3 {
			break
		}
	}
	if part != sum3 {
		return false
	}

	return true
}