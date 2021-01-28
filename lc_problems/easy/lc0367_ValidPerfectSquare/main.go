package leetcode

func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}

	low := 0
	high := num
	for low < high {
		mid := (low + high) >> 1
		p := mid * mid
		if p == num {
			return true
		}
		if p > num {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low*low == num
}
