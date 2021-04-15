func checkSubarraySum(nums []int, k int) bool {
	m := map[int]int{
		0: -1,
	}

	sum := 0
	for i, v := range nums {
		sum = (sum + v) % k
		if pre, ok := m[sum]; ok {
			if (i - pre) > 1 {
				return true
			} else {
				continue
			}
		}
		m[sum] = i
	}
	return false
}