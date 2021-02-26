func waysToMakeFair(nums []int) int {
	post0 := 0
	post1 := 0
	n := len(nums)
	for i := 1; i < n; i += 2 {
		post1 += nums[i]
		if i+1 < n {
			post0 += nums[i+1]
		}
	}

	count := 0
	pre0 := 0
	pre1 := 0
	for i := 0; i < n; i++ {
		sum0 := pre0 + post0
		sum1 := pre1 + post1
		if i > 0 {
			sum0 = pre0 + post1
			sum1 = pre1 + post0
		}
		if sum0 == sum1 {
			count++
		}

		if i%2 == 0 {
			pre0 += nums[i]
		} else {
			pre1 += nums[i]
		}

		if i+1 >= n {
			continue
		}
		if (i+1)%2 == 0 {
			post0 -= nums[i+1]
		} else {
			post1 -= nums[i+1]
		}
	}
	return count
}