func search(nums []int, target int) bool {
	first, last := 0, len(nums)-1
	mid := 0
	for first <= last {
		mid = first + (last-first)/2
		if nums[mid] == target {
			return true
		}
		if nums[first] == nums[mid] && nums[mid] == nums[last] {
			first++
			last--
			continue
		}
		if nums[first] <= nums[mid] {
			if target >= nums[first] && target < nums[mid] {
				last = mid - 1
				continue
			}
			first = mid + 1
			continue
		}
		if target > nums[mid] && target <= nums[last] {
			first = mid + 1
			continue
		}
		last = mid - 1
	}
	return false
}