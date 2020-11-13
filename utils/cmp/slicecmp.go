package slicecmp

func CmpIntSlice(arr1, arr2 []int) int8 {
	l1 := len(arr1)
	l2 := len(arr2)
	if l1 > l2 {
		return -CmpIntSlice(arr2, arr1)
	}

	for i, val := range arr1 {
		if val < arr2[i] {
			return -1
		} else if val > arr2[i] {
			return 1
		}
	}
	return 0
}
