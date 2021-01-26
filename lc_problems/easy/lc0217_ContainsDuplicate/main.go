package leetcode

func containsDuplicate(nums []int) bool {
	existing := make(map[int]struct{})

	for i, l := 0, len(nums); i < l; i++ {
		if _, ok := existing[nums[i]]; ok {
			return true
		}
		existing[nums[i]] = struct{}{}
	}
	return false
}
