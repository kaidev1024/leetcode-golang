func increasingTriplet(nums []int) bool {
	seq := make([]int, 0)
	for _, v := range nums {
		ns := len(seq)
		if ns == 0 {
			seq = append(seq, v)
			continue
		}
		if ns == 1 {
			if v > seq[0] {
				seq = append(seq, v)
			} else if v < seq[0] {
				seq[0] = v
			}
			continue
		}
		if v > seq[1] {
			return true
		}
		if v < seq[0] {
			seq[0] = v
		}
		if v > seq[0] && v < seq[1] {
			seq[1] = v
		}
	}
	return false
}