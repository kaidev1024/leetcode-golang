func arrayNesting(nums []int) int {
	n := len(nums)
	ret := 0
	for i := 0; i < n; i++ {
		if nums[i] == -1 {
			continue
		}
		arr := make([]int, 0)
		j := i
		for {
			arr = append(arr, nums[j])
			prev := j
			j = nums[j]
			nums[prev] = -1
			if j == i {
				ret = max(ret, len(arr))
				break
			}
		}
	}
	return ret
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}