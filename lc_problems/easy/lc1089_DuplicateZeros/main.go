func duplicateZeros(arr []int) {
	l := len(arr)

	for i := 0; i < l; {
		if arr[i] != 0 {
			i++
			continue
		}
		i++
		if i == l {
			return
		}
		arr = append(arr[:i+1], arr[i:l-1]...)
		arr[i] = 0
		i++
	}
}