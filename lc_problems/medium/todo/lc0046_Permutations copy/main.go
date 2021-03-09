func permute(nums []int) [][]int {
	n := len(nums)
	if n == 1 {
		return [][]int{nums}
	}

	ans := make([][]int, 0)

	for i := 0; i < n; i++ {
		v := make([]int, n)
		copy(v, nums)

		v[0], v[i] = v[i], v[0]
		v = v[1:]
		ans1 := permute(v)

		for _, j := range ans1 {
			j = append([]int{nums[i]}, j...)
			ans = append(ans, j)
		}
	}
	return ans
}