package leetcode

func isHappy(n int) bool {
	m := make(map[int]struct{})

	for {
		sum := 0
		for n > 0 {
			res := n % 10
			n /= 10
			sum += res * res
		}
		if sum == 1 {
			return true
		}
		if _, ok := m[sum]; ok {
			return false
		}
		m[sum] = struct{}{}
		n = sum
	}
}
