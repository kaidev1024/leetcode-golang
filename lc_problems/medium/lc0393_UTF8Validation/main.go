func check(num, mask, result int) bool {
	return num&mask == result
}

func validUtf8(data []int) bool {
	mask1 := 0b10000000
	mask2 := 0b11000000
	mask3 := 0b11100000
	mask4 := 0b11110000
	mask5 := 0b11111000

	for i, n := 0, len(data); i < n; {
		if check(data[i], mask1, 0) {
			i++
			continue
		}
		if check(data[i], mask3, mask2) {
			if i+2 > n {
				return false
			}
			if !check(data[i+1], mask2, mask1) {
				return false
			}
			i += 2
			continue
		}
		if check(data[i], mask4, mask3) {
			if i+3 > n {
				return false
			}
			if !check(data[i+1], mask2, mask1) || !check(data[i+2], mask2, mask1) {
				return false
			}
			i += 3
			continue
		}
		if check(data[i], mask5, mask4) {
			if i+4 > n {
				return false
			}
			if !check(data[i+1], mask2, mask1) || !check(data[i+2], mask2, mask1) || !check(data[i+3], mask2, mask1) {
				return false
			}
			i += 4
			continue
		}
		return false
	}
	return true
}

