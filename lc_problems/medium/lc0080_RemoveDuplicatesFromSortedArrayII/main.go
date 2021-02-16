func removeDuplicates(nums []int) int {
	n := len(nums)
	fast := 1
	slow := 1
	count := 1
	for fast < n {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
			fast++
			count = 1
		} else {
			count++
			if count > 2 {
				fast++
			} else {
				nums[slow] = nums[fast]
				slow++
				fast++
			}
		}
	}
	return slow
}