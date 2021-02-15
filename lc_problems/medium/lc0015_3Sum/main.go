func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return nil
	}
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < n-2; i++ {

		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l-1] == nums[l] {
					l++
				}
				for l < r && nums[r+1] == nums[r] {
					r--
				}
				continue
			}
			if sum < 0 {
				l++
			} else {
				r--
			}
		}
		for nums[i] == nums[i+1] && i < n-2 {
			i++
		}
	}
	return result
}