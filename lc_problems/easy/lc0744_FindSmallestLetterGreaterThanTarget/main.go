func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	l := 0
	r := n - 1
	if target >= letters[r] {
		return letters[l]
	}
	for l < r {
		mid := (l + r) >> 1
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return letters[l]
}