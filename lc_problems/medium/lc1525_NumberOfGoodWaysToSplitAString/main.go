func numSplits(s string) int {
	n := len(s)
	left := make([]int, n)
	right := make([]int, n)
	arrl := make([]bool, 26)
	arrr := make([]bool, 26)
	cl, cr := 0, 0
	for i := 0; i < n; i++ {
		posl := s[i] - 'a'
		if !arrl[posl] {
			arrl[posl] = true
			cl++
		}
		left[i] = cl

		posr := s[n-1-i] - 'a'
		if !arrr[posr] {
			arrr[posr] = true
			cr++
		}
		right[n-1-i] = cr
	}

	result := 0
	for i := 0; i < n-1; i++ {
		if left[i] == right[i+1] {
			result++
		}
	}
	return result
}