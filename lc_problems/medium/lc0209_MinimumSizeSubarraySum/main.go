func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	start := 0
	end := 0
	sum := 0

	result := n + 1
	for end < n {
		for end < n {
			sum += nums[end]
			end++
			if sum >= target {
				break
			}
		}
		if sum < target {
			break
		}
		dif := end - start
		if dif < result {
			result = dif
		}
		for start < end {
			sum -= nums[start]
			start++
			if sum < target {
				break
			}
		}
		dif = end - start + 1
		if dif < result {
			result = dif
		}
	}
	if result == n+1 {
		return 0
	}
	return result
}