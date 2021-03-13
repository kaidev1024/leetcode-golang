func numSubseq(nums []int, target int) int {
	const MOD = 1000000007
	exps := make([]int, len(nums))
	for i := range exps {
		exps[i] = 1
	}

	for i := 1; i < len(exps); i++ {
		exps[i] = (2 * exps[i-1]) % MOD
	}

	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	count := 0

	for left <= right {
		if nums[left]+nums[right] > target {
			right--
		} else {
			count = count + exps[right-left]
			left++
		}
	}

	return count % MOD
}