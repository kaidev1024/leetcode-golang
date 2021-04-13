func nextGreaterElements(nums []int) []int {
	maxNum := -1 << 31
	maxInd := -1

	n := len(nums)
	stack := make([]int, 0, n)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = -1
	}
	for i, v := range nums {
		if v > maxNum {
			maxNum = v
			maxInd = i
		}
		ls := len(stack)
		if ls == 0 || nums[stack[ls-1]] >= v {
			stack = append(stack, i)
			continue
		}
		for ls > 0 {
			if nums[stack[ls-1]] < v {
				result[stack[ls-1]] = v
				ls--
				stack = stack[:ls]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	for i, v := range nums[:maxInd+1] {
		ls := len(stack)
		if ls == 0 || nums[stack[ls-1]] >= v {
			stack = append(stack, i)
			continue
		}
		for ls > 0 {
			if nums[stack[ls-1]] < v {
				result[stack[ls-1]] = v
				ls--
				stack = stack[:ls]
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	return result
}