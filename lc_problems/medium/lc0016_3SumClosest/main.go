func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)

	sum := nums[0] + nums[1] + nums[2]
	dif := diff(sum, target)
	result := sum

	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			}
			if diff(sum, target) < dif {
				dif = diff(sum, target)
				result = sum
			}
			if sum > target {
				r--
			} else {
				l++
			}
		}
	}
	return result
}